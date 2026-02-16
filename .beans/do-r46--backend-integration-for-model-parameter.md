---
# do-r46
title: Backend Integration for Model Parameter
status: todo
type: feature
priority: normal
created_at: 2026-02-16T20:10:03Z
updated_at: 2026-02-16T20:10:03Z
parent: do-7gv
---

Integrate the model parameter into the core logic that calls the Gemini API.

### Checklist
- [ ] Update `config` to store default and user-specified models
- [ ] Pass the model parameter through the `runner` to the agent
- [ ] Ensure the agent uses the specified model in its requests
- [ ] Implement fallback logic if a model is not specified
