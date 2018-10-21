package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func main() {
	dateCmd := exec.Command("date")
	o, _ := dateCmd.Output()
	fmt.Println("> date")
	fmt.Println(string(o))

	lsCmd := exec.Command("bash", "-c", "ls -l -a -h")
	o, _ = lsCmd.Output()
	fmt.Println("> ls -l -a -h")
	fmt.Println(string(o))

	grepCmd := exec.Command("grep", "spawn")
	pipeIn, _ := grepCmd.StdinPipe()
	pipeOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	pipeIn.Write(o)
	pipeIn.Close()
	result, _ := ioutil.ReadAll(pipeOut)
	grepCmd.Wait()
	fmt.Println("> grep spawn")
	fmt.Println(string(result))
}
