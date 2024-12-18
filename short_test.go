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

}
