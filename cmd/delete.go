package cmd

import (
	"github.com/spf13/cobra"
	"strconv"
	"task-tracker-cli/internal/errors"
	"task-tracker-cli/internal/storage"
	"task-tracker-cli/internal/task"
)

var deleteCmd = &cobra.Command{
	Use:     "delete",
	Short:   "Delete a task by ID",
	Example: "task-tracker-cli delete 1 ",
	Args:    cobra.ExactArgs(1),
	Run:     deleteTask,
}

func deleteTask(cmd *cobra.Command, args []string) {

	//check user input
	if len(args) == 0 || args[0] == "" {
		handleErr(errors.NoArgumentsSpecified)
	} else if len(args) != 1 {
		handleErr(errors.InvalidNumOfArgs)
	}

	TaskIdString := args[0]

	//convert Task ID argument to int and handle error
	TaskId, err := strconv.Atoi(TaskIdString)
	if err != nil {
		handleErr(errors.InvalidTaskID)
	}

	//delete task
	if err := task.UpdateJSON(cmd.Use, "", storage.FilePath, TaskId); err != nil {
		handleErr(err)
	}
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
