package main

import (
	"bufio"
	"fmt"
	"github.com/quantstop/quantstopterminal/internal/config"
	"github.com/quantstop/quantstopterminal/internal/engine"
	"github.com/quantstop/quantstopterminal/internal/grpcserver"
	"github.com/quantstop/quantstopterminal/internal/grpcserver/auth"
	"github.com/quantstop/quantstopterminal/internal/qstcli/commands"
	"golang.org/x/net/context"
	"golang.org/x/term"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"time"
)

var (
	registry *commands.CommandRegistry
	certPath = filepath.Join(config.LocalConfig("QuantstopTerminal"), "tls", "cert.pem")
)

const (
	defaultTimeout = time.Second * 30
	defaultHost    = "localhost:9052"
)

func main() {

	// Get username and password from user
	un, psw, err := getCredentials()
	if err != nil {
		log.Fatalf("error getting user input: %v", err)
	}

	// Get tls certificate
	tlsCredentials, err := credentials.NewClientTLSFromFile(certPath, "")
	if err != nil {
		log.Fatalf("error getting cert: %v", err)
	}

	// Set up a context and cancel function for the gRPC connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	// Set dial options for the gRPC connection
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(tlsCredentials),
		grpc.WithPerRPCCredentials(auth.BasicAuth{
			Username: un,
			Password: psw,
		}),
	}

	// Try dialing gRPC server
	ctx, cancel = context.WithTimeout(ctx, defaultTimeout)
	conn, err := grpc.DialContext(ctx, defaultHost, opts...)
	if err != nil {
		log.Fatalf("error connecting: %v", err)
	}

	// Defer closing the gRPC connection till we exit the application
	defer func(conn *grpc.ClientConn) {
		cancel()
		_ = conn.Close()
	}(conn)

	// Set global client
	client := grpcserver.NewGRPCServerClient(conn)

	// Try getting info from server using connection above
	resp, err := client.GetInfo(ctx, &grpcserver.GetInfoRequest{})
	if err != nil {
		// If there is an error here, log fatal because it means we are not authenticated or the server is down
		log.Fatalf("%v: ", err)
	}

	// If we get here, the request was successful, which means we authenticated successfully
	// We can now continue setting up, and showing the prompt to the user

	// Setup command registry
	registry = commands.NewCommandRegistry()
	if err = registry.InitCommands(client); err != nil {
		log.Fatalf("error setting up command registry: %v", err)
	}

	// Print banner and version
	fmt.Println("")
	fmt.Println(commands.ColorPurple + engine.BannerGraffiti + commands.ColorReset)
	fmt.Println(resp.Version)

	fmt.Println("")

	// Print some system info
	fmt.Println("System Uptime:", resp.Uptime)
	fmt.Println("")

	// Create a new reader for stdin
	reader := bufio.NewReader(os.Stdin)

	// Infinite loop to wait for command input
	for {

		// Print the prompt to screen
		fmt.Print(commands.ColorRed, "qstcli> ", commands.ColorReset)

		// ReadString will block until the delimiter is entered
		cmdString, err := reader.ReadString('\n')
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
		}

		// Trim new line characters and create an array
		// The first element will be the command, and any following elements will be arguments
		commandStr := strings.TrimSuffix(cmdString, "\n")
		arrCommandStr := strings.Fields(commandStr)

		// Execute command with arguments
		if err := registry.Execute(arrCommandStr[0], arrCommandStr[1:]...); err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
		}

		// Todo: on windows (cmd), at the prompt if you enter tab it exits ...

	}

}

func getCredentials() (string, string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Username: ")
	username, err := reader.ReadString('\n')
	if err != nil {
		return "", "", err
	}

	fmt.Print("Enter Password: ")
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", "", err
	}

	password := string(bytePassword)
	return strings.TrimSpace(username), strings.TrimSpace(password), nil
}
