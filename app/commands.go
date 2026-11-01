package main

import (
	"os"
	"strings"
)

type CallbackFunc func([]string) string

type Command struct {
	Name     string
	Callback CallbackFunc
}

func getCommands() map[string]Command {
	commandList := map[string]Command{
		"type": {
			Name:     "type",
			Callback: commandType,
		},

		"exit": {
			Name:     "exit",
			Callback: commandExit,
		},
		"echo": {
			Name:     "echo",
			Callback: commandEcho,
		},
	}
	return commandList
}

func commandExit(string[]) string {
	os.Exit(0)
	return ""
}

func commandEcho(input string[]) string {
	return strings.Join(nput, " ")
}
func commandExit(string[]) string {
	os.Exit(0)
	return ""
}
