package task

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func IncrementId(path string) int {
	var tasks []Task

	data, err := os.ReadFile(path)

	if err != nil {
		if os.IsNotExist(err) {
			// If file doesn't exist, return id as 1
			return 1
		} else {
			log.Fatal("Error reading file: %w", err)
		}
	}

	// Unmarshal the existing content into the `tasks` slice
	errUnmarshal := json.Unmarshal(data, &tasks)
	if errUnmarshal != nil && len(data) > 0 { // Avoid error on empty file
		log.Fatal("Error unmarshalling JSON: %w", errUnmarshal)
	}

	lastId := tasks[len(tasks)-1].Id // Get the last element from the array and retrieve the ID of the last task
	return lastId + 1
}

func WriteToJSON(newTask Task, path string) error {
	var tasks []Task

	// Read the existing JSON file content
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			// If file doesn't exist, initialize an empty task list
			tasks = []Task{}
		} else {
			return fmt.Errorf("error reading file: %w", err)
		}
	}

	// Unmarshal the existing content into the `tasks` slice
	errUnmarshal := json.Unmarshal(data, &tasks)
	if errUnmarshal != nil && len(data) > 0 { // Avoid error on empty file
		return fmt.Errorf("error unmarshaling JSON: %w", errUnmarshal)
	}

	// Append the new task
	tasks = append(tasks, newTask)

	// Marshal updated tasks list to JSON
	result, errorMar := json.MarshalIndent(tasks, "", "  ")
	if errorMar != nil {
		return fmt.Errorf("error marshaling the file: %w", errorMar)
	}

	// Write to file and create it if not exists
	errFile := os.WriteFile(path, result, 0644)
	if errFile != nil {
		return fmt.Errorf("error writing or creating JSON file: %w", errFile)
	}

	return nil
}

func UpdateTask(taskName string, path string, id int) error {

	var tasks []Task

	data, err := os.ReadFile(path)

	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("JSON file does not exists")
		} else {
			return fmt.Errorf("error reading file: %w", err)
		}
	}

	// Unmarshal the existing content into the `tasks` slice
	errUnmarshal := json.Unmarshal(data, &tasks)
	if errUnmarshal != nil && len(data) > 0 { // Avoid error on empty file
		log.Fatal("Error unmarshalling JSON: %w", errUnmarshal)
	}

	//Check for task ID and update TaskName
	if id < 1 || id > len(tasks) {
		return fmt.Errorf("invalid task ID: %d, try to list all tasks", id)
	}
	tasks[id-1].Description = taskName
	tasks[id-1].UpdatedAt = TimeNow
	fmt.Printf("Task (ID: %d) updated successfully\n", id)

	// Marshal updated tasks list to JSON
	result, errorMar := json.MarshalIndent(tasks, "", "  ")
	if errorMar != nil {
		return fmt.Errorf("error marshaling the file: %w", errorMar)
	}

	// Write to file and create it if not exists
	errFile := os.WriteFile(path, result, 0644)
	if errFile != nil {
		return fmt.Errorf("error writing to the JSON file %w", errFile)
	}

	return nil
}
