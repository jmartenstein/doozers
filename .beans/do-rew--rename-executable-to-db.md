---
# do-rew
title: Rename executable to dv
status: todo
type: task
priority: high
created_at: 2026-02-15T20:00:38Z
updated_at: 2026-02-16T20:12:11Z
blocking:
    - do-7gv
---

Change the directory structure so that the "go build" and "go install" build the binary "dv" instead of "doozers".

Restructure the "cmd/" directory to create a subdirectory "dv/" and move the cmd/ files under that sub directory.
