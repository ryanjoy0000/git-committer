package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

const (
	FILENAME = "1.txt"
)

func main() {
	// Create a file
	ptrFile, err := os.Create(FILENAME)
	handleErr(err, "error while creating the file")

	// Make changes to the file
	text := "."
	_, err = ptrFile.WriteString(text)
	handleErr(err, "error while writing to file")

	// close the file
	ptrFile.Close()

	// Commit file changes to git repo
	ptrCmd := exec.Command("ls", "-ll")
	bSlice, err := ptrCmd.Output()
	handleErr(err, "error while executing command")
	fmt.Println(string(bSlice))
}

func handleErr(err error, msg string) {
	if err != nil {
		log.Panicln(msg, " : ", err)
	}
}
