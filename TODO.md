# Doozers Project Roadmap

## Project Summary
The `doozers` project is a Go-based framework (`dv` CLI) designed to orchestrate coding agents (personas like "Cary Coder") through standardized task cycles: **prep**, **make**, and **share**. It aims to accelerate the OODA loop by automating task discovery, environment isolation via Git worktrees, and agent execution.

## Initial Features
- **Persona-based Execution**: A command-line interface (`dv`) that supports different agent personas.
- **Task Management Integration**: Native support for `beans` or `beads` to track tasks and resolve dependencies.
- **Git Worktree Orchestration**: Automated creation and management of isolated worktrees for each task.
- **Scripted Workflows**: Parsing config file based on`SCRIPTS.md` to inject task context into agent prompts.
- **Vendor Agnostic Agent Support**: Initial implementation focusing on Gemini CLI with "YOLO" approval mode.

## Implementation Tasks

### 1. Project Scaffolding
- [ ] Initialize Go module: `go mod init github.com/jmartenstein/doozers`.
- [ ] Install Cobra: `go get github.com/spf13/cobra`.
- [ ] Scaffold `dv` root command and initial subcommands (e.g., `dv run`).

### 2. "Coder" Logic Migration (`run_cary.sh` -> Go)
- [ ] **Bean Client**: Implement a wrapper for `beans query` to:
    - Check if a task is blocked.
    - Recursively follow the blocking chain to find the actionable task.
- [ ] **Git Wrapper**: Implement Go functions to:
    - Identify repository roots.
    - Create/manage Git worktrees (`git worktree add`).
    - Handle branch creation and switching.
- [ ] **Config File**: Migrate `SCRIPTS.md` to a readable maintainable config file that:
    - Can be easily parsed
    - Can be edited and modified either by an agent or by a human.
    - Provides a persona name based on function, for example "Coder"
    - Each provider can also use an alternate alias such as "cary"
- [ ] **Script Parser**: Create a parser for the config file that:
    - Extracts specific persona blocks (e.g., everything under `# Coder / Cary`).
    - Performs template substitution (e.g., replacing `${TASK_ID}`).
- [ ] **Agent Runner**: Implement the execution logic for the Gemini CLI:
    - Change directory to the worktree.
    - Execute `gemini --prompt "..." --approval-mode yolo`.
- [ ] **Outcome Sharing**: Implement the final "share" step:
    - `git add .`, `git commit`, and `git push`.

### 3. CLI Interface Design
- [ ] Implement `dv run <persona> <task_id>` (e.g., `dv run cary <bean_id>`).
- [ ] Support `--dry-run` flag globally.
- [ ] Implement configurable script paths and agent defaults.
