package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func echocmd(first []string) error {
	if len(first) == 0 {
		fmt.Fprintln(os.Stdout)
		return nil

	}
	for i := 0; i < len(first)-1; i++ {
		fmt.Fprintf(os.Stdout, "%s ", first[i])
	}
	fmt.Fprintln(os.Stdout, first[len(first)-1])
	return nil
}
func Cdcmd(first []string) string {
	tomove := first[1]
	//source _ := os.Getwd()
	homedir, _ := os.UserHomeDir()
	if tomove == "~" {
		os.Chdir(homedir)
	} else {
		err := os.Chdir(tomove)
		if err != nil {
			return fmt.Sprintf("cd: %s: No such file or directory\n", tomove)
		}
	}
	return ""
}
func typecmd(path string, first []string) string {
	builtintype := []string{"echo", "exit", "type", "pwd", "cd"}
	for _, item := range builtintype {
		if first[1] == item {
			return fmt.Sprintf("%s is a shell builtin", first[1])
		}
	}
	directory := strings.Split(path, ":")
	for _, item := range directory {
		fullPath := fmt.Sprintf("%s/%s", item, first[1])
		if _, err := os.Stat(fullPath); err == nil {
			return fmt.Sprintf("%s is %s", first[1], fullPath)
		}
	}

	return fmt.Sprintf("%s: not found", first[1])
}
func main() {
	// Uncomment this block to pass the first stage
	fmt.Fprint(os.Stdout, "$ ")
	path := os.Getenv("PATH")

	userInput := bufio.NewReader(os.Stdin)

	for {
		command, err := userInput.ReadString('\n')

		if err != nil {
			fmt.Println("invalid_command: command not found")
			os.Exit(1)
		}
		command = strings.TrimSpace(command)

		var first []string
		for {
			start := strings.Index(command, "'")
			if start == -1 {
				first = append(first, strings.Fields(command)...)
				break
			}
			first = append(first, strings.Fields(command[:start])...)
			command = command[start+1:]
			end := strings.Index(command, "'")
			token := command[:end]
			first = append(first, token)
			command = command[end+1:]
		}
		if len(first) == 1 && first[0] != "pwd" {
			fmt.Printf("%s: command not found\n", first[0])
		} else if first[0] == "pwd" {
			dir, err := os.Getwd()
			if err != nil {
				fmt.Println("error")
			}
			fmt.Println(dir)

		} else if first[0] == "echo" {
			echocmd(first)

		} else if first[0] == "type" {
			fmt.Println(typecmd(path, first))

		} else if first[0] == "cd" {
			fmt.Print(Cdcmd(first))

		} else if command == "exit 0" {
			os.Exit(0)

		} else {
			paths := strings.Split(path, ":")
			isfound := false
			for _, item := range paths {
				fullpath := filepath.Join(item, first[0])
				if _, err := os.Stat(fullpath); err == nil {
					isfound = true

					// Execute the command
					cmd := exec.Command(first[0], first[1:]...)
					cmd.Stdout = os.Stdout
					cmd.Stderr = os.Stderr
					if err := cmd.Run(); err != nil {
						fmt.Printf("%s: command not found\n", first[0])
					}
					break
				}
			}
			if !isfound {
				fmt.Printf("%s: not found\n", first[1])
			}
		}
		fmt.Fprint(os.Stdout, "$ ")
	}
}
