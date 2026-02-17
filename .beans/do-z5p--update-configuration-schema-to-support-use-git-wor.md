---
# do-z5p
title: Update configuration schema to support use_git_worktree for personas
status: todo
type: task
priority: high
created_at: 2026-02-17T23:08:25Z
updated_at: 2026-02-17T23:10:03Z
parent: do-4dl
blocking:
    - do-nc4
    - do-a0n
---

Update the configuration parsing logic and types to include a 'use_git_worktree' boolean field for each persona in '.doozers.yaml'.

- [ ] Update internal/config/config.go to include 'UseGitWorktree' in persona struct.
- [ ] Update parser logic if necessary.
- [ ] Add tests for new configuration field.
