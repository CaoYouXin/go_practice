package main

import (
	"fmt"
	"strconv"
)

func main() {
	f, _ := strconv.ParseFloat("1.2345", 64)
	fmt.Printf("%T, %v\n", f, f)

	i, _ := strconv.ParseInt("1", 0, 64)
	fmt.Printf("%T, %v\n", i, i)

	ui, e := strconv.ParseUint("-1", 0, 64)
	fmt.Printf("%T, %v, %v\n", ui, ui, e)

	num, e := strconv.Atoi("-1")
	fmt.Printf("%T, %v, %v\n", num, num, e)
}
