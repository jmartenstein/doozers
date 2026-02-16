---
# do-ahe
title: CLI Support for Model Selection
status: todo
type: feature
priority: normal
created_at: 2026-02-16T20:09:23Z
updated_at: 2026-02-16T20:11:50Z
parent: do-7gv
blocking:
    - do-r46
---

Implement the CLI interface for selecting the Gemini model. This includes adding the flag to relevant commands and validating the input.

### Checklist
- [ ] Add `--model` flag to `run` command
- [ ] Implement validation for the model flag
- [ ] Add documentation for the model flag in `--help`
- [ ] Unit tests for CLI flag parsing
