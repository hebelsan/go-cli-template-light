package main

import (
	"github.com/spf13/cobra/doc"
	"go-cli-template/cmd"
	"log"
)

func main() {
	err := doc.GenMarkdownTree(cmd.GetRootCmd(), "./docs/")
	if err != nil {
		log.Fatal(err)
	}
}
