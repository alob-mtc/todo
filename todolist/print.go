package todolist

import (
	"fmt"
	"text/tabwriter"

	"github.com/alob-mtc/todo/util"
	"github.com/fatih/color"
)

type ScreenPrinter struct {
	Writer   *tabwriter.Writer
	fgGreen  func(a ...interface{}) string
	fgYellow func(a ...interface{}) string
	fgRed    func(a ...interface{}) string
}

func NewScreenPrinter() *ScreenPrinter {
	w := new(tabwriter.Writer)
	w.Init(color.Output, 0, 8, 1, ' ', 0) //Changed from os.Stdout to color.Output when compiled on Windows.
	//w.Init(color.Output, 5, 0, 1, ' ', tabwriter.StripEscape)
	green := color.New(color.FgGreen).Add(color.Bold).SprintFunc()
	yellow := color.New(color.FgYellow).Add(color.Bold).SprintFunc()
	red := color.New(color.FgRed).Add(color.Bold).SprintFunc()

	formatter := &ScreenPrinter{Writer: w, fgGreen: green, fgYellow: yellow, fgRed: red}
	return formatter
}

func (f *ScreenPrinter) PrintTodoDetail(todos []*Todo) {
	key := f.fgGreen
	val := f.fgYellow

	f.Println("==== TODO List ==== \n", "Y")
	for i, todo := range todos {
		if i > 0 {
			fmt.Println("")
		}
		fmt.Fprintf(f.Writer, " %s\t%s\n", key("ID:"), val(todo.Id))
		fmt.Fprintf(f.Writer, " %s\t%s\n", key("Subject:"), val(todo.Subject))
		fmt.Fprintf(f.Writer, " %s\t%s\n", key("Description:"), val(todo.Description))
		fmt.Fprintf(f.Writer, " %s\t%s\n", key("Status:"), val(util.CompletedToWords(todo.Completed)))
		fmt.Fprintf(f.Writer, " %s\t%s\n", key("CreatedDate:"), val(todo.CreatedDate))
		fmt.Fprintf(f.Writer, " %s\t%s\n", key("CompletedDate:"), val(todo.CompletedDate))
		f.Writer.Flush()
	}
	f.Println("\n============== \n", "Y")
}

func (f *ScreenPrinter) Println(val string, colour string) {

	switch colour {
	case "G":
		fmt.Fprintf(f.Writer, " %s\n", f.fgGreen(val))
	case "Y":
		f.fgYellow(val)
		fmt.Fprintf(f.Writer, " %s\n", f.fgYellow(val))
	case "R":
		fmt.Fprintf(f.Writer, " %s\n", f.fgRed(val))
	default:
		fmt.Printf("Color not Supported")
	}

}
