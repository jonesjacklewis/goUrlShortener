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

	t.Run("Same url should produce same hash", func(t *testing.T) {
		url_str := "https://www.google.com"
		shortened := Short(url_str)

		turns := 100

		for turns > 0 {
			turns -= 1

			round_hash := Short(url_str)

			if round_hash != shortened {
				t.Errorf("Short() = %s; want %s", round_hash, shortened)
			}
		}

	})

	t.Run("Given a known hash, should be able to get the url", func(t *testing.T) {
		url_str := "https://www.google.com"
		shortened := Short(url_str)

		long_form := Long(shortened)

		if long_form != url_str {
			t.Errorf("Long() = %s; want %s", long_form, url_str)
		}
	})

	t.Run("Given an unknown hash, should return 'N/a' as a literal string", func(t *testing.T) {
		long_form := Long("helloWorld")

		if long_form != "N/a" {
			t.Errorf("Long() = %s; want N/a", long_form)
		}
	})

}
