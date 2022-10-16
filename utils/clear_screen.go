package utils

import (
	"fmt"
	"os"
	"os/exec"
)

func ClearScreen() {
	fmt.Scanln()
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func DirectClearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
