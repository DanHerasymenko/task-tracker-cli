package task

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"task-tracker-cli/internal/storage"
)

func NewJSON(file storage.JSONStorage, jsonData []byte) error {

	// Construct the path to the internal storage directory
	path := filepath.Join("internal", "storage", file.FileName)

	err := os.WriteFile(path, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("Error writing file: %w", err)
	}
	return nil
}

func WriteToJSON(newTask Task) ([]byte, error) {
	jsonData, err := json.MarshalIndent(newTask, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("Error marshaling JSON: %w", err)
	}
	return jsonData, nil
}
