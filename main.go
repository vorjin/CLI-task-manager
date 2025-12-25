/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/mitchellh/go-homedir"
	"stask/cmd"
	"stask/db"
)

func main() {
	var path = `/task-manager.db`
	osDir, err := homedir.Dir()
	if err != nil {
		panic(err)
	}
	dbPath := osDir + path

	err = db.BoltDBInit(dbPath)
	if err != nil {
		panic(err)
	}
	cmd.Execute()

	err = db.CloseBoltDB()
	if err != nil {
		panic(err)
	}
}
