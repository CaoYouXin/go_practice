package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	seed := time.Now().UnixNano()
	mrand := rand.New(rand.NewSource(seed))
	fmt.Println(mrand.Intn(100))
}
