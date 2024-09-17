package task

import "fmt"

func Add(task string, file string) error {

	newTask := Task{
		Id:          IncrementId(file),
		Description: task,
		Status:      Todo,
		CreatedAt:   TimeNow,
	}
	fmt.Println(newTask)
	if err := WriteToJSON(newTask, file); err != nil {
		return err
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
