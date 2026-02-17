---
# do-r9r
title: Update git package to handle optional worktrees
status: todo
type: task
priority: high
created_at: 2026-02-17T23:08:35Z
updated_at: 2026-02-17T23:10:08Z
parent: do-4dl
blocking:
    - do-nc4
---

Ensure the 'internal/git' package functions handle cases where a worktree might not be initialized or required.

- [ ] Review internal/git/git.go for assumptions about active worktrees.
- [ ] Ensure 'Cleanup' doesn't fail if no worktree was created.
