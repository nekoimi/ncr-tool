package cmd

import (
	"github.com/nekoimi/ncr-tool/logs"
	"github.com/urfave/cli/v2"
	"os"
)

var (
	log  = logs.New()
	root = &cli.App{
		Name:  "ncr",
		Usage: "An ncr-mirror container image proxy command line tool",
	}
)

func Execute(args []string) {
	err := root.Run(args)
	if err != nil {
		log.Errorf("An error occurred：%s\n", err.Error())
		os.Exit(1)
	}
	os.Exit(0)
}
