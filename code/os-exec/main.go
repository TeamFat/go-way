package main

import (
	"fmt"
	"os/exec"
)

func main() {
	pwdCmd := exec.Command("ls")
	pwdOutput, _ := pwdCmd.Output()
	fmt.Println(string(pwdOutput))
}
