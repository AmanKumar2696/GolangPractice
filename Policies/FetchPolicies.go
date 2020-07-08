package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd, _ := exec.Command("bash", "script.sh").Output()
	fmt.Println(string(cmd))

}
