package function

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"strings"

	"github.com/ysugimoto/vintage/runtime/core"
)

const Digest_awsv4_hmac_Name = "digest.awsv4_hmac"

// Fastly built-in function implementation of digest.awsv4_hmac
// Arguments may be:
// - STRING, STRING, STRING, STRING, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/cryptographic/digest-awsv4-hmac/
func Digest_awsv4_hmac[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	key, dateStamp, region, service, stringToSign string,
) (string, error) {

	signature := []byte("AWS4" + key)
	hashes := []string{
		dateStamp,
		region,
		service,
		"aws4_request",
		stringToSign,
	}
	for i := range hashes {
		mac := hmac.New(sha256.New, signature)
		mac.Write([]byte(hashes[i]))
		signature = mac.Sum(nil)
	}

	return strings.ToLower(hex.EncodeToString(signature)), nil
}
