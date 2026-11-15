---
name: card-keys
description: "A skill that sets every saphan.card.* metadata key the engine reads."
metadata:
  saphan.origin: builtin
  saphan.default: "off"
  saphan.card.gives: "Read a file and tell the agent what it contains."
  saphan.card.when: "The task is to read a file the agent names."
  saphan.card.when_not: "The path does not exist or the file is binary."
  saphan.card.examples: "Read examples/valid/read-files/SKILL.md."
  saphan.card.anti_patterns: "Do not guess the file's contents; do not ask for the path again."
  saphan.card.related: "read-files"
  saphan.card.keywords: "read, file"
  saphan.card.ecosystems: "go"
---

# Card keys

Every `saphan.card.*` key the engine reads is set above. `skills verify`
refuses a malformed card with a rule of its own instead of dropping it
silently.
