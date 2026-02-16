---
# do-xz6
title: Update runner to handle model parameter
status: todo
type: task
priority: normal
created_at: 2026-02-16T20:10:14Z
updated_at: 2026-02-16T20:11:45Z
parent: do-r46
blocking:
    - do-bdf
---

Ensure the runner can accept and pass along the model parameter to the underlying components.

### Checklist
- [ ] Modify runner interface to include model parameter
- [ ] Update call sites to pass the model
