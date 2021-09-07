package todolist

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"strconv"
)

var (
	StatusPending   = "Pending"
	StatusCompleted = "Completed"

	pendingFileLocation   = ".pending_todos.csv"
	completedFileLocation = ".completed_todos.csv"
)

// FileStore: create an Abstraction around the data-store
type FileStore struct {
	Data []*Todo
}

func NewFileStore() *FileStore {
	returnValue := &FileStore{}
	returnValue.initialize()
	return returnValue
}

func (f *FileStore) initialize() {

	_, err := ioutil.ReadFile(pendingFileLocation)
	if err != nil {
		if err := ioutil.WriteFile(pendingFileLocation, []byte(""), 0644); err != nil {
			fmt.Println("Error writing csv file", err)
			os.Exit(1)
		}
	}
	_, err = ioutil.ReadFile(completedFileLocation)
	if err != nil {
		if err := ioutil.WriteFile(completedFileLocation, []byte(""), 0644); err != nil {
			fmt.Println("Error writing csv file", err)
			os.Exit(1)
		}
	}
}

func (f *FileStore) LoadTodoList(Status string) ([]*Todo, error) {
	var FileLocation string

	if Status == "Pending" {
		FileLocation = f.getPendingLocation()

	} else if Status == "Completed" {
		FileLocation = f.getCompletedLocation()
	}

	if err := f.load(FileLocation); err != nil {
		return nil, err
	}

	return f.Data, nil
}

func (f *FileStore) load(filepath string) error {
	f.Data = []*Todo{}
	data, err := os.Open(filepath)
	if err != nil {
		fmt.Println("No todo file found!")
		f.initialize()
		fmt.Println("todo file initialize")
		os.Exit(0)
	}
	defer data.Close()

	reader := csv.NewReader(data)
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		Id, err := strconv.Atoi(row[0])
		if err != nil {
			continue
		}
		Completed, err := strconv.ParseBool(row[3])
		if err != nil {
			continue
		}
		f.Data = append(f.Data, &Todo{Id: Id, Subject: row[1], Description: row[2], Completed: Completed, CreatedDate: row[4], CompletedDate: row[5]})
	}
	return nil
}

func (f *FileStore) SaveAll(Status string) {
	//Separate Completed and pending and save separately
	var FileLocation string
	if Status == "Pending" {
		FileLocation = f.getPendingLocation()

	} else if Status == "Completed" {
		FileLocation = f.getCompletedLocation()
	}
	//Save TODO

	data, err := os.Create(FileLocation)
	if err != nil {
		fmt.Println("No todo file found!")
		f.initialize()
		fmt.Println("todo file initialize")
		os.Exit(0)
	}

	defer data.Close()

	writer := csv.NewWriter(data)
	defer writer.Flush()

	for _, row := range f.Data {
		writer.Write(todoToCsv(row))
	}

}

func (f *FileStore) SaveTodo(todo *Todo, Status string) {
	//Separate Completed and pending and save separately
	var FileLocation string
	if Status == "Pending" {
		FileLocation = f.getPendingLocation()

	} else if Status == "Completed" {
		FileLocation = f.getCompletedLocation()
	}
	//Save TODO

	data, err := os.OpenFile(FileLocation, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("No todo file found!")
		fmt.Println("Initialize a new todo repo by running 'todolist init'")
		os.Exit(0)
	}
	defer data.Close()

	writer := csv.NewWriter(data)
	defer writer.Flush()

	if err := writer.Write(todoToCsv(todo)); err != nil {
		fmt.Println(err)
	}
}

func (f *FileStore) DeleteRecord(Status string) {
	var FileLocation string
	if Status == "Pending" {
		FileLocation = f.getPendingLocation()

	} else if Status == "Completed" {
		FileLocation = f.getCompletedLocation()
	}

	if err := ioutil.WriteFile(FileLocation, []byte(""), 0644); err != nil {
		fmt.Println("Error Deleting Record", err)
		os.Exit(1)
	}
}

func todoToCsv(t *Todo) []string {
	return []string{fmt.Sprintf("%d", t.Id), t.Subject, t.Description, strconv.FormatBool(t.Completed), t.CreatedDate, t.CompletedDate}
}

func (f *FileStore) getPendingLocation() string {

	localrepo := pendingFileLocation
	usr, _ := user.Current()
	homerepo := fmt.Sprintf("%s/.pending_todos.csv", usr.HomeDir)
	_, ferr := os.Stat(localrepo)

	if ferr == nil {
		return localrepo
	} else {
		return homerepo
	}
}

func (f *FileStore) getCompletedLocation() string {

	localrepo := completedFileLocation
	usr, _ := user.Current()
	homerepo := fmt.Sprintf("%s/.completed_todos.csv", usr.HomeDir)
	_, ferr := os.Stat(localrepo)

	if ferr == nil {
		return localrepo
	} else {
		return homerepo
	}
}

func (f *FileStore) NextId() int {
	maxId := 0
	for _, todo := range f.Data {
		if todo.Id > maxId {
			maxId = todo.Id
		}
	}
	return maxId + 1
}

func (f *FileStore) RemoveById(id int) *Todo {
	for i, todo := range f.Data {
		if todo.Id == id {
			f.Data = append(f.Data[0:i], f.Data[i+1:len(f.Data)]...)
			return todo
		}
	}
	return nil
}
