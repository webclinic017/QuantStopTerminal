package commands

import (
	"context"
	"errors"
	"fmt"
	"github.com/quantstop/quantstopterminal/internal/grpcserver"
)

type SetSubsystemsCommand struct {
	Command
	Client grpcserver.GRPCServerClient
}

func CreateSetSubsystemsCommand(client grpcserver.GRPCServerClient) (*SetSubsystemsCommand, error) {
	if client == nil {
		return nil, errors.New("engine instance is nil")
	}

	base := Command{
		name: "setsubsystem",
		help: "Sets the mode of a subsystem. Example: setsubsystem internet_monitor disable",
	}
	cmd := &SetSubsystemsCommand{
		Command: base,
		Client:  client,
	}
	return cmd, nil
}

func (c *SetSubsystemsCommand) GetName() string {
	return c.name
}

func (c *SetSubsystemsCommand) GetHelpText() string {
	return c.help
}

func (c *SetSubsystemsCommand) Execute(args ...string) error {

	if len(args) != 2 {
		return errors.New("invalid number of arguments supplied")
	}

	if args[0] == "" {
		return errors.New("invalid subsystem supplied")
	}

	switch args[1] {
	case "help":
		fmt.Println()
		fmt.Printf("%20v : %s\n", "setsubsystem", "arg 1 is the subsystem name, arg 2 is the state. Ex: setsubsystem database enable")
		fmt.Println()
		return nil
	case "enable":
		result, err := c.Client.EnableSubsystem(
			context.Background(),
			&grpcserver.GenericSubsystemRequest{
				Subsystem: args[0],
			},
		)
		if err != nil {
			return err
		}
		JsonOutput(result)
		return nil
	case "disable":
		result, err := c.Client.DisableSubsystem(
			context.Background(),
			&grpcserver.GenericSubsystemRequest{
				Subsystem: args[0],
			},
		)
		if err != nil {
			return err
		}
		JsonOutput(result)
		return nil
	default:
		return errors.New("invalid argument supplied. command must be 'enable' or 'disable'")
	}

}
