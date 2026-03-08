# GitHub CLI (gh) Help Reference

## gh issue

### gh issue list
List issues in a GitHub repository.
```
USAGE
  gh issue list [flags]

FLAGS
  -a, --assignee string    Filter by assignee
  -A, --author string      Filter by author
  -l, --label strings      Filter by label
  -L, --limit int          Maximum number of issues to fetch (default 30)
  -S, --search query       Search issues with query
  -s, --state string       Filter by state: {open|closed|all} (default "open")
```

### gh issue view
Display information about an issue.
```
USAGE
  gh issue view {<number> | <url>} [flags]

FLAGS
  -c, --comments          View issue comments
  -w, --web               Open an issue in the browser
```

### gh issue edit
Edit one or more issues.
```
USAGE
  gh issue edit {<numbers> | <urls>} [flags]

FLAGS
      --add-assignee login      Add assigned users by their login.
      --add-label name          Add labels by name
  -b, --body string             Set the new body.
  -t, --title string            Set the new title.
      --remove-label name       Remove labels by name
```

---

## gh pr

### gh pr create
Create a pull request on GitHub.
```
USAGE
  gh pr create [flags]

FLAGS
  -B, --base branch          The branch into which you want your code merged
  -b, --body string          Body for the pull request
  -d, --draft                Mark pull request as a draft
  -f, --fill                 Use commit info for title and body
  -t, --title string         Title for the pull request
  -w, --web                  Open the web browser to create a pull request
```

### gh pr ready
Mark a pull request as ready for review.
```
USAGE
  gh pr ready [<number> | <url> | <branch>] [flags]
```

### gh pr edit
Edit a pull request.
```
USAGE
  gh pr edit [<number> | <url> | <branch>] [flags]

FLAGS
      --add-assignee login      Add assigned users by their login.
      --add-label name          Add labels by name
      --add-reviewer login      Add reviewers by their login.
  -b, --body string             Set the new body.
  -t, --title string            Set the new title.
```
