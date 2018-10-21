package main

import (
	"os"
	"os/exec"
	"syscall"
	"time"
)

func main() {
	binary, _ := exec.LookPath("ls")
	args := []string{"ls", "-a", "-l"}
	syscall.Exec(binary, args, os.Environ())

	binary, _ = exec.LookPath("go")
	args = []string{"go", "run", "../../closing/closing.go"}
	syscall.Exec(binary, args, os.Environ())

	time.Sleep(time.Second * 6)
}
