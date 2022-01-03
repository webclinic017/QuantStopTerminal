package commands

import (
	"bufio"
	"fmt"
	"github.com/quantstop/quantstopterminal/internal/config"
	"io"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

type LogmonCommand struct {
	Command
}

func CreateLogmonCommand() *LogmonCommand {
	base := Command{
		name: "logmon",
		help: "Monitors the bot logs as they come in.",
	}
	cmd := &LogmonCommand{Command: base}
	return cmd
}

func (c *LogmonCommand) GetName() string {
	return c.name
}

func (c *LogmonCommand) GetHelpText() string {
	return c.help
}

func (c *LogmonCommand) Execute(args ...string) error {
	tailLogs(os.Stdout)
	return nil

}

func tailLogs(out io.Writer) {

	// Create a new ticker
	t := time.NewTicker(time.Second)

	// Try opening log file
	logFile, err := os.Open(config.LocalConfig("QuantstopTerminal") + "/logs/log.txt")
	if err != nil {
		panic(err)
	}

	// Make signal channel
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Create a new reader for the file
	fileReader := bufio.NewReader(logFile)
	info, err := logFile.Stat()
	if err != nil {
		panic(err)
	}

	// Set a variable that stores the size of the file, so we can compare and update the output if new stuff is added
	oldSize := info.Size()

	// Defer stopping the ticker until finished
	defer func() {
		t.Stop()
	}()

	// Goroutine to handle receiving an interrupt, and sending a done message to channel
	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	// Docs: https://tour.golang.org/concurrency/5
	// This is a blocking loop, that will run the channel tick case indefinitely until an interrupt is received
	for {
		select {
		case <-done: // if channel message is shutdown finish loop
			return
		case <-t.C: // on channel tick read the file and print logs

			// This will read one line at a time, and then pretty print the line
			for line, prefix, err := fileReader.ReadLine(); err != io.EOF; line, prefix, err = fileReader.ReadLine() {
				prettyPrintLogLine(out, string(line), prefix)
			}

			// Set position in file
			pos, err := logFile.Seek(0, io.SeekCurrent)
			if err != nil {
				fmt.Println(err)
				return
			}

			// Get new file info
			newInfo, err := logFile.Stat()
			if err != nil {
				fmt.Println(err)
				return
			}
			newSize := newInfo.Size()

			// Compare new size to previous size
			if newSize != oldSize {

				// If newSize is less than previous size, reset current position, otherwise seek
				if newSize < oldSize {
					logFile.Seek(0, io.SeekStart)
				} else {
					// Set current position
					logFile.Seek(pos, io.SeekStart)
				}

				// Set new fileReader, and oldSize
				fileReader = bufio.NewReader(logFile)
				oldSize = newSize
			}

		}
	}
}

func prettyPrintLogLine(out io.Writer, line string, prefix bool) {

	logLineArray := strings.Split(line, " | ")

	var newLine string

	switch logLineArray[0] {
	case "[INFO]":
		newLine = line
	case "[WARN]":
		newLine = ColorYellow + line + ColorReset
	case "[ERROR]":
		newLine = ColorRed + line + ColorReset
	case "[DEBUG]":
		newLine = ColorGreen + line + ColorReset
	default:
		newLine = line
	}

	if prefix {
		fmt.Fprint(out, newLine)
	} else {
		fmt.Fprintln(out, newLine)
	}

}
