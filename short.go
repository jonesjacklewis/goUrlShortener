package main

import (
	"encoding/base64"
	"fmt"
)

func Short(url string) string {

	hash := base64.StdEncoding.EncodeToString([]byte(url))

	return hash
}

func main() {
	hash := Short("https://www.google.com")

	fmt.Println(hash)
}
