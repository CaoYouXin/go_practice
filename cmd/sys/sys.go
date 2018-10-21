package main

import (
	"os"

	"golang.org/x/sys/unix"
)

func main() {
	unix.Exec("/usr/local/bin/go", []string{"go", "run", "../../closing/closing.go"}, os.Environ())
}
