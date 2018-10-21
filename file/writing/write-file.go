package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	fp := "-test.log"
	rand.NewSource(int64(time.Now().Nanosecond()))
	f1p := strconv.FormatInt(rand.Int63(), 10) + fp
	ioutil.WriteFile(f1p, []byte("this is whate i want"), 0644)
	data, _ := ioutil.ReadFile(f1p)
	fmt.Println(string(data))

	rand.NewSource(int64(time.Now().Nanosecond()))
	f2p := strconv.FormatInt(rand.Int63(), 10) + fp
	f, err := os.Create(f2p)
	check(err)
	defer f.Close()

	f.Write([]byte("first line\n"))
	f.WriteString("second line\n")
	f.Sync()

	w := bufio.NewWriter(f)
	w.WriteString("last line\n")
	w.Flush()

	data, _ = ioutil.ReadFile(f2p)
	fmt.Println(string(data))
}
