package function

import (
	"testing"
)

// Fastly built-in function testing implementation of digest.awsv4_hmac
// Arguments may be:
// - STRING, STRING, STRING, STRING, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/cryptographic/digest-awsv4-hmac/
func Test_Digest_awsv4_hmac(t *testing.T) {

	// example from https://docs.aws.amazon.com/AmazonS3/latest/API/sig-v4-header-based-auth.html
	ret, err := Digest_awsv4_hmac(
		newTestRuntime(),
		"wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY",
		"20130524",
		"us-east-1",
		"s3",
		"AWS4-HMAC-SHA256\n20130524T000000Z\n20130524/us-east-1/s3/aws4_request\n7344ae5b7ee6c3e7e6b0fe0640412a37625d1fbfff95c48bbb2dc43964946972",
	)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	expect := "f0e8bdb87c964420e857bd35b5d6ed310bd44f0170aba48dd91039c6036bdb41"
	if ret != expect {
		t.Errorf("Return value unmatch, expect=%s, got=%s", expect, ret)
	}
}
