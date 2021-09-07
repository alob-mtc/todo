package todolist

import (
	"fmt"
)

type App struct {
	TodoStore *FileStore
	Printer   *ScreenPrinter
}

func NewApp() *App {
	app := &App{
		Printer:   NewScreenPrinter(),
		TodoStore: NewFileStore(),
	}
	return app
}

// List: returns the pending Task
func (a *App) List() {
	todolist, err := a.TodoStore.LoadTodoList(StatusPending)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	if len(todolist) == 0 {
		a.Printer.Println("No Task", "R")
		return
	}
	a.Printer.PrintTodoDetail(todolist)
}

//Cleanup: clears all completed Task
func (a *App) Cleanup() {
	a.TodoStore.DeleteRecord(StatusCompleted)
	a.Printer.Println("Cleaned", "G")
}

// Add: Creates new task
func (a *App) Add(subject, description string) {
	a.TodoStore.LoadTodoList(StatusPending)
	todo := NewTodo(a.TodoStore.NextId(), subject, description)
	a.TodoStore.SaveTodo(todo, StatusPending)
	a.Printer.Println("Task added to TODO", "Y")
}

// Done: merk Task as completed
func (a *App) Done(id int) {
	if _, err := a.TodoStore.LoadTodoList(StatusPending); err != nil {
		fmt.Println(err)
		return
	}

	todo := a.TodoStore.RemoveById(id)
	if todo == nil {
		fmt.Println("Task does not exist")
		return
	}

	a.TodoStore.SaveAll(StatusPending)
	if _, err := a.TodoStore.LoadTodoList(StatusCompleted); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(a.TodoStore.NextId())
	todo.Complete(a.TodoStore.NextId())
	a.TodoStore.SaveTodo(todo, StatusCompleted)

	a.Printer.Println("Task marked as Done", "Y")
}

// Undone: mark Task has uncompleted
func (a *App) Undone(id int) {
	a.TodoStore.LoadTodoList(StatusCompleted)

	todo := a.TodoStore.RemoveById(id)
	if todo == nil {
		fmt.Println("Task does not exist")
		return
	}
	a.TodoStore.SaveAll(StatusCompleted)

	if _, err := a.TodoStore.LoadTodoList(StatusPending); err != nil {
		fmt.Println(err)
		return
	}
	todo.Uncomplete(a.TodoStore.NextId())
	a.TodoStore.SaveTodo(todo, StatusPending)
	a.Printer.Println("Task marked as Undone", "Y")

}
