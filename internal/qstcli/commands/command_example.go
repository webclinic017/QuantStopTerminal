package commands

import "fmt"

type ExampleCommand struct {
	Command
}

func CreateExampleCommand() *ExampleCommand {
	base := Command{
		name: "example",
		help: "An example command to use as a starting point.",
	}
	cmd := &ExampleCommand{Command: base}
	return cmd
}

func (c *ExampleCommand) GetName() string {
	return c.name
}

func (c *ExampleCommand) GetHelpText() string {
	return c.help
}

func (c *ExampleCommand) Execute(args ...string) error {
	fmt.Println("Executing example command")
	return nil
}
