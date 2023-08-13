package cmd

import (
	"github.com/maogou/pep-ebook/internal"
	"github.com/urfave/cli/v2"
)

func CreatePepEb() *cli.App {
	return internal.NewPepEbAppCommand()
}
