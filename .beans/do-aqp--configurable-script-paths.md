---
# do-aqp
title: Support Configurable Script and Agent Paths
status: completed
type: task
priority: low
created_at: 2026-02-13T00:28:14Z
updated_at: 2026-02-14T00:04:02Z
parent: do-h6e
---

Enhance the CLI to support custom paths for scripts and agent configurations. This allows for greater flexibility across different projects and developer environments.

### Acceptance Criteria
- [ ] Command-line flags are available to override default script paths.
- [ ] The tool correctly falls back to default paths defined in the config file.
- [ ] Validation ensures that specified paths exist and are accessible.

### Tasks
- [x] Support custom script paths
- [x] Support agent defaults

## Summary of Changes\n- Added --script and --script-path flags to override script content.\n- Added --agent flag to override agent executable.\n- Added DefaultAgent to config and Agent to Persona config.
