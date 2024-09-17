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
		log.Fatal("Error unmarshaling JSON: %w", errUnmarshal)
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
			return fmt.Errorf("Error reading file: %w", err)
		}
	}

	// Unmarshal the existing content into the `tasks` slice
	errUnmarshal := json.Unmarshal(data, &tasks)
	if errUnmarshal != nil && len(data) > 0 { // Avoid error on empty file
		return fmt.Errorf("Error unmarshaling JSON: %w", errUnmarshal)
	}

	// Append the new task
	tasks = append(tasks, newTask)

	// Marshal updated tasks list to JSON
	result, errorMar := json.MarshalIndent(tasks, "", "  ")
	if errorMar != nil {
		return fmt.Errorf("Error marshaling the file: %w", errorMar)
	}

	// Write to file and create it if not exists
	errFile := os.WriteFile(path, result, 0644)
	if errFile != nil {
		return fmt.Errorf("Error writing or creating JSON file: %w", errFile)
	}

	fmt.Println("Task successfully written to tasks.json")
	return nil
}
