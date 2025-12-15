package cmd

import (
	"encoding/binary"
	"fmt"
	"strconv"

	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do [task number]",
	Short: "Mark a task as complete",
	Long:  `Mark a task as complete by providing its number from the list.`,
	Run: func(cmd *cobra.Command, args []string) {
		db, err := bolt.Open("task-manager.db", 0644, nil)
		if err != nil {
			panic(err)
		}
		defer db.Close()

		err = db.Update(func(tx *bolt.Tx) error {
			bucket := tx.Bucket([]byte("tasks"))
			if bucket == nil {
				fmt.Println("bucket not found")
				return nil
			}

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				panic(err)
			}

			idBytes := make([]byte, 8)
			binary.BigEndian.PutUint64(idBytes, id)

			err = bucket.Delete(idBytes)
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
