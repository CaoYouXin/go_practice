package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	s := "this is a string"

	h := sha1.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)

	fmt.Printf("sha1 of '%v' is '%x'\n", s, bs)
}
