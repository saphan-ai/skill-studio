<!-- saphan-decision-projection-v1
_id: example--header-stripped
status: RATIFIED
A decision-projection header wraps this skill on disk. It is not part of the
skill: the canonical hash covers the text after the header's blank line, and
nothing else.
-->

---
name: header-stripped
description: A skill stored with a decision-projection header; the canonical hash covers the text after the header.
metadata:
  saphan.origin: project
  saphan.default: "off"
---

# Header stripped

The header above is removed before hashing or signing; the agent receives the
canonical text, never the stored file.
