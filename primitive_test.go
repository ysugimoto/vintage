package vintage

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestPrimitiveTypes(t *testing.T) {
	now := time.Now()
	tests := []struct {
		value        Value
		assertString string
		assertBool   bool
	}{
		{value: String("foo"), assertString: "foo", assertBool: true},
		{value: String(""), assertString: "", assertBool: false},
		{value: Integer(0), assertString: "0", assertBool: false},
		{value: Integer(100), assertString: "100", assertBool: false},
		{value: Float(0), assertString: "0.000", assertBool: false},
		{value: Float(10.0), assertString: "10.000", assertBool: false},
		{value: Bool(true), assertString: "1", assertBool: true},
		{value: Bool(false), assertString: "0", assertBool: false},
		{value: RTime(time.Second), assertString: "1.000", assertBool: false},
		{value: Time(now), assertString: now.Format("Mon, 02 Jan 2006 15:04:05 GMT"), assertBool: false},
	}

	for i, tt := range tests {
		if diff := cmp.Diff(tt.value.String(), tt.assertString); diff != "" {
			t.Errorf("[%d] Value unmatch, diff=%s", i, diff)
		}
		if diff := cmp.Diff(tt.value.Bool(), tt.assertBool); diff != "" {
			t.Errorf("[%d] Value unmatch, diff=%s", i, diff)
		}
	}
}

func TestAssignPrimitive(t *testing.T) {
	var str String
	str = "foo"
	if str.String() != "foo" {
		t.Errorf("unmatch value, expect=foo, got=%s", str.String())
	}
}
