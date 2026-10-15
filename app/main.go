package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Print("$ ")
	command, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		log.Fatalf("failed to read string")
	} else {
		fmt.Print(command[:len(command)-1] + ": command not found")
	}
}
