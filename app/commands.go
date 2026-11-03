package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type CallbackFunc func(cfg *Config, input []string) error

type Command struct {
	Name     string
	Callback CallbackFunc
}
type Config struct {
	Commands map[string]Command
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
		return nil
	}
	first := input[0]
	_, found := cfg.Commands[first]
	if found {
		fmt.Println(first + " is a shell builtin")
	} else if path, err := exec.LookPath(first); err == nil {
		fmt.Println(first + " is " + path)
	} else {
		fmt.Println(first + ": not found")
	}
	return nil
}
