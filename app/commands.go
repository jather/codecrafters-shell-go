package main

import (
	"fmt"
	"os"
	"strings"
)

type CallbackFunc func(cfg *Config, input []string) error

type Command struct {
	Name     string
	Type     string
	Callback CallbackFunc
}
type Config struct {
	Commands map[string]Command
}

func getCommands() map[string]Command {
	commandList := map[string]Command{
		"type": {
			Name:     "type",
			Type:     "builtin",
			Callback: commandType,
		},

		"exit": {
			Name:     "exit",
			Type:     "builtin",
			Callback: commandExit,
		},

		"echo": {
			Name:     "echo",
			Type:     "builtin",
			Callback: commandEcho,
		},
	}
	return commandList
}

func commandExit(cfg *Config, input []string) error {
	os.Exit(0)
	return nil
}

func commandEcho(cfg *Config, input []string) error {
	fmt.Println(strings.Join(input, " "))
	return nil
}
func commandType(cfg *Config, input []string) error {
	if len(input) < 1 {
		fmt.Println(": not found")
	}
	first := input[0]
	command, ok := cfg.Commands[first]
	if !ok {
		fmt.Println(first + ": not found")
	} else {
		fmt.Println(first + " is a shell " + command.Type)
	}
	return nil
}
