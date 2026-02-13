---
# do-7d7
title: Support '--dry-run' Flag for Safe Execution
status: draft
type: task
priority: low
created_at: 2026-02-13T00:28:09Z
updated_at: 2026-02-13T00:31:32Z
parent: do-h6e
---

Add a global `--dry-run` flag to the `dv` CLI. This flag should prevent any actual changes from being made (e.g., no git commits, no agent execution) while still logging what would have happened.

### Acceptance Criteria
- [ ] The `--dry-run` flag is recognized globally across all `dv` commands.
- [ ] When enabled, all destructive operations are bypassed and logged instead.
- [ ] The user receives clear feedback that they are in dry-run mode.

### Tasks
- [ ] Add global `--dry-run` flag
