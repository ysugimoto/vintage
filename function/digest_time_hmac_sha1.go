package function

import (
	"crypto/sha1"
	"encoding/base32"
	"encoding/base64"
	"time"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"github.com/ysugimoto/vintage/errors"
	"github.com/ysugimoto/vintage/runtime/core"
)

const Digest_time_hmac_sha1_Name = "digest.time_hmac_sha1"

// Fastly built-in function implementation of digest.time_hmac_sha1
// Arguments may be:
// - STRING, INTEGER, INTEGER
// Reference: https://developer.fastly.com/reference/vcl/functions/cryptographic/digest-time-hmac-sha1/
func Digest_time_hmac_sha1[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	secret string,
	interval, offset int64,
) (string, error) {

	return digest_time_hmac_sha1(time.Now(), secret, interval, offset)
}

func digest_time_hmac_sha1(
	baseTime time.Time,
	secret string,
	interval, offset int64,
) (string, error) {

	dec, err := base64.StdEncoding.DecodeString(secret)
	if err != nil {
		return "", errors.FunctionError(
			Digest_time_hmac_sha1_Name,
			"Failed to base64 decode secret string: %w", err,
		)
	}

	var skew uint
	if offset >= 0 {
		skew = uint(offset)
	}

	key := base32.StdEncoding.EncodeToString(dec)
	pass, err := totp.GenerateCodeCustom(key, baseTime, totp.ValidateOpts{
		Period:    uint(interval),
		Digits:    otp.DigitsSix,
		Algorithm: otp.AlgorithmSHA1,
		Skew:      skew,
	})
	if err != nil {
		return "", errors.FunctionError(
			Digest_time_hmac_sha1_Name,
			"Failed to generate TOTP password: %w", err,
		)
	}

	enc := sha1.Sum([]byte(pass))
	return base64.StdEncoding.EncodeToString(enc[:]), nil
}
