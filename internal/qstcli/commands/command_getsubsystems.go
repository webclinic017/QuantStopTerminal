package commands

import (
	"context"
	"errors"
	"github.com/quantstop/quantstopterminal/internal/grpcserver"
)

type GetSubsystemsCommand struct {
	Command
	Client grpcserver.GRPCServerClient
}

func CreateGetSubsystemsCommand(client grpcserver.GRPCServerClient) (*GetSubsystemsCommand, error) {
	if client == nil {
		return nil, errors.New("engine instance is nil")
	}

	base := Command{
		name: "getsubsystems",
		help: "Returns the status of all subsystems currently running on QSTrader",
	}
	cmd := &GetSubsystemsCommand{
		Command: base,
		Client:  client,
	}
	return cmd, nil
}

func (c *GetSubsystemsCommand) GetName() string {
	return c.name
}

func (c *GetSubsystemsCommand) GetHelpText() string {
	return c.help
}

func (c *GetSubsystemsCommand) Execute(args ...string) error {
	result, err := c.Client.GetSubsystems(
		context.Background(),
		&grpcserver.GetSubsystemsRequest{},
	)
	if err != nil {
		return err
	}
	JsonOutput(result)
	return nil
}
