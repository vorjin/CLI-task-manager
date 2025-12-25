package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"stask/db"
)

// completedCmd represents the completed command
var completedCmd = &cobra.Command{
	Use:   "completed",
	Short: "show completed tasks",
	Long:  "show completed tasks",
	Run: func(cmd *cobra.Command, args []string) {
		hours, err := cmd.Flags().GetInt("time")
		if err != nil {
			panic(err)
		}

		color.Cyan("This are your completed tasks for the last %d hours: \n", hours)

		tasks, err := db.ListCompletedTasks(hours)
		if err != nil {
			panic(err)
		}

		for i, task := range tasks {
			fmt.Printf("%d. %s\n", i+1, task.Task)
			i++
		}
	},
}

func init() {
	completedCmd.Flags().IntP("time", "t", 24, "how many hours ago tasks were completed")
	rootCmd.AddCommand(completedCmd)
}
