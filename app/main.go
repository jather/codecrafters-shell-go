package main

import (
	"fmt"
)

func main() {
	var command string
	fmt.Println("$ ")
	fmt.Scan(&command)
	fmt.Printf("%s: command not found/n", command)
}
