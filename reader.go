package erlgo

import (
	"fmt"
)

type ErlType interface {
	ToInteger() (int64, error)
	IsInteger() bool

	Matches(ErlType) bool
}

type ErlExtBinary []byte

const (
	smallInteger    uint8 = 97
	integer               = 98
	smallBigInteger       = 110
)

var funcMap = map[uint8]func(ErlExtBinary) (ErlType, []byte, error){
	smallInteger:    decodeSmallInteger,
	integer:         decodeInteger,
	smallBigInteger: decodeSmallBigInteger,
}

func (b ErlExtBinary) Decode() (ErlType, error) {
	if len(b) < 2 {
		return nil, fmt.Errorf("%v is to short", b)
	}

	if b[0] != 131 {
		return nil, fmt.Errorf("%v is an unknown version specifier", b[0])
	}

	return decodeRemaining(b[1:])
}

func decodeRemaining(b ErlExtBinary) (ErlType, error) {
	if f, ok := funcMap[b[0]]; ok {
		res, rem, err := f(b)

		if err != nil {
			return nil, err
		}

		if len(rem) != 0 {
			return nil, fmt.Errorf("There were bytes left to consume after resolving nesting")
		}

		return res, nil
	}
	return nil, fmt.Errorf("%v is an unknown tag", b[0])
}
