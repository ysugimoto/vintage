package function

import (
	"testing"
)

// Fastly built-in function testing implementation of uuid.version4
// Arguments may be:
// Reference: https://developer.fastly.com/reference/vcl/functions/uuid/uuid-version4/

func Test_Uuid_version4(t *testing.T) {
	t.Skip("uuid.version4 is randomized string, we trust uuid library")
}
