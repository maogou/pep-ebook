package main

import (
	"log"
	"os"

	"github.com/maogou/pep-ebook/cmd"
)

func main() {
	pepeb := cmd.CreatePepEb()

	if err := pepeb.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
