package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

var knownHashes map[string]string = make(map[string]string)

func isValidUrl(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)

	return err == nil
}

func short(url string) string {

	if !isValidUrl(url) {
		return "N/a"
	}

	hash := base64.StdEncoding.EncodeToString([]byte(url))

	knownHashes[hash] = url

	return hash
}

func long(hash string) string {
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
		http.Error(w, "Invalid Input", http.StatusBadRequest)
		return
	}

	val, ok := req["url"]

	if !ok {
		http.Error(w, "Missing Url Parameter", http.StatusBadRequest)
		return
	}

	hash := short(val)

	if hash == "N/a" {
		http.Error(w, fmt.Sprintf("%s is an invalid URL", val), http.StatusBadRequest)
		return
	}

	resp := map[string]string{"short": fmt.Sprintf("http://localhost:8080/long/%s", hash)}
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(resp)
}

func longHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
		return
	}

	hash := r.URL.Path
	hash = strings.TrimPrefix(hash, "/long/")

	url := long(hash)

	if url == "N/a" {
		http.Error(w, "Invalid hash", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, url, 301)
}

func main() {
	http.HandleFunc("/shorten", shortenHandler)
	http.HandleFunc("/long/", longHandler)
	http.ListenAndServe(":8080", nil)
}
