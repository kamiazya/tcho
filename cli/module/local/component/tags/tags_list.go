package tags

import (
	"fmt"

	"github.com/bndr/gotabulate"
	"gopkg.in/urfave/cli.v2"
)

func (c *tagComponent) listCommand() cli.ActionFunc {
	return func(ctx *cli.Context) error {
		mode := ctx.String("mode")
		tags, err := c.uc.GetTagsAll(c.context())
		if err != nil {
			return err
		}
		switch mode {
		case "simple":
			// simple table
			rows := make([][]string, 0, len(tags))
			for _, tag := range tags {
				rows = append(rows, []string{
					tag.ID.ID(),
					tag.Name,
					tag.Color.String(),
					tag.Description.Text,
				})
			}
			table := gotabulate.Create(rows)
			table.SetHeaders([]string{"ID", "Name", "Color", "Description"})
			fmt.Println(table.Render("simple"))

		default:
			for _, tag := range tags {
				c.Color(tag.Color).Println(tag.Name)
			}
		}

		return nil
	}
}

func (c *tagComponent) listFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:  "mode",
			Value: "",
		},
	}
}
