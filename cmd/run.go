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
	yolo       bool
	scriptStr  string
	scriptPath string
	agent      string
)

var runCmd = &cobra.Command{
	Use:   "run [persona] [task-id]",
	Short: "Run a script for a persona",
	Long:  `Run a script for a persona (e.g., Cary, Archie) for a specific task.`,
	Args:  cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		var personaName string
		var taskID string

		if len(args) == 2 {
			personaName = args[0]
			taskID = args[1]
		} else {
			taskID = args[0]
		}

		vars := map[string]string{
			"TASK_ID": taskID,
		}

		var scriptContent string
		var selectedAgent string

		cfg, err := config.Load(".doozers.yaml")
		if err != nil {
			// It's okay if config doesn't exist if we provide script directly
			if scriptStr == "" && scriptPath == "" {
				fmt.Printf("Error loading config: %v\n", err)
				os.Exit(1)
			}
		}

		if agent != "" {
			selectedAgent = agent
		} else if cfg != nil && cfg.DefaultAgent != "" {
			selectedAgent = cfg.DefaultAgent
		} else {
			selectedAgent = "gemini"
		}

		if scriptStr != "" {
			scriptContent = scriptStr
		} else if scriptPath != "" {
			data, err := os.ReadFile(scriptPath)
			if err != nil {
				fmt.Printf("Error reading script file: %v\n", err)
				os.Exit(1)
			}
			scriptContent = string(data)
		} else {
			if personaName == "" {
				fmt.Println("Persona name is required when --script or --script-path is not provided")
				os.Exit(1)
			}
			if cfg == nil {
				fmt.Println("Config file .doozers.yaml is required when using persona")
				os.Exit(1)
			}

			persona, ok := cfg.GetPersona(personaName)
			if !ok {
				fmt.Printf("Persona not found: %s\n", personaName)
				os.Exit(1)
			}
			scriptContent = persona.Script
			if agent == "" && persona.Agent != "" {
				selectedAgent = persona.Agent
			}
		}

		if taskID == "" {
			fmt.Println("Task ID is required")
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

		script := parser.Substitute(scriptContent, vars)

		if dryRun {
			fmt.Println("DRY RUN: The following script would be executed:")
			fmt.Println("--------------------------------")
			fmt.Println(script)
			fmt.Println("--------------------------------")
			fmt.Printf("Agent: %s, YOLO: %v\n", selectedAgent, yolo)
			fmt.Println("Outcome sharing would be triggered after successful execution.")
			return
		}

		fmt.Printf("Running script for task %s...\n", taskID)

		r := runner.NewRunner(yolo)
		r.Executable = selectedAgent

		if err := r.Run(script, "."); err != nil {
			fmt.Printf("Error running agent: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("Agent execution completed successfully. Sharing outcome...")
		g := git.NewGit()
		m := outcome.NewManager(g)
		if err := m.Share(taskID, "Completed task"); err != nil {
			fmt.Printf("Error sharing outcome: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("Outcome shared successfully.")
	},
}

func init() {
	runCmd.Flags().BoolVar(&yolo, "yolo", false, "Enable YOLO mode (automated execution)")
	runCmd.Flags().StringVar(&scriptStr, "script", "", "Provide the script content directly")
	runCmd.Flags().StringVar(&scriptPath, "script-path", "", "Path to a file containing the script")
	runCmd.Flags().StringVar(&agent, "agent", "gemini", "Override the agent executable")
	rootCmd.AddCommand(runCmd)
}
