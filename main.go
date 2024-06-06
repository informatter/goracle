package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)


func useOllama(userPrompt string) string {
	model :="llama3"
	return request(model,systemWindowsCmd,userPrompt)

}

func main(){

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("\n\nWelcome to goracle 🔮 !\n\nInteract in natural language with your shell and goracle 🔮 will respond with the appropiate shell command\n\n")
	fmt.Print("For example if you type:\n\n'Display all the contents of my current directory', goracle 🔮 should respond:\n\nls\n\n")
	fmt.Print("To execute any command simply append exec at the beginning:\n\n")
	fmt.Print("exec ls\n\n")
	
	for {
		fmt.Print("goracle 🔮💠 > ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		err = processUserInput(input)

		if err != nil {

			fmt.Fprintln(os.Stderr, err)
		}

	}

}

func processUserInput(commandStr string) error {

	commandStr = strings.TrimSpace(commandStr)

	if (commandStr == "exit"){
		os.Exit(0)
	}else if  (strings.Contains(commandStr,"exec")){
		// run coommand 

		arrCommandStr := strings.Fields(commandStr)[1:] // removes exec
		cmd := exec.Command(arrCommandStr[0], arrCommandStr[1:]...)
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		result := cmd.Run()

		return result
	}

	response:= useOllama(commandStr)
	fmt.Print(response)
	return nil

}