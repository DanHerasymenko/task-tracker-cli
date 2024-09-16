package task

func Add(task string, file string) error {

	newTask := Task{
		Id:          1,
		Description: task,
		Status:      Todo,
		CreatedAt:   TimeNow,
	}

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
