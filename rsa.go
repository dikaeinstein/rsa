package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	// BlockPublicKey contains the value which defines a Public Key in a PEM block
	BlockPublicKey = "PUBLIC KEY"
	// BlockPrivateKey contains the value which defines a PrivateKey in a PEM block
	BlockPrivateKey = "RSA PRIVATE KEY"
)

func readFileToBlock(file string, blockType string) (*pem.Block, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("Unable to open keyfile: %s", err)
	}
	defer f.Close()

	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("Unable to read data from file: %s", err)
	}

	block, _ := pem.Decode(bytes)
	if block == nil || block.Type != blockType {
		return nil, fmt.Errorf("Unable to decode %s key", blockType)
	}

	return block, nil
}

// UnmarshalPublicKeyFromFile reads the given file and returns an RSA PublicKey
func UnmarshalPublicKeyFromFile(file string) (*rsa.PublicKey, error) {
	block, err := readFileToBlock(file, BlockPublicKey)
	if err != nil {
		return nil, err
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("Unable to parse public key: %s", err)
	}

	return pub.(*rsa.PublicKey), nil
}

// UnmarshalPrivateKeyFromFile reads the given file and returns an RSA PrivateKey
func UnmarshalPrivateKeyFromFile(file string) (*rsa.PrivateKey, error) {
	block, err := readFileToBlock(file, BlockPrivateKey)
	if err != nil {
		return nil, err
	}

	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("Unable to parse private key: %s", err)
	}

	return priv, nil
}

// GenerateRandomString of given length with crypto/rand
func GenerateRandomString(length int) (string, error) {
	b := make([]byte, length)

	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", b), nil
}
