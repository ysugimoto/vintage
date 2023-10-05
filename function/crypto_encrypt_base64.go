package function

import (
	"encoding/base64"

	"github.com/ysugimoto/vintage/errors"
	"github.com/ysugimoto/vintage/lib"
	"github.com/ysugimoto/vintage/runtime/core"
)

const Crypto_encrypt_base64_Name = "crypto.encrypt_base64"

// Fastly built-in function implementation of crypto.encrypt_base64
// Arguments may be:
// - ID, ID, ID, STRING, STRING, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/cryptographic/crypto-encrypt-base64/
func Crypto_encrypt_base64[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	cipherId, mode, padding string, // IDENT
	key, iv, text string,
) (string, error) {

	buf, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		return "", errors.FunctionError(
			Crypto_encrypt_base64_Name,
			"Failed to decode base64 encoded string: %s", text,
		)
	}

	codec, err := lib.NewCryptoCodec(
		Crypto_encrypt_base64_Name,
		cipherId, mode, padding,
	)
	if err != nil {
		return "", errors.WithStack(err)
	}

	encrypted, err := codec.Encrypt(key, iv, buf)
	if err != nil {
		return "", errors.WithStack(err)
	}

	return base64.StdEncoding.EncodeToString(encrypted), nil
}
