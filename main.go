package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/fatih/color"
)

func useOllama(userPrompt string) string {

	model := "llama3"

	go loading(responseChan)
	response := request(model, systemPromptBash, userPrompt)
	responseChan <- "request-finished"
	color.Set(color.FgWhite)
	color.Set(color.BgBlack)
	fmt.Print("\n\n")
	return response
}

func loading(responseChan chan string) {
	goracleWhite := color.New(color.BgHiCyan)
	fmt.Print("\n")
	for {
		select {
		case msg := <-responseChan:
			switch msg {
			case "request-finished":
				return
			}
		default:
			goracleWhite.Print(" ")
			time.Sleep(700 * time.Millisecond)
		}
	}

}

func main() {

	reader := bufio.NewReader(os.Stdin)
	//processUserInput("exec clear",reader)

	fmt.Print("\n\nWelcome to goracle 🔮 !\n\nInteract in natural language with your shell and goracle  will respond with the appropiate shell command\n\n")
	fmt.Print("For example if you type:\n\n'Display all the contents of my current directory', goracle  should respond:\n\nls\n\n")
	fmt.Print("To execute any command simply append exec at the beginning:\n\n")
	fmt.Print("exec ls\n\n")

	for {

		displayGoracleSigntature("")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		processUserInput(input, reader)

	}

}

func executeCommand(command []string) {
	color.Set(color.FgYellow)
	cmd := exec.Command(command[0], command[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	color.Set(color.FgWhite)
}

func executeBashCommand(command string) {
	color.Set(color.FgYellow)
	cmd := exec.Command("sh", "-c", command)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	color.Set(color.FgWhite)
}
func processUserInput(commandStr string, reader *bufio.Reader) {

	commandStr = strings.TrimSpace(commandStr)

	if commandStr == "exit" {
		os.Exit(0)
	} else if strings.Contains(commandStr, "exec") {

		arrCommandStr := strings.Fields(commandStr)[1:] // removes exec

		command := strings.Join(arrCommandStr, " ")
		executeBashCommand(command)
		return
	}

	response := useOllama(commandStr)
	if response != "" {
		color.Set(color.FgGreen)
		color.Set(color.BgBlack)
		fmt.Print(response)
		fmt.Print("\n")
		// goracleWhite := color.New(color.FgHiWhite)
		// goracleWhiteBold := goracleWhite.Add(color.Bold)
		// goracleWhiteBold.Print("goracle 🔮 > execute? [y/n] ")
		displayGoracleSigntature(executePrompt)

		for {
			input, err := reader.ReadString('\n')
			input = strings.TrimSpace(strings.ToLower(input))
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}

			if input == "y" || input == "yes" {
				executeBashCommand(response) 
				return
			}
			if input == "n" || input == "no" {
				return
			}
			if input == "exit" {
				os.Exit(0)
			}

			return

		}
	}
	color.Set(color.FgWhite)

}
