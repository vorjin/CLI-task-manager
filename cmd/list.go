package cmd

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"stask/db"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all of your incomplete tasks",
	Long:  `List all of your incomplete tasks currently stored in the database.`,
	Run: func(cmd *cobra.Command, args []string) {
		color.Magenta("This are your tasks: \n")

		err := db.ListToDoTasks()
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
