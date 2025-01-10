package main

import (
	"bufio"
	"fmt"
	"os"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func main() {
	// Uncomment this block to pass the first stage
	fmt.Fprint(os.Stdout, "$ ")

	// Wait for user input
	userInput, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println("invalid_command")
		fmt.Println(fmt.Sprintf("invalid_command:%v not found", userInput))
	}
	fmt.Println(userInput)

}
