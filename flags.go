package main

import "github.com/urfave/cli/v2"

var SetupFlags = []cli.Flag{
	&cli.BoolFlag{
		Name:    "verbose",
		Aliases: []string{"v"},
		Usage:   "If set log level will be set to INFO.",
	},
	&cli.PathFlag{
		Name:        "source",
		Aliases:     []string{"s"},
		DefaultText: ".env.example",
		Usage:       "Path to env file from which settings will be generated.",
	},
	&cli.PathFlag{
		Name:        "destination",
		Aliases:     []string{"d"},
		DefaultText: "teaset/",
		Usage:       "Path to directory, where Teaset will be set up.",
	},
	&cli.BoolFlag{
		Name:    "interactive",
		Aliases: []string{"i"},
		Usage:   "Will use interactive mode.",
	},
}

var Env_genFlags = []cli.Flag{
	&cli.PathFlag{
		Name:        "source",
		Aliases:     []string{"s"},
		DefaultText: ".env.example",
		Usage:       "Path to env file from which settings will be generated (don't need to be complete).",
	},
	&cli.PathFlag{
		Name:        "destination",
		Aliases:     []string{"d"},
		DefaultText: "teaset/",
		Usage:       "Path to directory, where Teaset .env file will be stored. Usually this is Teaset home directory.",
	},
}

var StartFlags = []cli.Flag{
	// TODO: Add flags for 'start' command
}

var StopFlags = []cli.Flag{
	// TODO: Add flags for 'stop' command
}

var UpdateFlags = []cli.Flag{
	// TODO: Add flags for 'update' command
}
var VersionFlags = []cli.Flag{
	// TODO: Add flags for 'version' command
}
