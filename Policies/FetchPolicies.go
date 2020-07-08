package main

import (
	"os"
	"os/exec"
)

func main() {
	//cmd, _ := exec.Command("bash", "script.sh").Output()
	//fmt.Println(string(cmd))
	get()
}

func get() {
	path, _ := exec.LookPath("bash")
	cmd := &exec.Cmd{
		Path:   path,
		Args:   []string{"./", "script.sh"},
		Stdout: os.Stdout,
		Stdin:  os.Stdin,
	}
	cmd.Start()
	cmd.Wait()
}
