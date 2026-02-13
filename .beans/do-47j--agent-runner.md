---
# do-47j
title: Implement Agent Runner for Task Execution
status: draft
type: task
priority: normal
created_at: 2026-02-13T00:27:51Z
updated_at: 2026-02-13T00:31:09Z
parent: do-vul
---

Build the logic to execute the Gemini CLI agent. The runner must manage the agent's environment, handle directory changes (especially when using worktrees), and support "YOLO" mode for automated execution.

### Acceptance Criteria
- [ ] Successfully invokes the Gemini CLI with the correct persona and task context.
- [ ] Correctly manages the working directory within git worktrees.
- [ ] Supports a mode where agent actions are executed without manual confirmation.

### Tasks
- [ ] Execute Gemini CLI
- [ ] Handle worktree directory changes
- [ ] Set YOLO mode
