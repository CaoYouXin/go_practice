package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {
	binary, _ := exec.LookPath("ls")
	args := []string{"ls", "-a", "-l"}
	syscall.Exec(binary, args, os.Environ())
}
