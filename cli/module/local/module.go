package local

import (
	"database/sql"
	"io"
	"log"
	"strings"

	"github.com/kamiazya/tcho/cli/adapter/component"
	"github.com/kamiazya/tcho/cli/adapter/component/readline"
	"github.com/kamiazya/tcho/cli/module"
	attachmentscmp "github.com/kamiazya/tcho/cli/module/local/component/attachments"
	tagscmp "github.com/kamiazya/tcho/cli/module/local/component/tags"
	"github.com/kamiazya/tcho/core/application/usecase"
	"github.com/kamiazya/tcho/core/domain/model"
	"github.com/kamiazya/tcho/core/infrastructure/database/sqlite"
	r "github.com/chzyer/readline"
	"gopkg.in/urfave/cli.v2"
)

var _ module.Module = new(localModule)

func New(app *cli.App) module.Module {
	return &localModule{
		app: app,
	}
}

type localModule struct {
	db       *sql.DB
	app      *cli.App
	instance *r.Instance

	tagsComponent        component.Component
	attachmentsComponent component.Component
	readlineCmp          component.Component
}

func (p *localModule) Build() error {
	var err error
	// open db
	if p.db, err = sql.Open("sqlite3", "./data/tcho.sqlite"); err != nil {
		return err
	}

	// repository
	repos := sqlite.RepositoryContainer(p.db, model.Tag, model.Attachment)

	// usecase
	tagUsecase, err := usecase.NewTagUseCase(repos)
	if err != nil {
		return err
	}

	attachmentUsecase, err := usecase.NewAttachmentUseCase(repos)
	if err != nil {
		return err
	}

	// tags component
	if p.tagsComponent, err = tagscmp.NewComponent(tagUsecase); err != nil {
		return err
	}

	// attachments component
	if p.attachmentsComponent, err = attachmentscmp.NewComponent(attachmentUsecase); err != nil {
		return err
	}

	cmds := make([]*cli.Command, 0, 10)
	cmds = append(cmds, p.tagsComponent.Commands()...)
	cmds = append(cmds, p.attachmentsComponent.Commands()...)

	p.app.Commands = cmds

	p.app.Action = p.readline()
	return nil
}

func (p *localModule) Close() {
	if p.instance != nil {
		p.instance.Close()
	}
	if p.db != nil {
		p.db.Close()
	}
}

func (p *localModule) Run(arguments []string) (err error) {
	return p.app.Run(arguments)
}

func (p *localModule) readline() func(c *cli.Context) error {
	return func(c *cli.Context) (err error) {

		p.app.Action = func(*cli.Context) error {
			return nil
		}

		// create readline instance
		if p.instance, err = r.NewEx(&r.Config{
			Prompt:            "\033[31mtch »\033[0m ",
			VimMode:           false,
			InterruptPrompt:   "^C",
			EOFPrompt:         "exit",
			HistorySearchFold: true,
			HistoryFile:       "data/history",
			AutoComplete: r.NewPrefixCompleter(
				p.tagsComponent.Completer(),
				p.attachmentsComponent.Completer(),
			),
		}); err != nil {
			return err
		}

		// readline component
		p.readlineCmp = readline.NewComponent(p.instance)

		p.app.Commands = append(p.app.Commands, p.readlineCmp.Commands()...)

		log.SetOutput(p.instance.Stderr())
		// initial
		var (
			line        string
			readlineErr error
			args        = make([]string, 0, 10)
		)
		for {
			line, readlineErr = p.instance.Readline()
			if readlineErr == r.ErrInterrupt {
				if len(line) == 0 {
					break
				} else {
					continue
				}
			} else if err == io.EOF {
				break
			}

			line = strings.TrimSpace(line)
			args = append(args, p.app.Name)
			args = append(args, strings.Split(line, " ")...)
			p.app.Run(args)

			// reset
			args = args[:0]
		}
		return
	}
}
