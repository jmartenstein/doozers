Doozers
=======
An extensible framework for managing coding agents

# Background

The doozers project provides an extensible framework for orchestrating one or more coding agents to work together. It takes inspiration from the [Gas Town](https://steve-yegge.medium.com/welcome-to-gas-town-4f25ee16dd04) project, as well as the [Ralph Wiggum Loop](https://ghuntley.com/loop/) and the [Outcome Engineering](https://o16g.com/) manifesto. The core of this project is the "dv" utility (short for doozer village) that supports a number of different modes and tools for managing coding agents.

# Overview

The core of the doozers utility is a number of personas that are following scripts to complete tasks. Each script centers around 3 action verbs: prep, make and share.

Personas are named after their primary function, such as "Coder / Cary" for a coding agent, or "Architect / Archie" for a planning and design agent. Each of these verbs will pull from a configurable script file (see the SCRIPTS.md file documentation) that get fed into the agent.

In addition, other shell commands may be run, such as git or beans or beads.

# Requirements

The doozers project is vendor agnostic, able to run on a variety of command-line coding tools, such as Gemini, Code, Claude Caude, goose, etc. The initial implementation will use the Gemini CLI by default.

The "dv" utility is written primarily in the Go programming language. It uses the builds on the [Cobra]() library for managing command-line interfaces.

This project is meant to be used alongside a git-based project tracker like beans or beads. It is designed to be flexible in what tools it uses to track tasks. Future work may include its own minimal task tracker.

The language and structure of this project will be built with [John Boyd's OODA Loop](https://en.wikipedia.org/wiki/OODA_loop) in mind. The goal is not to entirely replace humans in coding, but instead to provide the tools and processes necessary to speed up the OODA loop as much as possible.

# Getting Started

## Build

To build the `dv` utility, ensure you have Go installed (version 1.24.8 or later).

```bash
go build -o dv ./cmd/dv
```

This will create a `dv` executable in the current directory. You can also install it to your `$GOPATH/bin`:

```bash
go install ./cmd/dv
```

## Usage

The `dv` utility uses a command-based interface. The primary command is `run`.

### Running a Persona

Personas are defined in your `.doozers.yaml` file. You can run a persona by name or alias for a specific task:

```bash
# Run using persona name
./dv run Cary do-123

# Run using persona alias
./dv run coder do-123
```

### Options and Flags

- **Dry Run**: See the script that would be executed without actually running it.
  ```bash
  ./dv run Cary do-123 --dry-run
  ```

- **YOLO Mode**: Enable automated execution (skipping confirmations if supported by the agent).
  ```bash
  ./dv run Cary do-123 --yolo
  ```

- **Direct Scripting**: Provide a script directly or via a file instead of using a persona.
  ```bash
  ./dv run do-123 --script "Review the code in main.go and suggest improvements."
  ./dv run do-123 --script-path ./scripts/refactor.txt
  ```

- **Agent Override**: Specify which agent to use (defaults to `gemini` or the persona's configured agent).
  ```bash
  ./dv run Cary do-123 --agent claude
  ```

## Testing

The project includes unit tests for its internal packages. To run all tests:

```bash
go test ./...
```

To run tests with verbosity and coverage:

```bash
go test -v -cover ./...
```

# Project Structure

- `cmd/`: CLI command definitions using Cobra.
- `internal/`: Core logic and internal packages.
  - `beans/`: Integration with the `beans` task tracker.
  - `config/`: Configuration loading and management (`.doozers.yaml`).
  - `git/`: Git operations wrapper.
  - `outcome/`: Logic for sharing task outcomes.
  - `parser/`: Script template parsing and variable substitution.
  - `runner/`: Agent execution logic.
- `.doozers.yaml`: Configuration file for personas and agents.

# Related Documentation

- [SCRIPTS.md](SCRIPTS.md): Detailed documentation on script templates for different personas.
- [OODA.md](docs/OODA.md): Deep dive into the OODA loop philosophy as applied to this project.
- [TODO.md](TODO.md): Current task list and roadmap.
- [GEMINI.md](GEMINI.md): Instructions for using the Gemini CLI within this project.
