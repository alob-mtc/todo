package todolist

import (
	"time"

	"github.com/alob-mtc/todo/util"
)

type Todo struct {
	Id            int
	Subject       string
	Description   string
	CreatedDate   string
	CompletedDate string
	Completed     bool
}

// NewTodo: create an new task item
func NewTodo(id int, subject, description string) *Todo {
	return &Todo{Id: id, Subject: subject, Description: description, CreatedDate: util.TimeToString(time.Now()), Completed: false}
}

func (t *Todo) Complete(newId int) {
	t.Id = newId
	t.Completed = true
	t.CompletedDate = util.TimeToString(time.Now())
}

func (t *Todo) Uncomplete(newId int) {
	t.Id = newId
	t.Completed = false
	t.CompletedDate = ""
}
