package variable

import "errors"

// ErrNotImplemented returns error which indicates variables is not implemented.
// Base variables interface implementation always returns this error
// so the any transformers MUST implement all of predefiend variables.
func ErrNotImplemented(name string) error {
	return errors.New("Variable " + name + " is not implemented")
}

// ErrNotFound indicates an error that vatiable is not found in our spec.
// Then we assume that we forget to implement that variable, please raise a report.
func ErrNotFound(name string) error {
	return errors.New("Variable " + name + " is not found")
}

// ErrCanntSet is similar to ErrNotFound but add more information that "variable cannot set")
func ErrCannotSet(name string) error {
	return errors.New("Variable " + name + " is not found or cannot set")
}

// ErrCanntUnset is similar to ErrNotFound but add more information that "variable cannot unset")
func ErrCannotUnset(name string) error {
	return errors.New("Variable " + name + " is not found or cannot unset")
}
