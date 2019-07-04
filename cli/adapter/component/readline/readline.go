package readline

import (
	"github.com/kamiazya/tcho/cli/adapter/component"
	"github.com/chzyer/readline"
	"gopkg.in/urfave/cli.v2"
)

func NewComponent(instance *readline.Instance) component.Component {
	return &readlineComponent{
		instance: instance,
	}
}

var _ component.Component = new(readlineComponent)

type readlineComponent struct {
	instance *readline.Instance
}

func (c *readlineComponent) Commands() []*cli.Command {
	return []*cli.Command{
		&cli.Command{},
	}
}

func (c *readlineComponent) Completer() readline.PrefixCompleterInterface {
	return nil
}
