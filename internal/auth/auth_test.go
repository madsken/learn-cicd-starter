package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	noKeyHdr := http.Header{}
	malformedKey := http.Header{}
	malformedKey.Set("Authorization", "wrongkey")
	correctKey := http.Header{}
	correctKey.Set("Authorization", "ApiKey key")

	input := []http.Header{noKeyHdr, malformedKey, correctKey}

	want := []string{"", "", "key"}

	for i := range input {
		got, _ := GetAPIKey(input[i])
		if got != want[i] {
			t.Fatalf("Failed on %d. Want: %s, got: %s", i, want[i], got)
		}
	}
	t.Fatalf("woops")
}
