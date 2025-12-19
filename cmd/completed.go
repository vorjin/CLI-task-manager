package cmd

import (
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
		color.Cyan("This are your completed tasks: \n")

		err := db.ListTasks("completed")
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(completedCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// completedCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// completedCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
