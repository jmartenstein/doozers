package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a script for a persona",
	Long:  `Run a script for a persona (e.g., Cary, Archie).`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run command called")
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
