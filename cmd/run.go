package cmd

import (
	"fmt"
	"os"

	"github.com/jmartenstein/doozers/internal/beans"
	"github.com/jmartenstein/doozers/internal/config"
	"github.com/jmartenstein/doozers/internal/git"
	"github.com/jmartenstein/doozers/internal/outcome"
	"github.com/jmartenstein/doozers/internal/parser"
	"github.com/jmartenstein/doozers/internal/runner"
	"github.com/spf13/cobra"
)

var (
	yolo bool
)

var runCmd = &cobra.Command{
	Use:   "run [persona] [task-id]",
	Short: "Run a script for a persona",
	Long:  `Run a script for a persona (e.g., Cary, Archie) for a specific task.`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		personaName := args[0]
		taskID := args[1]

		cfg, err := config.Load(".doozers.yaml")
		if err != nil {
			fmt.Printf("Error loading config: %v\n", err)
			os.Exit(1)
		}

		persona, ok := cfg.GetPersona(personaName)
		if !ok {
			fmt.Printf("Persona not found: %s\n", personaName)
			os.Exit(1)
		}

		fmt.Printf("Checking status for task %s...\n", taskID)
		bClient := beans.NewClient()
		blocked, err := bClient.IsBlocked(taskID)
		if err != nil {
			fmt.Printf("Error checking task status: %v\n", err)
			os.Exit(1)
		}
		if blocked {
			rootCause, err := bClient.GetRootCause(taskID)
			if err != nil {
				fmt.Printf("Task %s is blocked, but could not find root cause: %v\n", taskID, err)
			} else {
				fmt.Printf("Task %s is blocked by: %s (%s)\n", taskID, rootCause.Title, rootCause.ID)
			}
			os.Exit(1)
		}

		vars := map[string]string{
			"TASK_ID": taskID,
		}
		script := parser.Substitute(persona.Script, vars)

		fmt.Printf("Running script for persona %s on task %s...\n", persona.Name, taskID)

		r := runner.NewRunner(yolo)
		if err := r.Run(script, "."); err != nil {
			fmt.Printf("Error running agent: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("Agent execution completed successfully. Sharing outcome...")
		g := git.NewGit()
		m := outcome.NewManager(g)
		if err := m.Share(taskID, "Completed task using persona "+persona.Name); err != nil {
			fmt.Printf("Error sharing outcome: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("Outcome shared successfully.")
	},
}

func init() {
	runCmd.Flags().BoolVar(&yolo, "yolo", false, "Enable YOLO mode (automated execution)")
	rootCmd.AddCommand(runCmd)
}
