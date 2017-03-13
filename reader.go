package erlgo

import (
	"bytes"
	"fmt"
	"io"
)

type Term interface {
	ToInteger() (Int, error)
	IsInteger() bool

	Matches(Term) bool
}

type ErlExtBinary struct {
	scanner io.ByteScanner
}

const (
	smallInteger    uint8 = 97
	integer               = 98
	smallBigInteger       = 110
	largeBigInteger       = 111
)

var funcMap = map[uint8]func(ErlExtBinary) (Term, error){
	smallInteger:    decodeSmallInteger,
	integer:         decodeInteger,
	smallBigInteger: decodeSmallBigInteger,
	largeBigInteger: decodeLargeBigInteger,
}

func NewReader(data []byte) ErlExtBinary {
	return ErlExtBinary{bytes.NewReader(data)}
}

func (b ErlExtBinary) Decode() (Term, error) {
	if version, err := b.scanner.ReadByte(); err != nil {
		return nil, err
	} else if version != 131 {
		return nil, fmt.Errorf("%v is an unknown version specifier", version)
	}

	return decodeRemaining(b)
}

func decodeRemaining(b ErlExtBinary) (Term, error) {
	if tag, err := b.scanner.ReadByte(); err != nil {
		return nil, err
	} else {
		b.scanner.UnreadByte()
		if f, ok := funcMap[tag]; ok {
			if res, err := f(b); err != nil {
				return nil, err
			} else {
				return res, nil
			}
		}
		return nil, fmt.Errorf("%v is an unknown tag", tag)
	}
}
