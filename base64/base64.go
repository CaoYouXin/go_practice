package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := "abc123!?$*&()'-=@~"
	fmt.Println(data)

	encoded := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(encoded)
	decoded, _ := base64.StdEncoding.DecodeString(encoded)
	fmt.Println(string(decoded))

	encoded = base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(encoded)
	decoded, _ = base64.URLEncoding.DecodeString(encoded)
	fmt.Println(string(decoded))
}
