package rsa

import (
	"testing"
)

func TestUnmarshalPrivateKeyFromFile(t *testing.T) {
	_, err := UnmarshalPrivateKeyFromFile("testdata/sample_key.priv")
	if err != nil {
		t.Fatal("Unable to read private key", err)
	}
}

func TestUnmarshalPublicKeyFromFile(t *testing.T) {
	_, err := UnmarshalPublicKeyFromFile("testdata/sample_key.pub")
	if err != nil {
		t.Fatal("Unable to read public key", err)
	}
}

func TestGenerateRandomString(t *testing.T) {
	s, err := GenerateRandomString(16)
	if err != nil {
		t.Fatal("Unable to generate random string", err)
	}
	if len(s) != 32 {
		t.Fatalf("Invalid random string; expected: %d, got %d", 32, len(s))
	}
}
