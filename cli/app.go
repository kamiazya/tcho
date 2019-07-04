package main

import (
	"fmt"

	"github.com/kamiazya/tcho/cli/i18n"
	"gopkg.in/urfave/cli.v2"
)

var app = &cli.App{
	Name:  "tcho-cli",
	Usage: i18n.AppUsage,
	Authors: []*cli.Author{
		{
			Name:  "kamiazya",
			Email: "yuki@kamaizya.tech",
		},
	},
	Copyright: "(c) 2018 kamiazya",

	CommandNotFound: func(ctx *cli.Context, cmd string) {
		fmt.Printf("command not found: %s\n", cmd)
	},

	Version: version(),
}
