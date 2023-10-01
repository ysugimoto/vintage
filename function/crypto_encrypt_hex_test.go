package function

import (
	"testing"
)

// Fastly built-in function testing implementation of crypto.encrypt_hex
// Arguments may be:
// - ID, ID, ID, STRING, STRING, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/cryptographic/crypto-encrypt-hex/
func Test_Crypto_encrypt_hex(t *testing.T) {

	enc, err := Crypto_encrypt_hex(
		newTestRuntime(),
		"aes256",
		"cbc",
		"nopad",
		"603deb1015ca71be2b73aef0857d77811f352c073b6108d72d9810a30914dff4",
		"000102030405060708090a0b0c0d0e0f",
		"6bc1bee22e409f96e93d7e117393172a",
	)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if enc != "f58c4c04d6e5f1ba779eabfb5f7bfbd6" {
		t.Errorf("Encrypt value unmatch, expect=6bc1bee22e409f96e93d7e117393172a, got=%s", enc)
	}
}
