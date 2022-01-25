package file

import (
	"errors"
	"fmt"
	"os"
	"sync"
)

var lock = &sync.Mutex{}

type File struct {
	Name   string
	Folder string
	Path   string
}

func (f *File) IsFileExists() bool {
	if _, err := os.Stat(f.Path); err == nil {
		return true
	}
	return false
}

func (f *File) IsFolderExists() bool {
	if _, err := os.Stat(f.Name); err == nil {
		return true
	}
	return false
}

func (f *File) CreateFile() error {
	_, err := os.Create(f.Path)
	if err != nil {
		return errors.New("Dosya Oluşturulurken Bir Hata Meydana Geldi")
	}
	return nil
}

func (f *File) CreateFolder() error {
	err := os.Mkdir(f.Folder, 0755)
	if err != nil {
		return errors.New("Klasor Olusturulurken Hata Meydana Geldi")
	}
	return nil
}

var fileInstance *File

func GetInstance() *File {
	lock.Lock()
	defer lock.Unlock()
	if fileInstance == nil {
		fmt.Println("Get New Instance")

		fileInstance = &File{
			Name:   "TIMESTAMP-data.json",
			Folder: "tmp",
			Path:   "./tmp/TIMESTAMP-data.json",
		}

		return fileInstance
	}
	fmt.Println("Instance Already Created")
	return fileInstance
}
