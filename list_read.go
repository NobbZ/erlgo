package erlgo

import (
	"encoding/binary"
	"errors"
)

type List interface {
	ToSlice() ([]Term, err)
	Len() int
}

type Nil struct{}

func (n Nil) ToSlice() ([]Term, err) {
	return []Term{}
}

func (n Nil) Len() int {
	return 0
}

func (n Nil) IsInteger() bool {
	return false
}

func (n Nil) IsList() bool { return true }

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

type Cons struct {
	this Term
	next Term
}

func (c Cons) ToSlice() ([]Term, err) {
	return []Term{}
}

func (c Cons) Len() int {
	return 0
}

func (c Cons) IsInteger() bool {
	return false
}

func (c Cons) IsList() bool { return true }

func (c Cons) ToInteger() (Int, error) {
	return nil, errors.New("Not an Integer")
}

func (c Cons) Matches(other Term) bool {
	switch o := other.(type) {
	case Cons:
		x, y := c, o

		for x.IsList() && y.IsList() {

		}
	default:
		return false
	}
}

func NewListFromTerms(terms []Term) List {
	var result Cons = Nil{}

	for i := len(terms) - 1; i >= 0; i-- {
		result = Cons{
			this: terms[i],
			next: result,
		}
	}

	return result
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

	return Cons{result, Nil{}}, nil
}
