package perc

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"testing"
)

// Some convenience functions copied from github.com/bifurcation/mint
func assert(t *testing.T, test bool, msg string) {
	if !test {
		t.Fatalf(msg)
	}
}

func assertByteEquals(t *testing.T, a []byte, b []byte) {
	if !bytes.Equal(a, b) {
		assert(t, false, fmt.Sprintf("%v != %v", hex.EncodeToString(a), hex.EncodeToString(b)))
	}
}

var (
// TODO pre-computed test constants
)

func TestDoubleEncrypt(t *testing.T) {
	// TODO Test that encrypt(known plaintext) == known ciphertext
}

func TestDoubleDecrypt(t *testing.T) {
	// TODO Test that encrypt(known ciphertext) == known plaintext
}

func TestDoubleRoundTrip(t *testing.T) {
	// TODO Test that encrypt(decrypt(ciphertext)) == ciphertext
	// TODO Test that decrypt(encrypt(plaintext)) == plaintext
}
