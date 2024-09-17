package cmd

import (
	"github.com/spf13/cobra"
	"strconv"
	"task-tracker-cli/internal/errors"
	"task-tracker-cli/internal/storage"
	"task-tracker-cli/internal/task"
)

var updateCmd = &cobra.Command{
	Use:     "update",
	Short:   "Update a task by ID",
	Example: "task-tracker-cli update 1 \"Buy groceries and cook dinner\"",
	Args:    cobra.ExactArgs(2),
	Run:     update,
}

func update(_ *cobra.Command, args []string) {

	//check user input
	if len(args) == 0 || args[0] == "" {
		handleErr(errors.NoArgumentsSpecified)
	} else if len(args) != 2 {
		handleErr(errors.InvalidNumOfArgs)
	}

	TaskIdString := args[0]
	NewTaskName := args[1]

	//convert Task ID argument to int and handle error
	TaskId, err := strconv.Atoi(TaskIdString)
	if err != nil {
		handleErr(errors.InvalidTaskID)
	}

	// search a task in json file, return error if no file -> create a new file
	// update task by ID -> return error if task ID not found -> propose to list all tasks
	if err := task.UpdateTask(NewTaskName, storage.FilePath, TaskId); err != nil {
		handleErr(err)
	}

}

func init() {
	rootCmd.AddCommand(updateCmd)
}
