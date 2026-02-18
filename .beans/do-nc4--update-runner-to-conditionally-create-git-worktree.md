---
# do-nc4
title: Update runner to conditionally create git worktrees
status: completed
type: task
priority: high
created_at: 2026-02-17T23:08:30Z
updated_at: 2026-02-18T00:30:58Z
parent: do-4dl
---

Modify the execution flow to skip git worktree setup and cleanup if 'use_git_worktree' is set to false for the active persona.

- [ ] Update internal/runner/runner.go to check 'UseGitWorktree' before calling git worktree functions.
- [ ] Ensure that even without a worktree, the persona can still execute Gemini commands.
- [x] Add unit tests for conditional execution.
