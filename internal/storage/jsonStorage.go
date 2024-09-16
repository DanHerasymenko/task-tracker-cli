package storage

import "path/filepath"

const FileName = "tasks.json"

var FilePath = filepath.Join("internal", "storage", FileName)

//type JSONStorage struct {
//	FilePath string
//}
//
//var JsonStore = JSONStorage{FilePath: StoragePath}
