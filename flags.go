package main

import "github.com/urfave/cli/v2"

var setup_flags = []cli.Flag{
	&cli.BoolFlag{
		Name:    "verbose",
		Aliases: []string{"v"},
		Usage:   "If set log level will be set to INFO",
	},
}
