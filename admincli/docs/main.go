package main

import (
	"github.com/bixlabs/authentication/admincli/cmd"
	"github.com/spf13/cobra/doc"
	"os"
	"path"
)

func main() {
	header := &doc.GenManHeader{
		Title:   "Admin Client",
		Section: "1",
	}

	workingDir, err := os.Getwd()
	if err != nil {
		workingDir = "/tmp"
	} else {
		workingDir = path.Join(workingDir, "admincli/docs")
	}

	err = doc.GenManTree(cmd.GetRootCommand().Command, header, workingDir)

	if err != nil {
		panic(err)
	}
}
