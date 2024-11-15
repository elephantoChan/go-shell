package main

import (
	utils "MyShell/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Completed entry to shell.")
	var IsCurrentlyRunning bool = true

	for IsCurrentlyRunning {

		fmt.Print("$ ")
		input, _ := reader.ReadString('\n')
		var CurrentInput = strings.Split(input, " ")

		switch strings.TrimSpace(CurrentInput[0]) {
		case "echo":
			fmt.Print(utils.Echo(CurrentInput[1:]...), "\n")
		case "exit":
			fmt.Print("Exiting..")
			IsCurrentlyRunning = false
		case "cat":
			content, err := utils.Cat(strings.TrimSpace(CurrentInput[1]))
			if err == nil {
				fmt.Print(content)
			} else {
				fmt.Println(err)
			}
		case "grep":
			content := utils.Grep([]byte(strings.TrimSpace(CurrentInput[1])), strings.TrimSpace(CurrentInput[2]))
			fmt.Print(content)
		default:
			fmt.Println("error: unknown command", strings.TrimSpace(CurrentInput[0]))
		}
	}
}
