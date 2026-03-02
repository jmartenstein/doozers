# Beans Operations: Create and Update

Detailed procedures for creating and modifying beans.

## Create a New Bean
Use `beans create` for any new task, bug, or feature.

### Basic Syntax
```bash
beans create "Title" -t <type> -d "Description..." -s <status>
```

### Advanced Options
- **Relationships**: Use `--parent <id>`, `--blocking <id>`, or `--blocked-by <id>`.
- **Priority**: Use `-p <priority>`.
- **Tags**: Use `--tag <tag-name>`.
- **Prefix**: Use `--prefix <custom-id-prefix>`.

## Update an Existing Bean
Use `beans update` to change status, priority, or body content.

### Metadata Updates
```bash
beans update <id> -s <status> -p <priority> --title "New Title"
```

### Body Content Modification
Beans supports targeted body updates without overwriting the entire file:

**Replace Text (Exact Match):**
```bash
beans update <id> --body-replace-old "old text" --body-replace-new "new text"
```
*Note: This will fail if "old text" is found more than once or not found at all.*

**Append Content:**
```bash
beans update <id> --body-append "New section content"
```

**Combined Update:**
```bash
beans update <id> -s completed --body-replace-old "- [ ] Task" --body-replace-new "- [x] Task"
```

## Optimistic Locking (Etags)
To prevent overwriting someone else's changes:
1. Get the current etag: `ETAG=$(beans show <id> --etag-only)`
2. Update with etag check: `beans update <id> --if-match "$ETAG" ...`

## Multiple Replacements (GraphQL)
For complex updates with multiple replacements and an append:
```bash
beans query 'mutation {
  updateBean(id: "<id>", input: {
    status: "completed"
    bodyMod: {
      replace: [
        { old: "- [ ] Task 1", new: "- [x] Task 1" }
        { old: "- [ ] Task 2", new: "- [x] Task 2" }
      ]
      append: "## Summary

All tasks completed!"
    }
  }) { id body etag }
}'
```
