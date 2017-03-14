package erlgo

import "errors"

type Float float64

func (f Float) IsInteger() bool {
	return false
}

func (f Float) ToInteger() (Int, error) {
	return nil, errors.New("Not an Integer")
}

func (f Float) Matches(o Term) bool {
	return false
}
