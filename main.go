package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	query := "any:"
	if len(os.Args) > 1 {
		query = os.Args[1]
	}

	programFilesX86, ok := os.LookupEnv("ProgramFiles(x86)")
	if !ok {
		fmt.Fprintln(os.Stderr, "ProgramFiles(x86) could not be found.")
		os.Exit(1)
	}
	path := programFilesX86 + "/Evernote/Evernote/ENScript.exe"

	err := exec.Command(path, "showNotes", "/q", query).Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
