# Project Planning Assumptions

- **Epic do-7gv**: The "model parameter" task (do-7gv) is treated as an epic that encompasses both CLI (Feature 1) and Backend (Feature 2) changes.
- **Priority**: Task do-rew (Rename executable to dv) is considered critical priority as it defines the CLI tool's identity and should ideally be completed before adding new features.
- **CLI-First**: We assume a CLI-first approach where the flag and its validation (Feature 1) are defined before the backend implementation (Feature 2) is finalized.
- **Configuration**: The model parameter (--model) is assumed to be stored in the configuration object and passed through the runner to the Gemini API client.
- **Model Support**: The backend implementation assumes that the Gemini API supports selecting specific models by name.
- **Feature do-4dl**: The task "Configure which personas need git worktrees" is treated as a feature that affects multiple components: configuration schema, runner execution, and git package handling.
- **Git Worktree Optionality**: We assume that certain personas (like `archie / architect`) only need to interact with the LLM and do not require a separate git worktree or branch for their operations.
- **Doozers.yaml Integration**: We assume the configuration for git worktree necessity belongs in the persona-specific configuration within `.doozers.yaml`.
- **CLI Override**: We assume that `do-8k0` (--no-worktree option) will provide a manual override for the configuration-based setting.
- **Tool Renaming**: We assume that renaming the tool to `dv` (do-rew) is a prerequisite for all other work to ensure consistent CLI branding and help documentation.
