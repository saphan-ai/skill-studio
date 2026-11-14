package skillfmt

import (
	"fmt"
	"strings"
)

// cardKeys are the eight metadata keys the engine reads as skill cards
// (THE CONTRACT): five text keys and three comma lists. The engine trims
// text values, splits list values on "," and silently drops empty items, so
// an author never learns a card was mis-read. The profile names each
// malformation with a rule of its own instead.
var cardKeys = []string{
	"saphan.card.gives",
	"saphan.card.when",
	"saphan.card.when_not",
	"saphan.card.examples",
	"saphan.card.anti_patterns",
	"saphan.card.related",
	"saphan.card.keywords",
	"saphan.card.ecosystems",
}

// cardListKeys are the card keys the engine splits on ",".
var cardListKeys = []string{"saphan.card.related", "saphan.card.keywords", "saphan.card.ecosystems"}

// cardViolations checks the saphan.card.* keys of one front matter document.
// The rules live in code: the in-repo schema keywords cannot express an
// unknown key carrying a known prefix, the absence of an empty list item, or
// a per-item name rule after splitting on ",".
func cardViolations(doc map[string]any) []Violation {
	var vs []Violation
	md, _ := doc["metadata"].(map[string]any)
	if md == nil {
		return nil
	}
	known := map[string]bool{}
	for _, k := range cardKeys {
		known[k] = true
	}
	isList := map[string]bool{}
	for _, k := range cardListKeys {
		isList[k] = true
	}
	for _, k := range sortedKeys(md) {
		if !strings.HasPrefix(k, "saphan.card.") {
			continue
		}
		v, _ := md[k].(string)
		if !known[k] {
			vs = append(vs, Violation{
				Rule:    RuleCardKeyUnknown,
				Path:    "SKILL.md#/metadata/" + k,
				Message: fmt.Sprintf("metadata key %q is not one of the eight card keys the engine reads (saphan.card.gives, saphan.card.when, saphan.card.when_not, saphan.card.examples, saphan.card.anti_patterns, saphan.card.related, saphan.card.keywords, saphan.card.ecosystems); write one of them or omit the key", k),
			})
			continue
		}
		if v == "" || strings.TrimSpace(v) == "" {
			vs = append(vs, Violation{
				Rule:    RuleCardEmpty,
				Path:    "SKILL.md#/metadata/" + k,
				Message: fmt.Sprintf("%s is empty or only whitespace; write the card text or omit the key", k),
			})
			continue
		}
		if !isList[k] {
			continue
		}
		items := strings.Split(v, ",")
		for _, item := range items {
			if strings.TrimSpace(item) == "" {
				vs = append(vs, Violation{
					Rule:    RuleCardListItem,
					Path:    "SKILL.md#/metadata/" + k,
					Message: fmt.Sprintf("%s holds an empty comma-separated item in %q; write items separated by commas with none empty, or omit the key", k, v),
				})
				break
			}
		}
		if k != "saphan.card.related" {
			continue
		}
		for _, item := range items {
			item = strings.TrimSpace(item)
			if item == "" {
				continue
			}
			if !isCardName(item) {
				vs = append(vs, Violation{
					Rule:    RuleCardRelatedName,
					Path:    "SKILL.md#/metadata/" + k,
					Message: fmt.Sprintf("%s item %q is not a skill name (a-z, 0-9 and hyphens only, at most 64 characters); write a skill name or drop the item", k, item),
				})
				break
			}
		}
	}
	return vs
}

// isCardName is the format's own name rule: a-z, 0-9 and hyphen only, at most
// 64 characters.
func isCardName(item string) bool {
	if len(item) == 0 || len(item) > 64 {
		return false
	}
	for _, r := range item {
		if !('a' <= r && r <= 'z' || '0' <= r && r <= '9' || r == '-') {
			return false
		}
	}
	return true
}
