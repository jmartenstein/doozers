---
# do-m2e
title: Implement Bean Client for Task Integration
status: completed
type: task
priority: high
created_at: 2026-02-13T00:27:31Z
updated_at: 2026-02-13T17:18:54Z
parent: do-vul
---

Create a Go wrapper for the `beans` CLI tool to allow the program to query task status, check for blockers, and traverse task dependency chains.

### Acceptance Criteria
- [ ] Can execute `beans query` and parse the JSON/output.
- [ ] Can determine if a task is currently blocked by another task.
- [ ] Can follow a chain of blocking tasks to find the root cause.

### Tasks
- [ ] Implement wrapper for beans query
- [ ] Check if task is blocked
- [x] Follow blocking chain
