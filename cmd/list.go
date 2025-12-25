package cmd

import (
	"fmt"
	"os"
	"stask/db"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all of your incomplete tasks",
	Long:  `List all of your incomplete tasks currently stored in the database.`,
	Run: func(cmd *cobra.Command, args []string) {
		color.Magenta("This are your tasks: \n")

		tasks, err := db.ListToDoTasks()
		if err != nil {
			fmt.Printf("Error listing TODO tasks. Err: %v", err)
			os.Exit(1)
		}

		for _, task := range tasks {
			fmt.Printf("%d. %s\n", task.ID, task.Task)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
