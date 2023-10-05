package function

import (
	"regexp"

	"github.com/ysugimoto/vintage/errors"
	"github.com/ysugimoto/vintage/runtime/core"
)

const Regsub_Name = "regsub"

func Regsub_convertGoExpandString(replacement string) (string, bool) {
	var converted []rune
	var found bool
	repl := []rune(replacement)

	for i := 0; i < len(repl); i++ {
		r := repl[i]
		if r != 0x5C { // escape sequence, "\"
			converted = append(converted, r)
			continue
		}
		// If rune is escape sequence, find next numeric character which indicates matched index like "\1"
		var matchIndex []rune
		for {
			if i+1 > len(repl)-1 {
				break
			}
			r = repl[i+1]
			if r >= 0x31 && r <= 0x39 {
				matchIndex = append(matchIndex, r)
				i++
				continue
			}
			break
		}
		if len(matchIndex) > 0 {
			converted = append(converted, []rune("${"+string(matchIndex)+"}")...)
			found = true
		}
	}

	return string(converted), found
}

// Fastly built-in function implementation of regsub
// Arguments may be:
// - STRING, STRING, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/strings/regsub/
func Regsub[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	input, pattern, replacement string,
) (string, error) {

	re, err := regexp.Compile(pattern)
	if err != nil {
		ctx.FastlyError = "EREGRECUR"
		return input, errors.FunctionError(
			Regsub_Name, "Invalid regular expression pattern: %s, error: %w", pattern, err,
		)
	}

	// Note: VCL's regsub uses PCRE regexp but golang is not PCRE
	matches := re.FindStringSubmatchIndex(input)
	if matches == nil {
		return input, nil
	}

	if expand, found := Regsub_convertGoExpandString(replacement); found {
		replaced := re.ExpandString([]byte{}, expand, input, matches)
		return string(replaced), nil
	}
	var replaced string
	if matches[0] > 0 {
		replaced += input[:matches[0]]
	}
	replaced += replacement
	if matches[1] < len(input)-1 {
		replaced += input[matches[1]:]
	}
	return replaced, nil
}
