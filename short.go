package main

import (
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

var knownHashes map[string]string = make(map[string]string)
var filename string = "short.db"

func addHash(path string, hash string, target string) bool {
	if !createDatabase(path) {
		return false
	}

	db, err := sql.Open("sqlite3", path)

	if err != nil {
		return false
	}

	defer db.Close()

	hash = strings.TrimSpace(hash)

	if len(hash) == 0 {
		return false
	}

	if !isValidUrl(target) {
		return false
	}

	insertQuerySql := `
	INSERT INTO
	urls (
		hash,
		url
	)
	VALUES (
		?,
		?
	);
	`

	_, err = db.Exec(insertQuerySql, hash, target)

	return err == nil
}

func createDatabase(fn string) bool {
	db, err := sql.Open("sqlite3", fn)

	if err != nil {
		return false
	}

	defer db.Close()

	createTableSql := `
	CREATE TABLE
	IF NOT EXISTS
	urls (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		hash TEXT NOT NULL UNIQUE,
		url TEXT NOT NULL
	);
	`

	_, err = db.Exec(createTableSql)

	return err == nil
}

func isValidUrl(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)
	return err == nil
}

func short(url string) string {

	if !isValidUrl(url) {
		return "N/a"
	}

	hash := base64.StdEncoding.EncodeToString([]byte(url))

	addHash(filename, hash, url)

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
		http.Error(w, "Invalid Hash", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, url, http.StatusPermanentRedirect)
}

func main() {
	createDatabase(filename)
	http.HandleFunc("/shorten", shortenHandler)
	http.HandleFunc("/long/", longHandler)
	http.ListenAndServe(":8080", nil)
}
