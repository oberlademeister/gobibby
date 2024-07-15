package main

import (
	_ "embed"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:   "gobibby",
		Usage:  "Compile .jsons to .bib",
		Action: run,
		Flags: []cli.Flag{
			&cli.PathFlag{
				Name:     "dbpath",
				Required: true,
			},
			&cli.PathFlag{
				Name:     "outfile",
				Value:    "out.bib",
				Required: true,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}

//go:embed structure.cue
var cuefile []byte
