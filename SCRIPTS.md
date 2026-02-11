Scripts
=======

Documentation file for scripts for each individual AI command.

# Cary Coder

Start working on feature ${TASK_ID}. Runs "beans prime" and "beans show ${TASK_ID}". Study the parents and child tickets. Work on all subtasks for ${TASK_ID}. Use "beans update" with the "--status" option to mark the status of each ticket as in_progress. Update the beans to mark commands as done. Do not stop working until "${TASK_ID}" is deemed complete. When complete, create a file to be added to the pull request, titled "summary.md".

# Archie Architect

Help me do project planning. First, run "beans prime" to understand how the beans issue tracker works. Then run "beans show silo-rl4" to and review the output to understand the silo-rl4 task. Come up with at least 2 features under this epic, then come up with 3-4 tasks that go under these features. Make sure each feature and epic is associated with the correct parent. Ensure that each feature and task has a detailed description with a checklist.

Run "beans list" to get a full list of all of the beans. Then iterate over each bean, running "beans show" to better understand each item in the list. For each bean not in a "completed" state, analyze which beans have the highest priority, and which beans are blocking other beans.  

First, run "beans update --help" to understand how the "beans update" command works.

Use the "beans update" command with the "--blocking" parameter to denote when one bean is blocking another. Use the "--priority" argument to set the priority to one of the following values: critical, high, normal, low, deferred.

 Write whatever biases and assumptions you find to the file ASSUMPTIONS.md file.
