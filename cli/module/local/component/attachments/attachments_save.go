package attachments

import (
	"github.com/k0kubun/pp"
	"gopkg.in/urfave/cli.v2"
)

func (c *attachments) saveFlags() []cli.Flag {
	return []cli.Flag{
	// &cli.StringFlag{
	// 	Name:        "color",
	// 	Aliases:     []string{"c"},
	// 	Usage:       i18n.TagsAddFlagsColorUsage,
	// 	Destination: &i18n.TagsAddFlagsColorDestination,
	// },
	}
}

func (c *attachments) saveCommand() cli.ActionFunc {
	return func(ctx *cli.Context) error {
		files := ctx.Args().Slice()
		for _, file := range files {
			pp.Println(file)
		}

		// filePath := ctx.Args().First()
		// a := &attachment.Attachment{
		// 	Name: filePath,
		// }

		// // set color
		// color := c.ColorCode(ctx.String("color"))
		// if color != colorcode.NoColor {
		// 	t.Color = color
		// }

		// _, err := c.uc.SaveAttachment(c.context(), t, nil)
		// return err
		return nil
	}
}
