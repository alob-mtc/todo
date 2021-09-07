package cmd

import (
	"strconv"

	"github.com/alob-mtc/todo/todolist"
	"github.com/spf13/cobra"
)

// undoneCmd represents the undone command
var undoneCmd = &cobra.Command{
	Use:   "undone",
	Short: "Mark a task as undone: todo undone -a=task_id",
	Long:  "would mark a task as undone",
	Run: func(cmd *cobra.Command, args []string) {

		Id, err := strconv.Atoi(id)
		if err != nil {
			return
		}
		newApp := todolist.NewApp()
		newApp.Undone(Id)
	},
}

func init() {
	undoneCmd.Flags().StringVarP(&id, "task_id", "a", "", "Task ID")
	undoneCmd.MarkFlagRequired("task_id")
	rootCmd.AddCommand(undoneCmd)
}
