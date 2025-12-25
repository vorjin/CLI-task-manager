package cmd

import (
	"fmt"
	"os"
	"stask/db"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// clearCmd represents the clear command
var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "clears TODO tasks list",
	Run: func(cmd *cobra.Command, args []string) {
		err := db.DeleteTasksBucket()
		if err != nil {
			fmt.Printf("Error clearing TODO list. Err: %v", err)
			os.Exit(1)
		}

		color.Green("TODO list was succesfully cleared!")
	},
}

func init() {
	rootCmd.AddCommand(clearCmd)
}
