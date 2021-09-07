package cmd

import (
	"github.com/alob-mtc/todo/todolist"
	"github.com/spf13/cobra"
)

// listCmd represents the serve command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks still to do: todo list",
	Long:  "list out all tasks that are not done",
	Run: func(cmd *cobra.Command, args []string) {
		newApp := todolist.NewApp()
		newApp.List()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
