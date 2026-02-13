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
