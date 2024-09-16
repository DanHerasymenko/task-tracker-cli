package cmd

import (
	"github.com/spf13/cobra"
	"task-tracker-cli/internal/errors"
	"task-tracker-cli/internal/storage"
	"task-tracker-cli/internal/task"
)

var addCmd = &cobra.Command{
	Use:     "add",
	Short:   "Add new task",
	Example: "task-tracker-cli add \"Buy groceries\"",
	Args:    cobra.ExactArgs(1),
	Run:     add,
}

func add(_ *cobra.Command, args []string) {

	//check user input
	if len(args) == 0 || args[0] == "" {
		handleErr(errors.NoArgumentsSpecified)
	}
	taskName := args[0]

	//add new task to the json file, create it if it does not exist
	if err := task.Add(taskName, storage.FilePath); err != nil {
		handleErr(err)
	}

}
