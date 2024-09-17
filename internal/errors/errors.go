package errors

import "errors"

var (
	NoArgumentsSpecified = errors.New("No arguments specified")
	InvalidNumOfArgs     = errors.New("Invalid number of arguments")
	InvalidTaskID        = errors.New("Invalid task ID")

	//ErrorCreatingJSONFile = errors.New("there is an error when creatin a new JSON file")
)
