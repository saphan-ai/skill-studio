---
name: header-in-body
description: A skill whose text holds a header-like block below line one; only a header on the first line is removed, so this block stays in the canonical text and is hashed with everything else.
---

# Header-like text below line one

This paragraph is ordinary skill text.

<!-- saphan-decision-projection-v1
_id: not-a-header
status: RATIFIED
-->
It looks like a decision-projection header, but it is not on the first line of
the file, so no header is removed and the whole file is hashed.
