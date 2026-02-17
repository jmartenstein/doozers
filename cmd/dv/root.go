package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	dryRun bool
)

var rootCmd = &cobra.Command{
	Use:   "dv",
	Short: "dv is a CLI tool for managing doozers",
	Long:  `dv is a CLI tool for managing doozers, built with Cobra.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&dryRun, "dry-run", false, "Simulate actions without making any changes")
}
