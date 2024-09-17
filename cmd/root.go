package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "task-tracker-cli",
	Short: "CLI Task Tracker",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		handleErr(err)
	}
}

func handleErr(err error) {
	_, _ = fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
