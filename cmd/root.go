package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
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
