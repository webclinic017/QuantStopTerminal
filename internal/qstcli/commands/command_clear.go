package commands

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

type ClearCommand struct {
	Command
}

func CreateClearCommand() *ClearCommand {
	base := Command{
		name: "clear",
		help: "Clear's the terminal window output.",
	}
	cmd := &ClearCommand{Command: base}
	return cmd
}

func (c *ClearCommand) GetName() string {
	return c.name
}

func (c *ClearCommand) GetHelpText() string {
	return c.help
}

func (c *ClearCommand) Execute(args ...string) error {

	switch runtime.GOOS {
	case "linux":
		return clearLinux()
	case "windows":
		return clearWindows()
	default:
		return fmt.Errorf("system not supported: %s", runtime.GOOS)
	}

}

func clearWindows() error {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func clearLinux() error {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
