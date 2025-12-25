/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"os"
	"stask/cmd"
	"stask/db"
)

func main() {
	var path = `/task-manager.db`
	osDir, err := homedir.Dir()
	if err != nil {
		fmt.Printf("Error getting user home directory. Err: %v", err)
		os.Exit(1)
	}
	dbPath := osDir + path

	err = db.BoltDBInit(dbPath)
	if err != nil {
		fmt.Printf("Error initialising Bolt database. Err: %v", err)
		os.Exit(1)
	}
	cmd.Execute()

	err = db.CloseBoltDB()
	if err != nil {
		fmt.Printf("Error closing Bolt database. Err: %v", err)
		os.Exit(1)
	}
}
