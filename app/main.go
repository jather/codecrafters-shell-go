package main

import (
	"bufio"
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
				path, err := exec.LookPath(firstWord)
				if err != nil {
					fmt.Println(firstWord + ": command not found")
				} else {
					cmd := exec.Command(path, strings.Join(args, " "))
					cmd.Stdin = os.Stdin
					cmd.Stdout = os.Stdout
					cmd.Stderr = os.Stderr
					if err = cmd.Run(); err != nil {
						fmt.Println(err.Error())
					}
				}
			} else {
				command.Callback(&cfg, args[1:])
			}
		}
	}
}
