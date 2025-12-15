package cmd

import (
	"encoding/binary"
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"
	"strings"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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

			id, err := bucket.NextSequence()
			if err != nil {
				panic(err)
			}

			idBytes := make([]byte, 8)
			binary.BigEndian.PutUint64(idBytes, id)

			todoTask := []byte(strings.Join(args, " "))

			err = bucket.Put(idBytes, todoTask)
			if err != nil {
				panic(err)
			}

			return nil
		})

		if err != nil {
			panic(err)
		}

		fmt.Println("Task was added succesfully.")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
