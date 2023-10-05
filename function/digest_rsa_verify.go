package function

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"

	"github.com/ysugimoto/vintage/errors"
	"github.com/ysugimoto/vintage/runtime/core"
)

const Digest_rsa_verify_Name = "digest.rsa_verify"

// Fastly built-in function implementation of digest.rsa_verify
// Arguments may be:
// - ID, STRING, STRING, STRING, ID
// - ID, STRING, STRING, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/cryptographic/digest-rsa-verify/
func Digest_rsa_verify[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	method string, // IDENT
	publicKey, payload, digest string,
	optional ...string, // IDENT, base64Method
) (bool, error) {

	base64Method := "url_nopad"
	if len(optional) > 0 {
		base64Method = optional[0]
	}

	hashMethod, err := digest_rsa_verify_HashMethod(method)
	if err != nil {
		return false, errors.WithStack(err)
	}

	payloadBytes := digest_rsa_verify_HashSum(payload, hashMethod)
	digestBytes, err := digest_rsa_verify_DecodeArgument(digest, base64Method)
	if err != nil {
		return false, errors.WithStack(err)
	}

	block, _ := pem.Decode([]byte(publicKey))
	if block == nil {
		return false, errors.FunctionError(
			Digest_rsa_verify_Name,
			"Failed to parse pem block of public key",
		)
	}
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return false, errors.FunctionError(
			Digest_rsa_verify_Name,
			"Failed to parse public key, %w", err,
		)
	}
	rsaKey, ok := pub.(*rsa.PublicKey)
	if !ok {
		return false, errors.FunctionError(
			Digest_rsa_verify_Name,
			"Provided public key does not seem to be RSA Public Key",
		)
	}
	if err := rsa.VerifyPKCS1v15(rsaKey, hashMethod, payloadBytes, digestBytes); err != nil {
		return false, nil
	}
	return true, nil
}

func digest_rsa_verify_HashMethod(method string) (crypto.Hash, error) {
	switch method {
	case "sha1":
		return crypto.SHA1, nil
	case "sha256":
		return crypto.SHA256, nil
	case "sha384":
		return crypto.SHA384, nil
	case "sha512":
		return crypto.SHA512, nil
	case "default":
		return crypto.SHA256, nil
	default:
		return crypto.Hash(0), errors.FunctionError(
			Digest_rsa_verify_Name,
			"Invalid hash_method %s provided on first argument",
			method,
		)
	}
}

func digest_rsa_verify_HashSum(payload string, hash crypto.Hash) []byte {
	switch hash {
	case crypto.SHA1:
		sum := sha1.Sum([]byte(payload))
		return sum[:]
	case crypto.SHA256:
		sum := sha256.Sum256([]byte(payload))
		return sum[:]
	case crypto.SHA384:
		sum := sha512.Sum384([]byte(payload))
		return sum[:]
	case crypto.SHA512:
		sum := sha512.Sum512([]byte(payload))
		return sum[:]
	default:
		return []byte(payload)
	}
}

func digest_rsa_verify_DecodeArgument(digest, b64 string) ([]byte, error) {
	switch b64 {
	case "standard":
		return base64.StdEncoding.DecodeString(digest)
	case "url":
		// Trick: url decoding may error. Then we try to decode as nopadding
		dec, err := base64.RawURLEncoding.DecodeString(digest)
		if err != nil {
			return base64.RawURLEncoding.DecodeString(digest)
		}
		return dec, nil
	case "url_nopad":
		// Trick: url decoding may error. Then we try to decode with padding
		dec, err := base64.RawURLEncoding.DecodeString(digest)
		if err != nil {
			return base64.URLEncoding.DecodeString(digest)
		}
		return dec, nil
	default:
		return nil, errors.FunctionError(
			Digest_rsa_verify_Name,
			"Invalid base64_method %s, 5th argument of digest.rsa_verify", b64,
		)
	}
}
