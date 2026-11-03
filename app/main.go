package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
				fmt.Println(firstWord + ": command not found")
			} else {
				command.Callback(&cfg, args[1:])
			}
		}
	}
}
