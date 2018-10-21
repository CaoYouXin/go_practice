package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(os.Args)
	fmt.Println(os.Args[1:])

	wordPtr := flag.String("foo", "bar", "foo bar")
	var aVar int
	flag.IntVar(&aVar, "i", 0, "a int")

	flag.Parse()
	fmt.Println(flag.Args())
	fmt.Println("word:", *wordPtr)
	fmt.Println("int:", aVar)

	for _, v := range os.Environ() {
		pair := strings.Split(v, "=")
		fmt.Println(pair)
	}
}
