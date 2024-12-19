package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
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

func shortenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
		return
	}

	var req map[string]string

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	val, ok := req["url"]

	if !ok {
		http.Error(w, "Missing url parameter", http.StatusBadRequest)
		return
	}

	hash := Short(val)

	resp := map[string]string{"short": fmt.Sprintf("http://localhost:8080/long/%s", hash)}
	w.Header().Set("Content-Type", "application.json")

	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/shorten", shortenHandler)
	http.ListenAndServe(":8080", nil)
}
