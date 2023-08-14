package main

import (
	"github.com/hebelsan/go-template-cli-light/cmd"
	"github.com/spf13/cobra/doc"
	"log"
)

func main() {
	err := doc.GenMarkdownTree(cmd.GetRootCmd(), "./docs/")
	if err != nil {
		log.Fatal(err)
	}
}
