---
# do-rew
title: Rename executable to dv
status: completed
type: task
priority: critical
created_at: 2026-02-15T20:00:38Z
updated_at: 2026-02-17T23:42:20Z
blocking:
    - do-7gv
---

Change the directory structure so that the "go build" and "go install" build the binary "dv" instead of "doozers".

Restructure the "cmd/" directory to create a subdirectory "dv/" and move the cmd/ files under that sub directory.

## Summary of Changes

- Created `cmd/dv/` directory.
- Moved `main.go` to `cmd/dv/main.go` and updated it to be part of the `main` package.
- Moved `cmd/root.go` and `cmd/run.go` to `cmd/dv/` and updated them to be part of the `main` package.
- Verified that `go build ./cmd/dv` produces the `dv` executable.
- The binary is now named `dv` when built or installed from `cmd/dv/`.
