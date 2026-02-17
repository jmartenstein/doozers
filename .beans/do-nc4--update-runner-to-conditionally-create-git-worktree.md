---
# do-nc4
title: Update runner to conditionally create git worktrees
status: todo
type: task
priority: high
created_at: 2026-02-17T23:08:30Z
updated_at: 2026-02-17T23:10:13Z
parent: do-4dl
---

Modify the execution flow to skip git worktree setup and cleanup if 'use_git_worktree' is set to false for the active persona.

- [ ] Update internal/runner/runner.go to check 'UseGitWorktree' before calling git worktree functions.
- [ ] Ensure that even without a worktree, the persona can still execute Gemini commands.
- [ ] Add unit tests for conditional execution.
