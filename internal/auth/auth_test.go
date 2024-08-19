package auth

import (
	"net/http"
	"testing"
)

func TestBadAuth(t *testing.T) {
	h := http.Header{}
	h.Add("Authorizaiton", "Badapikey")
	key, err := GetAPIKey(h)
	if err == nil {
		t.Fatalf("expected: error, got %s", key)
	}
}

func TestAuth(t *testing.T) {
	h := http.Header{}
	h.Add("Authorization", "ApiKey blahblahblah")
	key, err := GetAPIKey(h)
	if err != nil {
		t.Fatal(err.Error())
	}
	if key != "blahblahblah" {
		t.Fatalf("expected: blahblahblah, got: %s", key)
	}
}
