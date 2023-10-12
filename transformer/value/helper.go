package value

import "fmt"

var temporaryVarCount int
var useFixedName bool

func UseFixedTemporalValue() {
	useFixedName = true
}

func Temporary() string {
	if useFixedName {
		return "v__fixed"
	}
	temporaryVarCount++
	return fmt.Sprintf("v__%d", temporaryVarCount)
}

var ErrorCheck = "if err != nil {\nreturn vintage.NONE, err\n}"
