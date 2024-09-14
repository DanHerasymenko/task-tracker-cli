package task

import (
	"fmt"
	"os"
	"task-tracker-cli/internal/storage"
)

func Add(task string, file *storage.JSONStorage) error {

	if _, err := os.Stat(file.FileName); os.IsNotExist(err) {
		newTask := Task{
			Id:          1,
			Description: task,
			Status:      Todo,
			CreatedAt:   TimeNow,
		}
		// if there is no JSON file create it and write task added by the user
		jsonData, err := WriteToJSON(newTask)
		if err != nil {
			fmt.Println(err)
		}

		errJson := NewJSON(*file, jsonData)
		if err != nil {
			fmt.Println(errJson)
		}
		fmt.Println("Task successfully written to tasks.json")

	}

	return nil

}

//func (t *Tasks) create() error {
//	return nil
//
//}
//
//func (t *Tasks) update() error {
//	return nil
//}
//
//func (t *Tasks) delete() error {
//	return nil
//
//}
//
//func (t *Tasks) list() error {
//	return nil
//
//}
