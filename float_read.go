package erlgo

import (
	"encoding/binary"
	"errors"
	"math"
)

type Float float64

func (f Float) IsInteger() bool {
	return false
}

func (f Float) ToInteger() (Int, error) {
	return nil, errors.New("Not an Integer")
}

func (f Float) Matches(other Term) bool {
	switch o := other.(type) {
	case Float:
		return float64(f) == float64(o)
	default:
		return false
	}
}

func decodeNewFloat(b ErlExtBinary) (Term, error) {
	_, _ = b.bs.ReadByte() // skip tag; TODO: remove this when ready!

	var resultBytes = make([]byte, 8)

	for i := 0; i < 8; i++ {
		if byte, err := b.bs.ReadByte(); err != nil {
			return Float(0.0), err
		} else {
			// result |= uint64(byte) << uint(i*8)
			resultBytes[i] = byte
		}
	}

	resultBits := binary.BigEndian.Uint64(resultBytes)
	result := math.Float64frombits(resultBits)

	return Float(result), nil
}
