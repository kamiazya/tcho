package component

import (
	"github.com/chzyer/readline"
	"gopkg.in/urfave/cli.v2"
)

type Component interface {
	Commands() []*cli.Command
	Completer() readline.PrefixCompleterInterface
}
