package game

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"unicode"
)

func OnlyLetters(s string) bool {
	for _, char := range s {
		if !('a' <= char && char <= 'z' || 'A' <= char && char <= 'Z') {
			return false
		}
	}
	return true
}
func ClearConsole() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func IsValidName(name string, maxLetters int) bool {
	if len(name) > maxLetters {
		return false
	}

	for _, char := range name {
		if !unicode.IsLetter(char) {
			return false
		}
	}

	return true
}

func AskForUserChoice(prompt string) int {
	var choice int
	fmt.Println(prompt)
	_, err := fmt.Scan(&choice)
	if err != nil {
		fmt.Println("Error reading user input:", err)
	}
	return choice
}
