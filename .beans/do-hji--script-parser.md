---
# do-hji
title: Implement Script Parser with Template Substitution
status: completed
type: task
priority: normal
created_at: 2026-02-13T00:27:46Z
updated_at: 2026-02-13T17:23:50Z
parent: do-vul
---

Create a parser that can read script files, identify specific persona blocks, and perform variable substitution (e.g., replacing `{{TASK_ID}}` with the actual ID).

### Acceptance Criteria
- [ ] Successfully extracts the correct text block for a given persona.
- [ ] Performs template substitution for all defined variables.
- [ ] Handles errors gracefully if a persona or variable is missing.

### Tasks
- [ ] Extract persona blocks
- [x] Perform template substitution
