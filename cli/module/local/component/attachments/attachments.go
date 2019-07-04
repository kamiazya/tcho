package attachments

import (
	"context"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/kamiazya/tcho/cli/adapter/component"
	"github.com/kamiazya/tcho/cli/extarnal"
	"github.com/kamiazya/tcho/cli/i18n"
	"github.com/kamiazya/tcho/core/application/usecase"
	r "github.com/chzyer/readline"
	"gopkg.in/urfave/cli.v2"
)

type Option func(*attachments) error

func ContextFunc(contextFunc func() context.Context) Option {
	return func(c *attachments) error {
		c.contextFunc = contextFunc
		return nil
	}
}

func NewComponent(uc usecase.AttachmentUseCase, opts ...Option) (component.Component, error) {
	c := &attachments{
		uc: uc,
	}

	for _, opt := range opts {
		if err := opt(c); err != nil {
			return nil, err
		}
	}
	return c, nil
}

type attachments struct {
	extarnal.TechoCliColor

	uc          usecase.AttachmentUseCase
	contextFunc func() context.Context
}

func (c *attachments) context() context.Context {
	if c.contextFunc != nil {
		return c.contextFunc()
	}
	return context.Background()
}

func (c *attachments) Commands() []*cli.Command {
	return []*cli.Command{
		&cli.Command{
			Name:        "attachments",
			Aliases:     []string{"a"},
			Usage:       i18n.AttachmentsUsage,
			Description: i18n.AttachmentsDescription,
			Flags:       c.listFlags(),
			Action:      c.listCommand(),
			Subcommands: []*cli.Command{
				{
					Name:        "save",
					Usage:       i18n.AttachmentsSaveUsage,
					Description: i18n.AttachmentsSaveDestination,
					Action:      c.saveCommand(),
					Flags:       c.saveFlags(),
				},
				{
					Name:        "list",
					Usage:       i18n.AttachmentsListUsage,
					Description: i18n.AttachmentsListDestination,
					Action:      c.listCommand(),
					Flags:       c.listFlags(),
				},
			},
		},
	}
}

func (c *attachments) paths(dir string) []string {
	if dir == "" {
		dir = "."
	}
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return []string{}
	}

	paths := make([]string, 0, len(files))
	for _, file := range files {
		if file.IsDir() {
			paths = append(paths, c.paths(filepath.Join(dir, file.Name()))...)
			continue
		}
		paths = append(paths, filepath.Join(dir, file.Name()))
	}

	return paths
}

func (c *attachments) Completer() r.PrefixCompleterInterface {
	return r.PcItem(
		"attachments",
		r.PcItem(
			"save",
			r.PcItemDynamic(func(line string) []string {
				words := strings.Split(line, " ")
				return c.paths(words[len(words)-1])
			}),
		),
		r.PcItem(
			"list",
			r.PcItem(
				"--mode=simple",
			),
		),
		r.PcItem("help"),
	)
}
