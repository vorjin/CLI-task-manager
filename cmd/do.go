package cmd

import (
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"
	"strings"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "A brief description of your command",
	Long:  `A longer description`,
	Run: func(cmd *cobra.Command, args []string) {
		db, err := bolt.Open("task-manager.db", 0644, nil)
		if err != nil {
			panic(err)
		}
		defer db.Close()

		err = db.Update(func(tx *bolt.Tx) error {
			bucket, err := tx.CreateBucketIfNotExists([]byte("tasks"))
			if err != nil {
				panic(err)
			}

			doTask := []byte(strings.Join(args, " "))

			err = bucket.Delete(doTask)
			if err != nil {
				panic(err)
			}

			return nil
		})

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
