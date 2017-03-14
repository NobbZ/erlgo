package erlgo

import (
	"encoding/binary"
	"errors"
)

type Nil struct{}

func (n Nil) IsInteger() bool {
	return false
}

func (n Nil) ToInteger() (Int, error) {
	return nil, errors.New("Not an Integer")
}

func (n Nil) Matches(other Term) bool {
	switch other.(type) {
	case Nil:
		return true
	default:
		return false
	}
}

type List struct {
	firsts []Term
	tail   Term
}

func (l List) IsInteger() bool {
	return false
}

func (l List) ToInteger() (Int, error) {
	return nil, errors.New("Not an Integer")
}

func (l List) Matches(other Term) bool {
	switch o := other.(type) {
	case List:
		if len(l.firsts) != len(o.firsts) {
			return false
		}

		for idx := range l.firsts {
			if l.firsts[idx].Matches(o.firsts[idx]) {
				return false
			}
		}
		return l.tail.Matches(o.tail)
	default:
		return false
	}
}

func NewListFromTerms(terms []Term) List {
	return List{
		firsts: terms,
		tail:   Nil{},
	}
}

func decodeStringExt(b ErlExtBinary) (Term, error) {
	_, err := b.bs.ReadByte()

	lengthBytes := []byte{0, 0}
	if lengthBytes[0], err = b.bs.ReadByte(); err != nil {
		return nil, err
	} else if lengthBytes[1], err = b.bs.ReadByte(); err != nil {
		return nil, err
	}

	length := binary.BigEndian.Uint16(lengthBytes)
	result := make([]Term, length)

	for i := uint16(0); i < length; i++ {
		if byte, err := b.bs.ReadByte(); err != nil {
			return nil, err
		} else {
			result[i] = Int64(byte)
		}
	}

	return List{result, Nil{}}, nil
}
