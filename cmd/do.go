package cmd

import (
	postgres "cli-task-manager/db"
	"fmt"

	"github.com/spf13/cobra"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "A brief description of your command",
	Long:  `A longer description`,
	Run: func(cmd *cobra.Command, args []string) {
		db, err := postgres.ConnectToDatabase()
		if err != nil {
			panic(err)
		}
		defer db.Close()

		sqlInsertQuery := `
		INSERT INTO completed_tasks(title)
		SELECT title
		FROM tasks
		WHERE id=$1;
		`
		sqlDeleteQuery := `
		DELETE FROM tasks
		WHERE id=$1;
		`

		_, err = db.Exec(sqlInsertQuery, args[0])
		if err != nil {
			panic(err)
		}

		_, err = db.Exec(sqlDeleteQuery, args[0])
		if err != nil {
			panic(err)
		}

		fmt.Println("Task was marked as done!")
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
