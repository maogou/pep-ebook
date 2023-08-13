package main

import (
	"log"
	"os"

	"github.com/maogou/pep-ebook/internal"
)

func main() {
	pepeb := internal.NewPepEbAppCommand()

	if err := pepeb.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
