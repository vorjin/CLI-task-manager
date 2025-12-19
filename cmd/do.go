package cmd

import (
	"fmt"
	"stask/db"
	"strconv"
	"time"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do [task number]",
	Short: "Mark a task as complete",
	Long:  `Mark a task as complete by providing its number from the list.`,
	Run: func(cmd *cobra.Command, args []string) {

		for _, taskID := range args {
			id, err := strconv.ParseUint(taskID, 10, 64)
			if err != nil {
				panic(err)
			}

			taskDesc, err := db.TaskByID(id)
			if err != nil {
				panic(err)
			}

			err = db.AddTask(taskDesc, "completed")
			if err != nil {
				panic(err)
			}

			timeNow := []byte(time.Now().Format(time.RFC3339))

			err = db.AddTask(timeNow, "completed_time")
			if err != nil {
				panic(err)
			}

			err = db.DeleteTask(id)

			if err != nil {
				panic(err)
			}

			color.Set(color.FgGreen)
			fmt.Printf("Task #%s was marked as done!\n", taskID)
		}
	},
}

func init() {
	rootCmd.AddCommand(doCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
