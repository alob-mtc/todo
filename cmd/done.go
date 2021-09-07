package cmd

import (
	"strconv"

	"github.com/alob-mtc/todo/todolist"
	"github.com/spf13/cobra"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Mark a task as done: todo done -a=task_id",
	Long:  "would mark a task as done",
	Run: func(cmd *cobra.Command, args []string) {

		Id, err := strconv.Atoi(id)
		if err != nil {
			return
		}
		newApp := todolist.NewApp()
		newApp.Done(Id)
	},
}

func init() {
	doneCmd.Flags().StringVarP(&id, "task_id", "a", "", "Task ID")
	doneCmd.MarkFlagRequired("task_id")
	rootCmd.AddCommand(doneCmd)

}
