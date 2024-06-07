package main

import (
	"fmt"

	"github.com/fatih/color"
)

const executePrompt string = "execute? [y/n] "
const goracleIcon string = "🔮"

func goracleColor() *color.Color {
	goracleWhite := color.New(color.FgHiWhite)
	goracleWhiteBold := goracleWhite.Add(color.Bold)
	return goracleWhiteBold
}

func displayGoracleSigntature(message string) {
	goracleColor := goracleColor()
	formatedMessage := fmt.Sprintf("goracle %s > %s", goracleIcon, message)
	goracleColor.Print(formatedMessage)
}
