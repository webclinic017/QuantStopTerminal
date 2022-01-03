package commands

import (
	"fmt"
	"github.com/quantstop/quantstopterminal/internal/grpcserver"
	"reflect"
)

// CommandRegistry provides a useful pattern for managing commands.
// It allows for ease of dependency management and ensures commands
// dependent on others use the same references in memory.
type CommandRegistry struct {
	commands     map[reflect.Type]iCommand // map of types to iCommand.
	commandTypes []reflect.Type            // keep an ordered slice of registered service types.
}

// NewCommandRegistry starts a registry instance for convenience
func NewCommandRegistry() *CommandRegistry {
	return &CommandRegistry{
		commands: make(map[reflect.Type]iCommand),
	}
}

// InitCommands provides a function call to initialize all commands in the registry
func (s *CommandRegistry) InitCommands(client grpcserver.GRPCServerClient) error {

	var err error

	clear := CreateClearCommand()
	err = s.RegisterCommand(clear)
	if err != nil {
		return err
	}

	getsubsys, err := CreateGetSubsystemsCommand(client)
	if err != nil {
		return err
	}
	err = s.RegisterCommand(getsubsys)
	if err != nil {
		return err
	}

	exit := CreateQuitCommand()
	err = s.RegisterCommand(exit)
	if err != nil {
		return err
	}

	setsubsys, err := CreateSetSubsystemsCommand(client)
	if err != nil {
		return err
	}
	err = s.RegisterCommand(setsubsys)
	if err != nil {
		return err
	}

	logmon := CreateLogmonCommand()
	err = s.RegisterCommand(logmon)
	if err != nil {
		return err
	}

	return nil
}

// RegisterCommand appends a command constructor function to the command registry.
func (s *CommandRegistry) RegisterCommand(command iCommand) error {
	kind := reflect.TypeOf(command)
	if _, exists := s.commands[kind]; exists {
		return fmt.Errorf("command already exists: %v", kind)
	}
	s.commands[kind] = command
	s.commandTypes = append(s.commandTypes, kind)
	return nil
}

// FetchCommand takes in a struct pointer and sets the value of that pointer
// to a command currently stored in the command registry. This ensures the input argument is
// set to the right pointer that refers to the originally registered command.
func (s *CommandRegistry) FetchCommand(command interface{}) error {
	if reflect.TypeOf(command).Kind() != reflect.Ptr {
		return fmt.Errorf("input must be of pointer type, received value type instead: %T", command)
	}
	element := reflect.ValueOf(command).Elem()
	if running, ok := s.commands[element.Type()]; ok {
		element.Set(reflect.ValueOf(running))
		return nil
	}
	return fmt.Errorf("unknown command: %T", command)
}

// ExecuteCommand takes in command, looks up the command in the registry, and calls the commands Execute() function.
func (s *CommandRegistry) ExecuteCommand(command iCommand, args ...string) error {
	kind := reflect.TypeOf(command)
	if _, exists := s.commands[kind]; exists {
		if err := s.commands[kind].Execute(args[0:]...); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("unknown command: %T", command)
	}
	return nil
}

// Execute takes in a command string, and args, attempts to find the command by its name, and then Execute() with args.
func (s *CommandRegistry) Execute(name string, args ...string) error {

	if name == "help" {
		s.Help()
		return nil
	}

	found := false

	for _, kind := range s.commandTypes {
		if s.commands[kind].GetName() == name {
			found = true
			if err := s.commands[kind].Execute(args[0:]...); err != nil {
				return err
			}
			return nil
		}
	}

	if !found {
		return fmt.Errorf("unknown command: %s", name)
	}

	return nil
}

// Help command to show help text of all commands
func (s *CommandRegistry) Help() {
	fmt.Println()
	helpCmd := ColorYellow + "help" + ColorReset
	fmt.Printf("%25v : %s\n", helpCmd, "Shows a list of all commands, and a short help text.")
	for _, kind := range s.commandTypes {
		cmd := ColorYellow + s.commands[kind].GetName() + ColorReset
		fmt.Printf("%25v : %s\n", cmd, s.commands[kind].GetHelpText())
	}
	fmt.Println()
}
