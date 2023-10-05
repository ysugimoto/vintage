package function

import (
	"encoding/hex"

	"github.com/ysugimoto/vintage/errors"
	"github.com/ysugimoto/vintage/lib"
	"github.com/ysugimoto/vintage/runtime/core"
)

const Crypto_decrypt_hex_Name = "crypto.decrypt_hex"

// Fastly built-in function implementation of crypto.decrypt_hex
// Arguments may be:
// - ID, ID, ID, STRING, STRING, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/cryptographic/crypto-decrypt-hex/
func Crypto_decrypt_hex[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	cipherId, mode, padding string, // IDENT
	key, iv, text string,
) (string, error) {

	buf, err := hex.DecodeString(text)
	if err != nil {
		return "", errors.FunctionError(
			Crypto_decrypt_hex_Name,
			"Failed to decode hex string: %s", text,
		)
	}

	codec, err := lib.NewCryptoCodec(
		Crypto_decrypt_hex_Name,
		cipherId, mode, padding,
	)
	if err != nil {
		return "", errors.WithStack(err)
	}

	decrypted, err := codec.Decrypt(key, iv, buf)
	if err != nil {
		return "", errors.WithStack(err)
	}

	return hex.EncodeToString(decrypted), nil
}
