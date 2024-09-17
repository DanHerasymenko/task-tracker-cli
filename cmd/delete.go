package cmd

import (
	"github.com/spf13/cobra"
	"task-tracker-cli/internal/errors"
)

var deleteCmd = &cobra.Command{
	Use:     "delete",
	Short:   "Delete a task by ID",
	Example: "task-tracker-cli update 1 \"Buy groceries and cook dinner\"",
	Args:    cobra.ExactArgs(1),
	Run:     deleteTask,
}

func deleteTask(_ *cobra.Command, args []string) {

	//check user input
	if len(args) == 0 || args[0] == "" {
		handleErr(errors.NoArgumentsSpecified)
	} else if len(args) != 1 {
		handleErr(errors.InvalidNumOfArgs)
	}

}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
