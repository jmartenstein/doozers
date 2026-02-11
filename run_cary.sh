#!/bin/bash
set -e

DRY_RUN=false
BEAN_ID=""

# Parse arguments
while [[ $# -gt 0 ]]; do
  case $1 in
    --dry-run|-d)
      DRY_RUN=true
      shift
      ;;
    *)
      if [ -z "$BEAN_ID" ]; then
        BEAN_ID=$1
      fi
      shift
      ;;
  esac
done

# Check if bean ID is provided
if [ -z "$BEAN_ID" ]; then
    echo "Usage: $0 [--dry-run|-d] <bean_id>"
    exit 1
fi

# Capture the directory where the script is located to find SCRIPTS.md later
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
DOOZERS_ROOT=$(git -C "$SCRIPT_DIR" rev-parse --show-toplevel)

# Ensure we are running from the target repository root (the one we are in)
REPO_ROOT=$(git rev-parse --show-toplevel)
cd "$REPO_ROOT"

# 1. Check if bean and children are blocked
echo "Checking if $BEAN_ID and its children are blocked..."
QUERY='{
  bean(id: "'$BEAN_ID'") {
    id
    blockedBy(filter: { excludeStatus: ["completed"] }) { id }
    children {
      id
      blockedBy(filter: { excludeStatus: ["completed"] }) { id }
    }
  }
}'

RESULT=$(beans query --json "$QUERY")

# Check if bean exists
if [ "$(echo "$RESULT" | jq '.bean')" == "null" ]; then
    echo "Error: Bean $BEAN_ID not found."
    exit 1
fi

# Extract any blockers from the bean itself or its children
BLOCKERS=$(echo "$RESULT" | jq -r '
  [.bean.blockedBy[].id, (.bean.children[].blockedBy[].id // empty)] | unique | join(", ")
' | sed 's/^, //; s/ ,$//')

if [ -n "$BLOCKERS" ]; then
    echo "Error: Cannot proceed. Bean $BEAN_ID or its children are blocked by: $BLOCKERS"
    exit 1
fi

# 2. Handle git worktree and branch
REPO_NAME=$(basename "$REPO_ROOT")
BRANCH_NAME="${REPO_NAME}-${BEAN_ID}"
# Create worktree in a sibling directory
WORKTREE_PATH="$REPO_ROOT/../${BRANCH_NAME}"

if [ "$DRY_RUN" = true ]; then
    if [ -d "$WORKTREE_PATH" ]; then
        echo "[DRY RUN] Worktree directory '$WORKTREE_PATH' already exists."
    elif git rev-parse --verify "$BRANCH_NAME" >/dev/null 2>&1; then
        echo "[DRY RUN] Branch '$BRANCH_NAME' already exists. Would add worktree at '$WORKTREE_PATH' using existing branch."
    else
        echo "[DRY RUN] Would create git worktree for branch '$BRANCH_NAME' at '$WORKTREE_PATH'"
    fi
else
    if [ -d "$WORKTREE_PATH" ]; then
        echo "Worktree directory '$WORKTREE_PATH' already exists. Skipping creation."
    elif git rev-parse --verify "$BRANCH_NAME" >/dev/null 2>&1; then
        echo "Branch '$BRANCH_NAME' already exists. Adding worktree at '$WORKTREE_PATH' using existing branch..."
        git worktree add "$WORKTREE_PATH" "$BRANCH_NAME"
    else
        echo "Creating git worktree for branch '$BRANCH_NAME' at '$WORKTREE_PATH'..."
        git worktree add -b "$BRANCH_NAME" "$WORKTREE_PATH"
    fi
fi

# 4. Run gemini headless with "Cary Coder" script
# Extract the script content from SCRIPTS.md and replace TASK_ID
SCRIPT_CONTENT=$(awk '/# Cary Coder/{flag=1; next} /^# /{flag=0} flag' "$DOOZERS_ROOT/SCRIPTS.md" | sed "s/\${TASK_ID}/$BEAN_ID/g" | tr '\n' ' ' | sed 's/  */ /g; s/^ //; s/ $//')

if [ -z "$SCRIPT_CONTENT" ]; then
    echo "Error: Could not extract 'Cary Coder' script from SCRIPTS.md"
    exit 1
fi

if [ "$DRY_RUN" = true ]; then
    echo "[DRY RUN] Would run Gemini in headless mode for task $BEAN_ID"
    echo "[DRY RUN] Script content: $SCRIPT_CONTENT"
else
    echo "Running Gemini in headless mode for task $BEAN_ID..."
    # Execute Gemini from within the worktree
    (cd "$WORKTREE_PATH" && gemini --prompt "$SCRIPT_CONTENT" --approval-mode yolo)
fi

# 5. Commit and push changes
if [ "$DRY_RUN" = true ]; then
    echo "[DRY RUN] Would commit and push changes to remote"
else
    echo "Committing and pushing changes to remote..."
    (cd "$WORKTREE_PATH" && 
        git add . && 
        git commit -m "Cary Coder: Implementation for $BEAN_ID" && 
        git push origin "$BRANCH_NAME")
fi

if [ "$DRY_RUN" = true ]; then
    echo "[DRY RUN] Process completed successfully for $BEAN_ID."
else
    echo "Process completed successfully for $BEAN_ID."
fi
