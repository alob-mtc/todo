package cmd

import (
	"github.com/alob-mtc/todo/todolist"
	"github.com/spf13/cobra"
)

// cleanupCmd represents the cleanup command
var cleanupCmd = &cobra.Command{
	Use:   "cleanup",
	Short: "Clean up done task: todo cleanup",
	Long:  "delete tasks that have been done",
	Run: func(cmd *cobra.Command, args []string) {
		newApp := todolist.NewApp()
		newApp.Cleanup()
	},
}

func init() {
	rootCmd.AddCommand(cleanupCmd)

}
