package storage

type JSONStorage struct {
	FileName string
}

var JsonStore = JSONStorage{FileName: "tasks.json"}
