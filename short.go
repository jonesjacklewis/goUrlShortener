package main

import (
	"encoding/base64"
	"fmt"
)

var knownHashes map[string]string = make(map[string]string)

func Short(url string) string {

	hash := base64.StdEncoding.EncodeToString([]byte(url))

	knownHashes[hash] = url

	return hash
}

func Long(hash string) string {
	val, ok := knownHashes[hash]

	if !ok {
		return "N/a"
	}

	return val
}

func main() {
	hash := Short("https://www.google.com")

	fmt.Println(hash)
}
