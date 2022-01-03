package commands

import (
	"encoding/json"
	"fmt"
)

var ColorReset = "\033[0m"
var ColorRed = "\033[31m"
var ColorGreen = "\033[32m"
var ColorYellow = "\033[33m"
var ColorBlue = "\033[34m"
var ColorPurple = "\033[35m"
var ColorCyan = "\033[36m"
var ColorGray = "\033[37m"
var ColorWhite = "\033[97m"

type iCommand interface {
	GetName() string
	GetHelpText() string
	Execute(args ...string) error
}

type Command struct {
	name string
	help string
}

func JsonOutput(in interface{}) {
	j, err := json.MarshalIndent(in, "", "  ")
	if err != nil {
		return
	}
	fmt.Print(string(j))
	fmt.Println("")
}
