package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestShort(t *testing.T) {
	t.Run("Should convert a URL to a hash string", func(t *testing.T) {
		url_str := "https://www.google.com"
		shortened := short(url_str)

		if shortened == url_str {
			t.Errorf("Short() = %s; want hash", shortened)
		}

	})

	t.Run("Same url should produce same hash", func(t *testing.T) {
		url_str := "https://www.google.com"
		shortened := short(url_str)

		turns := 100

		for turns > 0 {
			turns -= 1

			round_hash := short(url_str)

			if round_hash != shortened {
				t.Errorf("Short() = %s; want %s", round_hash, shortened)
			}
		}

	})

	t.Run("An invalid url should return a 'N/a' as a literal string", func(t *testing.T) {
		url_str := "hello world"
		shortened := short(url_str)

		if shortened != "N/a" {
			t.Errorf("Short() = %s; want 'N/a'", shortened)
		}
	})

	t.Run("Given a known hash, should be able to get the url", func(t *testing.T) {
		url_str := "https://www.google.com"
		shortened := short(url_str)

		long_form := long(shortened)

		if long_form != url_str {
			t.Errorf("Long() = %s; want %s", long_form, url_str)
		}
	})

	t.Run("Given an unknown hash, should return 'N/a' as a literal string", func(t *testing.T) {
		long_form := long("helloWorld")

		if long_form != "N/a" {
			t.Errorf("Long() = %s; want N/a", long_form)
		}
	})

	t.Run("Given a valid url, should return true", func(t *testing.T) {
		valid_url := "https://www.google.com"

		if !isValidUrl(valid_url) {
			t.Errorf("isValidUrl() = false; want true")
		}
	})

	t.Run("Given an ivalid url, should return true", func(t *testing.T) {
		valid_url := "hello world"

		if isValidUrl(valid_url) {
			t.Errorf("isValidUrl() = true; want false")
		}
	})

	t.Run("Given a none post request on /shorten, should return an error", func(t *testing.T) {
		w := httptest.NewRecorder()

		r := httptest.NewRequest(http.MethodGet, "/shorten", strings.NewReader(`{"url": "https://www.google.com/"}`))
		r.Header.Set("Content-Type", "application/json")

		shortenHandler(w, r)

		resp := w.Result()

		if resp.StatusCode != http.StatusMethodNotAllowed {
			t.Errorf("shortenHandler() = %d; want %d", resp.StatusCode, http.StatusMethodNotAllowed)
		}

		// check error message
		expectedMessage := "Invalid Method"
		if !strings.Contains(w.Body.String(), expectedMessage) {
			t.Errorf("shortenHandler() = %s; want %s", w.Body.String(), expectedMessage)
		}

	})

	t.Run("Given invalid json /shorten, should return an error", func(t *testing.T) {
		w := httptest.NewRecorder()

		r := httptest.NewRequest(http.MethodPost, "/shorten", strings.NewReader(``))
		r.Header.Set("Content-Type", "application/json")

		shortenHandler(w, r)

		resp := w.Result()

		if resp.StatusCode != http.StatusBadRequest {
			t.Errorf("shortenHandler() = %d; want %d", resp.StatusCode, http.StatusBadRequest)
		}

		expectedMessage := "Invalid Input"
		if !strings.Contains(w.Body.String(), expectedMessage) {
			t.Errorf("shortenHandler() = %s; want %s", w.Body.String(), expectedMessage)
		}

	})

	t.Run("Given no url parameter on /shorten, should return an error", func(t *testing.T) {
		w := httptest.NewRecorder()

		r := httptest.NewRequest(http.MethodPost, "/shorten", strings.NewReader(`{"uri": "https://www.google.com/"}`))
		r.Header.Set("Content-Type", "application/json")

		shortenHandler(w, r)

		resp := w.Result()

		if resp.StatusCode != http.StatusBadRequest {
			t.Errorf("shortenHandler() = %d; want %d", resp.StatusCode, http.StatusBadRequest)
		}

		expectedMessage := "Missing Url Parameter"
		if !strings.Contains(w.Body.String(), expectedMessage) {
			t.Errorf("shortenHandler() = %s; want %s", w.Body.String(), expectedMessage)
		}

	})

	t.Run("Given invalid url parameter on /shorten, should return an error", func(t *testing.T) {
		w := httptest.NewRecorder()

		val := "hello world"
		json_string := fmt.Sprintf(`{"url": "%s"}`, val)

		r := httptest.NewRequest(http.MethodPost, "/shorten", strings.NewReader(json_string))
		r.Header.Set("Content-Type", "application/json")

		shortenHandler(w, r)

		resp := w.Result()

		if resp.StatusCode != http.StatusBadRequest {
			t.Errorf("shortenHandler() = %d; want %d", resp.StatusCode, http.StatusBadRequest)
		}

		expectedMessage := fmt.Sprintf("%s is an invalid URL", val)
		if !strings.Contains(w.Body.String(), expectedMessage) {
			t.Errorf("shortenHandler() = %s; want %s", w.Body.String(), expectedMessage)
		}

	})

	t.Run("Given valid url on /shorten, should return a success with hash", func(t *testing.T) {
		w := httptest.NewRecorder()

		val := "https://www.google.com/"
		json_string := fmt.Sprintf(`{"url": "%s"}`, val)

		r := httptest.NewRequest(http.MethodPost, "/shorten", strings.NewReader(json_string))
		r.Header.Set("Content-Type", "application/json")

		shortenHandler(w, r)

		resp := w.Result()

		if resp.StatusCode != http.StatusOK {
			t.Errorf("shortenHandler() = %d; want %d", resp.StatusCode, http.StatusOK)
		}

		// response should be a json object with key 'short'
		expectedMessage := fmt.Sprintf(`"short":"http://localhost:8080/long/%s"`, short(val))

		if !strings.Contains(w.Body.String(), expectedMessage) {
			t.Errorf("shortenHandler() = %s; want %s", w.Body.String(), expectedMessage)
		}

	})

	t.Run("Given none GET request on /long/, should return error", func(t *testing.T) {
		w := httptest.NewRecorder()

		r := httptest.NewRequest(http.MethodPost, "/long/", nil)

		longHandler(w, r)

		resp := w.Result()

		if resp.StatusCode != http.StatusMethodNotAllowed {
			t.Errorf("longHandler() = %d; want %d", resp.StatusCode, http.StatusMethodNotAllowed)
		}

		// check error message
		expectedMessage := "Invalid Method"
		if !strings.Contains(w.Body.String(), expectedMessage) {
			t.Errorf("longHandler() = %s; want %s", w.Body.String(), expectedMessage)
		}
	})

	t.Run("Given unseen hash on /long/ should return erorr", func(t *testing.T) {
		w := httptest.NewRecorder()

		r := httptest.NewRequest(http.MethodGet, "/long/hello-world", nil)

		longHandler(w, r)

		resp := w.Result()

		if resp.StatusCode != http.StatusNotFound {
			t.Errorf("longHandler() = %d; want %d", resp.StatusCode, http.StatusNotFound)
		}

		// check error message
		expectedMessage := "Invalid Hash"
		if !strings.Contains(w.Body.String(), expectedMessage) {
			t.Errorf("longHandler() = %s; want %s", w.Body.String(), expectedMessage)
		}
	})

	t.Run("Given seen hash on /long/ should redirect to a url", func(t *testing.T) {
		w := httptest.NewRecorder()

		target := "https://www.google.com/"
		hash := short(target)

		r := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/long/%s", hash), nil)

		longHandler(w, r)

		resp := w.Result()

		if resp.StatusCode != http.StatusPermanentRedirect {
			t.Errorf("longHandler() = %d; want %d", resp.StatusCode, http.StatusPermanentRedirect)
		}

		redirectTarget := resp.Header.Get("Location")

		if redirectTarget != target {
			t.Errorf("longHandler() = %s; want %s", redirectTarget, target)
		}

	})

}
