---
# do-4dl
title: Configure which personas need git worktrees
status: draft
type: task
priority: normal
created_at: 2026-02-17T22:58:51Z
updated_at: 2026-02-17T23:00:42Z
---

In the current ".doozers.yaml" file, there are two personas. Only the "cary / coder" persona needs to create a git worktree. The "archie / architect" persona does not need a git work tree and should not push changes to a git branch.

Add functionality to the existing config and git settings so that the the git steps are optional, and a persona can run ONLY a gemini command without creating a git worktree or pushing to a branch.
