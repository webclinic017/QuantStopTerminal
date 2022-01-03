package commands

import (
	"os"
)

type QuitCommand struct {
	Command
}

func CreateQuitCommand() *QuitCommand {
	base := Command{
		name: "exit",
		help: "Exit QST-Console",
	}
	cmd := &QuitCommand{Command: base}
	return cmd
}

func (c *QuitCommand) GetName() string {
	return c.name
}

func (c *QuitCommand) GetHelpText() string {
	return c.help
}

func (c *QuitCommand) Execute(args ...string) error {
	os.Exit(1)
	return nil
}
