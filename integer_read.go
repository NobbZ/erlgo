package erlgo

import (
	"fmt"
	"math/big"
)

var twoFiveSix = big.NewInt(256)

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

	if len(binary)-3 < int(binary[1]) {
		return nil, nil, fmt.Errorf("%#v has less bytes then its size specifies", binary)
	}

	var bigRes *big.Int // := big.NewInt(0)
	var bigMul *big.Int // := big.NewInt(1)

	res := int64(0)
	mul := int64(1)

	for i := 0; i < int(binary[1]); i++ {
		if i < 7 {
			dig := int64(binary[3+i])
			dig *= mul
			res += dig
			mul *= 256
		} else if i == 7 {
			bigRes = big.NewInt(res)
			bigMul = big.NewInt(mul)
			bigDig := big.NewInt(int64(binary[3+i]))
			bigDig.Mul(bigDig, bigMul)
			bigRes.Add(bigRes, bigDig)
			bigMul.Mul(bigMul, twoFiveSix)
		} else {
			bigDig := big.NewInt(int64(binary[3+i]))
			bigDig.Mul(bigDig, bigMul)
			bigRes.Add(bigRes, bigDig)
			bigMul.Mul(bigMul, twoFiveSix)
		}
	}

	if binary[2] == 1 {
		if bigRes != nil {
			bigRes.Neg(bigRes)
		} else {
			res *= -1
		}
	}

	var rem []byte

	if len(binary) == 3+int(binary[1]) {
		rem = []byte{}
	} else {
		rem = binary[5:]
	}

	if bigRes != nil {
		return ErlBigInt{bigRes}, rem, nil
	} else {
		return ErlInt(res), rem, nil
	}
}
