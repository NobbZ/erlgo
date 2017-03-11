package erl_ext

import (
	"fmt"
	"math/big"
)

type ErlInt int64

func (ei ErlInt) ToInteger() (int64, error) {
	return int64(ei), nil
}

func (ei ErlInt) IsInteger() bool {
	return true
}

func (ei ErlInt) Matches(other ErlType) bool {
	switch o := other.(type) {
	case ErlInt:
		return int64(ei) == int64(o)
	case ErlBigInt:
		return big.NewInt(int64(ei)).Cmp(o.Int) == 0
	default:
		return false
	}
}

type ErlBigInt struct {
	*big.Int
}

func (ei ErlBigInt) ToInteger() (int64, error) {
	return 0, nil
}

func (ei ErlBigInt) IsInteger() bool {
	return true
}

func (ei ErlBigInt) Matches(other ErlType) bool {
	switch o := other.(type) {
	case ErlInt:
		return o.Matches(ei)
	case ErlBigInt:
		return ei.Int.Cmp(o.Int) == 0
	default:
		return false
	}
}

func decodeSmallInteger(binary ErlExtBinary) (ErlType, []byte, error) {
	if binary[0] != smallInteger {
		return nil, nil, fmt.Errorf("%v is not tagging a small integer", binary[0])
	}

	if len(binary) < 2 {
		return nil, nil, fmt.Errorf("%#v has not enough bytes for a small integer", binary)
	}

	var rem []byte

	if len(binary) == 2 {
		rem = []byte{}
	} else {
		rem = binary[2:]
	}

	return ErlInt(binary[1]), rem, nil
}

func decodeInteger(binary ErlExtBinary) (ErlType, []byte, error) {
	if binary[0] != integer {
		return nil, nil, fmt.Errorf("%v is not tagging a integer", binary[0])
	}

	if len(binary) < 5 {
		return nil, nil, fmt.Errorf("%#v has not enough bytes for a integer", binary)
	}

	var rem []byte

	if len(binary) == 5 {
		rem = []byte{}
	} else {
		rem = binary[5:]
	}

	res := ErlInt(int32(binary[1])<<24 | int32(binary[2])<<16 | int32(binary[3])<<8 | int32(binary[4]))

	return res, rem, nil
}

func decodeSmallBigInteger(binary ErlExtBinary) (ErlType, []byte, error) {
	if binary[0] != smallBigInteger {
		return nil, nil, fmt.Errorf("%v is not tagging a small big integer", binary[0])
	}

	if len(binary) < 4 {
		return nil, nil, fmt.Errorf("%#v has not enough bytes for a small big integer", binary)
	}

	if len(binary) < (3 + int(binary[1])) {
		return nil, nil, fmt.Errorf("%#v has less bytes then its size specifies", binary)
	}

	res := big.NewInt(0)
	mul := big.NewInt(1)
	twoFiveSix := big.NewInt(256)

	for i := 3; i < int(3+binary[1]); i++ {
		dig := big.NewInt(int64(binary[i]))
		dig.Mul(dig, mul)
		res.Add(res, dig)
		if i < int(2+binary[1]) {
			mul.Mul(mul, twoFiveSix)
		}
	}

	if binary[2] == 1 {
		res.Neg(res)
	}

	var rem []byte

	if len(binary) == int(3+binary[1]) {
		rem = []byte{}
	} else {
		rem = binary[5:]
	}

	return ErlBigInt{res}, rem, nil
}
