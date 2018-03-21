package attachments

import (
	"fmt"

	"github.com/bndr/gotabulate"
	"gopkg.in/urfave/cli.v2"
)

func (c *attachments) listCommand() cli.ActionFunc {
	return func(ctx *cli.Context) error {
		mode := ctx.String("mode")
		attachments, err := c.uc.GetAttachmentsAll(c.context())
		if err != nil {
			return err
		}
		switch mode {
		case "simple":
			// simple table
			rows := make([][]string, 0, len(attachments))
			for _, attachment := range attachments {
				rows = append(rows, []string{
					attachment.ID.ID(),
					attachment.Name,
					attachment.Mime,
				})
			}
			table := gotabulate.Create(rows)
			table.SetHeaders([]string{"ID", "Name", "Mime"})
			fmt.Println(table.Render("simple"))

		default:
			for _, attachment := range attachments {
				fmt.Println(attachment.Name)
			}
		}

		return nil
	}
}

func (c *attachments) listFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:  "mode",
			Value: "",
		},
	}
}
