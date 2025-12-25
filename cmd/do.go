package cmd

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"stask/db"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do [task number]",
	Short: "Mark a task as complete",
	Long:  `Mark a task as complete by providing its number from the list.`,
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.DoTask(args)

		if err != nil {
			panic(err)
		}

		for _, taskID := range tasks {
			color.Green("Task #%s was marked as done!\n", taskID)
		}
	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}
