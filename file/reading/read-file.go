package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	fp := "/Users/cao/Dev/VarUsage.java"

	data, err := ioutil.ReadFile(fp)
	check(err)
	fmt.Println(string(data))

	f, err := os.Open(fp)
	check(err)
	defer f.Close()

	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1))

	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 5)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes form %d: %s\n", n2, o2, string(b2))

	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 5)
	n3, err := io.ReadAtLeast(f, b3, 5)
	check(err)
	fmt.Printf("%d bytes form %d: %s\n", n3, o3, string(b3))

	_, err = f.Seek(0, 0)
	check(err)

	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))
}
