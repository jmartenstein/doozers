# Doozers Project Roadmap

## Project Summary
The `doozers` project is a Go-based framework (`dv` CLI) designed to orchestrate coding agents (personas like "Cary Coder") through standardized task cycles: **prep**, **make**, and **share**. It aims to accelerate the OODA loop by automating task discovery, environment isolation via Git worktrees, and agent execution.

## Initial Features
- **Persona-based Execution**: A command-line interface (`dv`) that supports different agent personas.
- **Task Management Integration**: Native support for `beans` or `beads` to track tasks and resolve dependencies.
- **Git Worktree Orchestration**: Automated creation and management of isolated worktrees for each task.
- **Scripted Workflows**: Parsing config file based on `SCRIPTS.md` to inject task context into agent prompts.
- **Vendor Agnostic Agent Support**: Initial implementation focusing on Gemini CLI with "YOLO" approval mode.
- **Adapt 3-Step Persona**: Modular task cycles via configurable `prep`, `make`, and `share` steps for all personas.

## Implementation Tasks

### 1. Project Scaffolding
- [x] **Initialize Go module**: `go mod init github.com/jmartenstein/doozers`.
- [x] **Install Cobra**: `go get github.com/spf13/cobra`.
- [x] **Scaffold `dv` root command**: Implement initial subcommands (e.g., `dv run`).

### 2. "Coder" Logic Migration (Initial Implementation)
- [x] **Bean Client**: Implement a wrapper for `beans query` to check task status and find root causes.
- [x] **Git Wrapper**: Implement Go functions for repo root identification and worktree management.
- [x] **Config File**: Migrate `SCRIPTS.md` to a maintainable `.doozers.yaml` configuration.
- [x] **Script Parser**: Create a parser for template substitution (e.g., replacing `{{TASK_ID}}`).
- [x] **Agent Runner**: Implement initial execution logic for the Gemini CLI with YOLO mode.
- [x] **Outcome Sharing**: Implement the hardcoded "share" step (git add/commit/push).

### 3. CLI Interface Design
- [ ] **Refine `dv run`**: Finalize `dv run <persona> <task_id>` command structure.
- [ ] **Dry Run Support**: Implement a global `--dry-run` flag to preview actions without execution.
- [ ] **Configurable Paths**: Implement support for custom script paths and agent defaults.

### 4. Adapt to 3-Step Persona Paradigm (Prep, Make, Share)
- [ ] **Update Configuration Model**
    - Modify `internal/config/config.go` to update the `Persona` struct with `Prep`, `Make`, and `Share` fields.
    - Support the existing `Script` field as a fallback for the `Make` step for backward compatibility.
    - Update `Config.Load` to parse the new modular schema in `.doozers.yaml`.
- [ ] **Implement General Shell Runner**
    - Refactor `internal/runner/runner.go` to execute arbitrary shell scripts instead of hardcoded `gemini` commands.
    - Ensure multi-line script support and direct output streaming to `os.Stdout` and `os.Stderr`.
    - Maintain `Yolo` flag functionality through environment variables or template substitution.
- [ ] **Enhance Variable Substitution**
    - Update `internal/parser/parser.go` to ensure `{{TASK_ID}}` and context are substituted across all three steps.
- [ ] **Orchestrate Execution in `dv run`**
    - Update `cmd/run.go` to execute `Prep`, `Make`, and `Share` steps in sequence.
    - Implement abort-on-failure logic: if any step fails, the entire sequence is halted.
    - Manage working directory context (e.g., allowing a `Prep` step to set the environment for `Make`).
- [ ] **Refactor Outcome Sharing**
    - Remove hardcoded `git add`/`commit` logic from `internal/outcome/outcome.go`.
    - Delegate sharing responsibility to the configurable `Share` step in `.doozers.yaml`.
- [ ] **Migrate `.doozers.yaml`**
    - Update existing persona definitions (Cary, Archie) to follow the new 3-step modular paradigm.

#### Example of Future `.doozers.yaml`

```yaml
personas:
  - name: Cary
    aliases: ["cary", "coder"]
    # Prep: Set up the environment, e.g., create a worktree
    prep: |
      git worktree add -b task/{{TASK_ID}} ../doozers-{{TASK_ID}} main
    # Make: The core agent execution. Note: cd into the worktree if needed.
    make: |
      cd ../doozers-{{TASK_ID}} && gemini chat -p "Start working on feature {{TASK_ID}}. ..."
    # Share: Commit changes and push
    share: |
      cd ../doozers-{{TASK_ID}} && git add . && git commit -m "Completed {{TASK_ID}}" && git push origin task/{{TASK_ID}}

  - name: Archie
    aliases: ["archie", "architect"]
    prep: |
      beans prime
    make: |
      gemini chat -p "Help me do project planning for {{TASK_ID}}..."
    share: |
      git add ASSUMPTIONS.md && git commit -m "Archie: Planning for {{TASK_ID}}"
```
