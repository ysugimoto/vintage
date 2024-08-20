package function

import (
	"math/rand"
	"time"

	"github.com/ysugimoto/vintage/runtime/core"
)

const Randomstr_Name = "randomstr"

var Randomstr_default_characters = []rune(
	"abcdedfhijklmnopqrstuvwxyzABCDEDFHIJKLMNOPQRSTUVWXYZ0123456789-_",
)

// Fastly built-in function implementation of randomstr
// Arguments may be:
// - INTEGER
// - INTEGER, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/randomness/randomstr/
func Randomstr[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	length int64,
	optional ...string,
) (string, error) {

	characters := Randomstr_default_characters
	if len(optional) > 0 {
		characters = []rune(optional[0])
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	ret := make([]rune, int(length))

	for i := 0; i < int(length); i++ {
		ret[i] = characters[r.Intn(len(characters)-1)]
	}

	return string(ret), nil
}
