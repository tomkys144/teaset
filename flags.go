package main

import "github.com/urfave/cli/v2"

var setup_flags = []cli.Flag{
	&cli.BoolFlag{
		Name:    "verbose",
		Aliases: []string{"v"},
		Usage:   "If set log level will be set to INFO.",
	},
	&cli.PathFlag{
		Name:    "path",
		Aliases: []string{"p"},
		Usage:   "Path to env file from which settings will be generated.",
	},
	&cli.PathFlag{
		Name:    "destination",
		Aliases: []string{"d"},
		Usage:   "Path to directory, where Teaset will be set up.",
	},
	&cli.BoolFlag{
		Name:    "interactive",
		Aliases: []string{"i"},
		Usage:   "Will use interactive mode.",
	},
}

var env_gen_flags = []cli.Flag{
	&cli.PathFlag{
		Name:    "path",
		Aliases: []string{"p"},
		Usage:   "Path to env file from which settings will be generated (don't need to be complete).",
	},
	&cli.PathFlag{
		Name:    "destination",
		Aliases: []string{"d"},
		Usage:   "Path to directory, where Teaset .env file will be stored. Usuall this is Teaset home directory.",
	},
}
