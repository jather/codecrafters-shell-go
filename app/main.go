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
	for {
		fmt.Print(prompt + " ")

		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("failed to read string")
		} else {
			args := strings.Fields(input)
			command := args[0]
			switch command {
			case "exit":
				os.Exit(0)
			case "echo":
				fmt.Println(strings.Join(args[1:], " "))
			default:
				fmt.Println(command + ": command not found")
			}
		}
	}
}
