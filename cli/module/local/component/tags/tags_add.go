package tags

import (
	"io/ioutil"

	"github.com/k0kubun/pp"

	"bitbucket.org/kamiazya/tcho/cli/i18n"
	model "bitbucket.org/kamiazya/tcho/core/domain/model/tag"
	"bitbucket.org/kamiazya/tcho/core/domain/value/colorcode"
	"bitbucket.org/kamiazya/tcho/core/domain/value/memo"
	"gopkg.in/urfave/cli.v2"
)

func (c *tagComponent) addFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "color",
			Aliases:     []string{"c"},
			Usage:       i18n.TagsAddFlagsColorUsage,
			Destination: &i18n.TagsAddFlagsColorDestination,
		},
		&cli.StringFlag{
			Name:    "desc",
			Aliases: []string{"d"},
			Value:   "",
			// Usage:       i18n.TagsAddFlagsColorUsage,
			// Destination: &i18n.TagsAddFlagsColorDestination,
		},
		&cli.BoolFlag{
			Name:    "plain-desc",
			Aliases: []string{"plain"},
			Value:   false,
			// Usage:       i18n.TagsAddFlagsColorUsage,
			// Destination: &i18n.TagsAddFlagsColorDestination,
		},
		// &cli.BoolFlag{
		// 	Name:    "markdown-desc",
		// 	Aliases: []string{"md"},
		// 	Value:   false,
		// 	// Usage:       i18n.TagsAddFlagsColorUsage,
		// 	// Destination: &i18n.TagsAddFlagsColorDestination,
		// },
	}
}

func (c *tagComponent) addCommand() cli.ActionFunc {
	return func(ctx *cli.Context) error {
		if ctx.NArg() == 0 {
			return nil
		}

		// set name
		name := ctx.Args().First()
		t := &model.Tag{
			Name: name,
		}

		// set color
		color := c.ColorCode(ctx.String("color"))
		if color != colorcode.NoColor {
			t.Color = color
		}

		desc := ctx.String("desc")
		if desc != "" {
			t.Description.Type = memo.Plain
			t.Description.Text = desc
		} else {
			if ctx.Bool("plain-desc") {
				t.Description.Type = memo.Plain
				file, err := c.NewMemo(t.Description.Type)
				pp.Println(file)
				defer file.Close()
				if err != nil {
					return err
				}

				bs, err := ioutil.ReadAll(file)
				if err != nil {
					return err
				}
				t.Description.Text = string(bs)
				pp.Println(t.Description.Text)
			}
			// markdown := ctx.Bool("markdown-desc")
			// plain := ctx.Bool("plain-desc")
			// if markdown && plain {
			// 	return
			// }

		}

		_, err := c.uc.AddTag(c.context(), t)
		return err
	}
}
