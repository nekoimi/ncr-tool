package cmd

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

var root = &cli.App{
	Name:  "ncr",
	Usage: "An ncr-mirror container image proxy command line tool",
}

func Execute(args []string) {
	err := root.Run(args)
	if err != nil {
		log.Fatalln(err)
	}
	os.Exit(0)
}
