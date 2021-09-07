package cmd

import (
	"github.com/alob-mtc/todo/todolist"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: `Add task to the list: todo add -s="TASK NAME" -d="TASK DESCRIPTION" `,
	Long:  "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		newApp := todolist.NewApp()
		newApp.Add(subject, des)
	},
}

func init() {
	addCmd.Flags().StringVarP(&subject, "subject", "s", "", "Task Subject")
	addCmd.Flags().StringVarP(&des, "description", "d", "", "Task Description")
	addCmd.MarkFlagRequired("subject")
	addCmd.MarkFlagRequired("description")
	rootCmd.AddCommand(addCmd)

}
