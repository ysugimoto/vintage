package function

import (
	"testing"
)

// Fastly built-in function testing implementation of crypto.decrypt_base64
// Arguments may be:
// - ID, ID, ID, STRING, STRING, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/cryptographic/crypto-decrypt-base64/
func Test_Crypto_decrypt_base64(t *testing.T) {

	dec, err := Crypto_decrypt_base64(
		newTestRuntime(),
		"aes256",
		"cbc",
		"nopad",
		"603deb1015ca71be2b73aef0857d77811f352c073b6108d72d9810a30914dff4",
		"000102030405060708090a0b0c0d0e0f",
		"9YxMBNbl8bp3nqv7X3v71g==",
	)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if dec != "a8G+4i5An5bpPX4Rc5MXKg==" {
		t.Errorf("Decrypt value unmatch, expect=a8G+4i5An5bpPX4Rc5MXKg==, got=%s", dec)
	}
}
