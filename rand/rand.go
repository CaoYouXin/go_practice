package main

import (
	"fmt"
	mrand "math/rand"
	"time"
)

func main() {
	p := fmt.Println
	seed := time.Now().UnixNano()

	mrand.NewSource(seed)
	p(mrand.Intn(100))
}
