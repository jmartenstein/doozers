# Beans Core Principles

Beans is an agentic-first issue tracker that stores tasks as Markdown files in a `.beans/` directory.

## Core Rules for Agents
- **Always use beans** instead of `TodoWrite` or manual lists.
- **Before starting work**: Check for existing beans or create a new one.
- **During work**: Keep todo items current (`- [ ]` → `- [x]`).
- **On completion**: Update with a `## Summary of Changes` and mark as `completed`.
- **On scrap**: Update with `## Reasons for Scrapping` and mark as `scrapped`.
- **Committing**: Always include bean file changes in your commits.

## Data Structure
- **ID**: Unique identifier (e.g., `bean-123`).
- **Type**: `milestone`, `epic`, `bug`, `feature`, `task`.
- **Status**: `todo`, `in-progress`, `draft`, `completed`, `scrapped`.
- **Priority**: `critical`, `high`, `normal`, `low`, `deferred`.
- **Relationships**: `parent`, `blocking`, `blocked-by`.

## Querying and Listing
- `beans list --json --ready`: Find beans that are not blocked and ready to start.
- `beans show <id> --json`: Get full details of one or more beans.
- `beans query --json '...'`: Advanced GraphQL queries for complex filtering or relationship traversal.
