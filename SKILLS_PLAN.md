# Agent Skills Migration Plan

This plan outlines the steps to adapt the personas defined in `.doozers.yaml` into the official "agent skills" format according to the [specification](https://agentskills.io/specification).

## 1. Directory Structure
We will create a `skills/` directory at the project root to house each unique skill.
```text
skills/
├── coder/
│   └── SKILL.md
└── architect/
    └── SKILL.md
```

## 2. Skill Mapping Strategy

### Cary (The Developer)
- **Directory**: `skills/coder/`
- **Frontmatter**:
    - `name`: `coder`
    - `description`: A skill for executing development tasks using the `beans` utility. Handles task initialization, child ticket management, status updates, and PR summary generation.
    - `metadata`:
        - `aliases`: `["cary", "coder"]`
        - `use_git_worktree`: `true`
- **Instructions**: The `script` content will be moved to the Markdown body, formatted as procedural steps.

### Archie (The Architect)
- **Directory**: `skills/architect/`
- **Frontmatter**:
    - `name`: `architect`
    - `description`: A skill for project planning and backlog management. Categorizes beans (epic, feature, task), manages subtasks, sets priorities, and identifies blocking dependencies.
    - `metadata`:
        - `aliases`: `["archie", "architect"]`
        - `use_git_worktree`: `false`
- **Instructions**: The `script` content will be moved to the Markdown body, formatted as architectural guidelines and planning steps.

## 3. Implementation Steps
1. **Directory Preparation**: Create the root `skills/` parent directory.
2. **Automated Migration Script**: Create a script (e.g., `migrate_to_skills.py`) to handle the conversion of `.doozers.yaml` personas into the `skills/` structure.
   - **Parsing**: Read the `personas` list from the existing `.doozers.yaml`.
   - **Generation**: For each persona, automatically generate the `SKILL.md` with:
     - `name`: Role name (e.g., `coder` for Cary, `architect` for Archie).
     - `description`: A generated summary based on the persona's role.
     - `metadata`: Original properties like `aliases` and `use_git_worktree`.
     - `Body`: The original `script` content, formatted into a clean Markdown block.
   - **Directory Creation**: Create `skills/<name>/` and write the `SKILL.md`.
3. **Deprecation & Configuration Update**: Replace the legacy `personas` block in `.doozers.yaml` with a simplified `skills` reference.

### Example: New `.doozers.yaml`
```yaml
# Updated .doozers.yaml
version: 2.0
skills:
  - path: ./skills/coder
  - path: ./skills/architect
```

4. **Manual Cleanup**: Optionally remove the old `personas` YAML structure once the migration is validated.

## 4. Verification
- Ensure directory names match the `name` field in the frontmatter.
- Validate that all `beans` command instructions are preserved accurately.
- Verify that the `SKILL.md` files follow the Markdown + YAML frontmatter requirement.
