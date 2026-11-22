<!-- saphan-decision-projection-v1
_id: example--close-not-alone
status: RATIFIED
x -->
-->

---
name: close-must-be-alone
description: A skill whose header holds a line that only looks like a close; the header runs to the line that is exactly `-->`.

---

# Close must be alone

The line `x -->` does not close the header: only a line that is exactly
`-->` does. The canonical hash covers the text after the blank line that
follows it.
