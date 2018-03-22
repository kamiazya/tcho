package tags

import (
	"context"

	"bitbucket.org/kamiazya/tcho/cli/adapter/component"
	"bitbucket.org/kamiazya/tcho/cli/extarnal"
	"bitbucket.org/kamiazya/tcho/cli/extarnal/editor"
	"bitbucket.org/kamiazya/tcho/cli/i18n"
	"bitbucket.org/kamiazya/tcho/core/application/usecase"
	r "github.com/chzyer/readline"
	"gopkg.in/urfave/cli.v2"
)

type Option func(*tagComponent) error

func Editor(editor editor.TechoCliMemoEditor) Option {
	return func(c *tagComponent) error {
		c.TechoCliMemoEditor = editor
		return nil
	}
}

func ContextFunc(contextFunc func() context.Context) Option {
	return func(c *tagComponent) error {
		c.contextFunc = contextFunc
		return nil
	}
}

func NewComponent(uc usecase.TagUseCase, opts ...Option) (component.Component, error) {
	c := &tagComponent{
		uc:                 uc,
		TechoCliMemoEditor: editor.Default,
	}

	for _, opt := range opts {
		if err := opt(c); err != nil {
			return nil, err
		}
	}
	return c, nil
}

type tagComponent struct {
	extarnal.TechoCliColor
	editor.TechoCliMemoEditor

	uc          usecase.TagUseCase
	contextFunc func() context.Context
}

func (c *tagComponent) context() context.Context {
	if c.contextFunc != nil {
		return c.contextFunc()
	}
	return context.Background()
}

func (c *tagComponent) Commands() []*cli.Command {
	return []*cli.Command{
		&cli.Command{
			Name:        "tags",
			Aliases:     []string{"t"},
			Usage:       i18n.TagsUsage,
			Description: i18n.TagsDescription,
			Flags:       c.listFlags(),
			Action:      c.listCommand(),
			Subcommands: []*cli.Command{
				{
					Name:        "add",
					Usage:       i18n.TagsAddUsage,
					Description: i18n.TagsAddDestination,
					Action:      c.addCommand(),
					Flags:       c.addFlags(),
				},
				{
					Name:        "list",
					Usage:       i18n.TagsListUsage,
					Description: i18n.TagsListDestination,
					Action:      c.listCommand(),
					Flags:       c.listFlags(),
				},
				{
					Name: "edit",
					// Usage:       i18n.TagsListUsage,
					// Description: i18n.TagsListDestination,
					Action: c.editCommand(),
					Flags:  c.editFlags(),
				},
			},
		},
	}
}

func (c *tagComponent) Completer() r.PrefixCompleterInterface {
	return r.PcItem(
		"tags",
		r.PcItem(
			"add",
			r.PcItemDynamic(func(val string) []string {
				return []string{
					"--color=red",
					"--color=green",
					"--color=yellow",
					"--color=blue",
					"--color=magenta",
					"--color=cyan",
				}
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
