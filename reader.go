package erlgo

import (
	"bufio"
	"bytes"
	"compress/zlib"
	"fmt"
)

type Term interface {
	ToInteger() (Int, error)
	IsInteger() bool

	Matches(Term) bool
}

type ErlExtBinary struct {
	bs *bufio.Reader
}

const (
	newFloatExt        uint8 = 70
	bitBinaryExt             = 77
	compressed               = 80
	atomCacheRef             = 82
	smallIntegerExt          = 97
	integerExt               = 98
	floatExt                 = 99
	atomExt                  = 100
	reference                = 101
	portExt                  = 102
	pidExt                   = 103
	smallTupleExt            = 104
	largeTupleExt            = 105
	nilExt                   = 106
	stringExt                = 107
	listExt                  = 108
	binaryExt                = 109
	smallBigIntegerExt       = 110
	largeBigIntegerExt       = 111
	newFunExt                = 112
	exportExt                = 113
	newReferenceExt          = 114
	smallAtomExt             = 115
	mapExt                   = 116
	funExt                   = 117
	atomUtf8Ext              = 118
	smallAtomUtf8Ext         = 119
)

var funcMap = map[uint8]func(ErlExtBinary) (Term, error){
	newFloatExt:        decodeNewFloat,
	bitBinaryExt:       undefined, // TODO: as soon as there is no `undefined` left, remove that function
	atomCacheRef:       undefined,
	smallIntegerExt:    decodeSmallInteger,
	integerExt:         decodeInteger,
	floatExt:           decodeFloatExt,
	atomExt:            undefined,
	reference:          undefined,
	portExt:            undefined,
	pidExt:             undefined,
	smallTupleExt:      undefined,
	largeTupleExt:      undefined,
	nilExt:             undefined,
	stringExt:          undefined,
	listExt:            undefined,
	binaryExt:          undefined,
	smallBigIntegerExt: decodeSmallBigInteger,
	largeBigIntegerExt: decodeLargeBigInteger,
	newFunExt:          undefined,
	exportExt:          undefined,
	newReferenceExt:    undefined,
	smallAtomExt:       undefined,
	mapExt:             undefined,
	funExt:             undefined,
	atomUtf8Ext:        undefined,
	smallAtomUtf8Ext:   undefined,
}

// TODO: remove this function when there is no undefined left in the map above
func undefined(b ErlExtBinary) (Term, error) {
	tag, _ := b.bs.ReadByte()
	return nil, fmt.Errorf("Undefined parser for tag %v", tag)
}

func FromBytes(data []byte) ErlExtBinary {
	return ErlExtBinary{bufio.NewReader(bytes.NewReader(data))}
}

func (b ErlExtBinary) Decode() (Term, error) {
	if version, err := b.bs.ReadByte(); err != nil {
		return nil, err
	} else if version != 131 {
		return nil, fmt.Errorf("%v is an unknown version specifier", version)
	}

	return decodeRemaining(b)
}

func decodeRemaining(b ErlExtBinary) (Term, error) {
	if tag, err := b.bs.ReadByte(); err != nil {
		return nil, err
	} else {
		if tag == compressed {
			b.bs.ReadByte() // TODO: Instead of just skipping the uncompressed size, use it to verify data!
			b.bs.ReadByte()
			b.bs.ReadByte()
			b.bs.ReadByte()
			rc, err := zlib.NewReader(b.bs)

			b.bs = bufio.NewReader(rc)
			if tag, err = b.bs.ReadByte(); err != nil {
				return nil, err
			}
		}

		b.bs.UnreadByte() // TODO: as soon as undefined has been removed, we can get rid of this unreading
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
