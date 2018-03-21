package tags

import (
	"gopkg.in/urfave/cli.v2"
)

func (c *tagComponent) editFlags() []cli.Flag {
	return []cli.Flag{}
}

func (c *tagComponent) editCommand() cli.ActionFunc {
	return func(ctx *cli.Context) error {
		return nil
	}
}
