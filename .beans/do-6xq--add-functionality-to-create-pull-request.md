---
# do-6xq
title: Add functionality to create pull request
status: todo
type: feature
priority: normal
created_at: 2026-02-16T17:51:06Z
updated_at: 2026-02-16T20:11:55Z
---

In its current state, the "dv" utility creates a branch and pushes it to origin. 
Once the commit is pushed to the branch, the dv utility should also create the 
pull request, and use the file "PR_SUMMARY.md" as the pull request's description.

Once the pull request is confirmed completed, the PR_SUMMARY.md file can be
deleted
