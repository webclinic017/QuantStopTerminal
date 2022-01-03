package commands

import "fmt"

type HelpCommand struct {
	Command
}

func CreateHelpCommand() *HelpCommand {
	base := Command{
		name: "help",
		help: "Displays the help text for each command.",
	}
	cmd := &HelpCommand{Command: base}
	return cmd
}

func (c *HelpCommand) GetName() string {
	return c.name
}

func (c *HelpCommand) GetHelpText() string {
	return c.help
}

func (c *HelpCommand) Execute(args ...string) error {

	fmt.Println("Executing help command")
	return nil
}
