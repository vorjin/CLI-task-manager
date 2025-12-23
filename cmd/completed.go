package cmd

import (
	"fmt"
	"stask/db"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
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

		color.Set(color.FgCyan)
		fmt.Printf("This are your completed tasks for the last %d hours: \n", hours)
		color.Unset()

		err = db.ListCompletedTasks(hours)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	completedCmd.Flags().IntP("time", "t", 24, "how many hours ago tasks were completed")
	rootCmd.AddCommand(completedCmd)
}
