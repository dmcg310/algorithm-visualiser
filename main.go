package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:      "Algorithm Visualiser",
		Usage:     "Visualise algorithms in the terminal!",
		UsageText: "algorithm-visualiser",
		Description: `This program allows for a number of algorithms 
		to be visualised in the terminal.`,
		Action: func(ctx *cli.Context) error {
			if s, err := InitScreen(); err != nil {
				log.Fatal(err)
			} else {
				NewApp(s, *NewGrid(s.Size())).Run()
			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
