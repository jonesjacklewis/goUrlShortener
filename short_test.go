package main

import (
	"testing"
)

func TestShort(t *testing.T) {
	t.Run("Should convert a URL to a hash string", func(t *testing.T) {
		url_str := "https://www.google.com"
		shortened := Short(url_str)

		if shortened == url_str {
			t.Errorf("Short() = %s; want hash", shortened)
		}

	})
}
