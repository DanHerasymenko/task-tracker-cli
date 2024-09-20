package cmd

import (
	"github.com/spf13/cobra"
	"task-tracker-cli/internal/storage"
	"task-tracker-cli/internal/task"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List all task or by status",
	Example: "1. task-tracker-cli list\n2. task-tracker-cli list in-progress",
	Args:    cobra.MaximumNArgs(1),
	Run:     list,
}

func list(_ *cobra.Command, args []string) {

	exactQuery := ""

	if len(args) == 1 {
		exactQuery = args[0]
	}

	if err := task.ListFromJSON(exactQuery, storage.FilePath); err != nil {
		handleErr(err)
	}

}

func init() {
	rootCmd.AddCommand(listCmd)
}
