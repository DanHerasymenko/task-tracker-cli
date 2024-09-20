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
	} else if len(args) != 1 {
		handleErr(errors.InvalidNumOfArgs)
	}
	taskName := args[0]

	newTask := task.Task{
		Id:          task.IncrementId(storage.FilePath),
		Description: taskName,
		Status:      task.Todo,
		CreatedAt:   task.TimeNow,
	}

	if err := task.WriteToJSON(newTask, storage.FilePath); err != nil {
		handleErr(err)
	}

}

func init() {
	rootCmd.AddCommand(addCmd)
}
