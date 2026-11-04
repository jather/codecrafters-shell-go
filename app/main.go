package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

const prompt = "$"

func main() {
	reader := bufio.NewReader(os.Stdin)
	commands := getCommands()
	cfg := Config{Commands: commands}
	for {
		fmt.Print(prompt + " ")

		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("failed to read string")
		} else {
			args := strings.Fields(input)
			if len(args) == 0 {
				continue
			}
			firstWord := args[0]
			command, ok := cfg.Commands[firstWord]
			if !ok {
				cmd := exec.Command(firstWord, args[1:]...)
				cmd.Stdin = os.Stdin
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				err = cmd.Run()
				var exitErr *exec.ExitError
				switch {
				case err == nil:
				case errors.Is(err, exec.ErrNotFound):
					fmt.Println(firstWord + ": command not found")
				case errors.As(err, &exitErr):
					//executable exited with failure
				default:
					fmt.Fprintln(os.Stderr, err)
				}
			} else {
				command.Callback(&cfg, args[1:])
			}
		}
	}
}
