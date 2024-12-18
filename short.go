package main

import "crypto/sha256"

func Short(url string) string {

	h := sha256.New()

	h.Write([]byte(url))

	hash := string(h.Sum(nil)[:])

	return hash
}

func main() {
	hash := Short("https://www.google.com")

	println(hash)
}
