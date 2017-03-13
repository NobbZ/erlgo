package erlgo

import (
	"fmt"
	"math"
	"math/big"
)

var int64MaxBig = big.NewInt(math.MaxInt64)
var int64MinBig = big.NewInt(math.MinInt64)

var twoFiveSix = big.NewInt(256)

type Int64 int64

func (ei Int64) ToInteger() (Int, error) {
	return ei, nil
}

func (ei Int64) IsInteger() bool {
	return true
}

func (ei Int64) Matches(other Term) bool {
	switch o := other.(type) {
	case Int64:
		return int64(ei) == int64(o)
	case ErlBigInt:
		return big.NewInt(int64(ei)).Cmp(o.Int) == 0
	default:
		return false
	}
}

func (ei Int64) Int64() (int64, bool) {
	return int64(ei), true
}

func (ei Int64) BigInt() *big.Int {
	return big.NewInt(int64(ei))
}

type ErlBigInt struct {
	*big.Int
}

func (ei ErlBigInt) ToInteger() (Int, error) {
	return ei, nil
}

func (ei ErlBigInt) IsInteger() bool {
	return true
}

func (ei ErlBigInt) Matches(other Term) bool {
	switch o := other.(type) {
	case Int64:
		return o.Matches(ei)
	case ErlBigInt:
		return ei.Int.Cmp(o.Int) == 0
	default:
		return false
	}
}

func (ebi ErlBigInt) Int64() (int64, bool) {
	if ebi.Cmp(int64MaxBig) == 1 || ebi.Cmp(int64MinBig) == -1 {
		return 0, false
	}
	return ebi.Int.Int64(), true
}

func (ebi ErlBigInt) BigInt() *big.Int {
	return ebi.Int
}

func decodeSmallInteger(binary ErlExtBinary) (Term, []byte, error) {
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

	return Int64(binary[1]), rem, nil
}

func decodeInteger(binary ErlExtBinary) (Term, []byte, error) {
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

	res := Int64(int32(binary[1])<<24 | int32(binary[2])<<16 | int32(binary[3])<<8 | int32(binary[4]))

	return res, rem, nil
}

func decodeSmallBigInteger(binary ErlExtBinary) (Term, []byte, error) {
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
		return Int64(res), rem, nil
	}
}

func decodeLargeBigInteger(binary ErlExtBinary) (Term, []byte, error) {
	if binary[0] != largeBigInteger {
		return nil, nil, fmt.Errorf("%v is not tagging a small big integer", binary[0])
	}

	if len(binary) < 7 {
		return nil, nil, fmt.Errorf("%#v has not enough bytes for a small big integer", binary)
	}

	if len(binary)-6 < int(binary[1]) {
		return nil, nil, fmt.Errorf("%#v has less bytes then its size specifies", binary)
	}

	numBytes := uint32(binary[1])<<24 | uint32(binary[2])<<16 | uint32(binary[3])<<8 | uint32(binary[4])

	var bigRes *big.Int // := big.NewInt(0)
	var bigMul *big.Int // := big.NewInt(1)

	res := int64(0)
	mul := int64(1)

	for i := uint32(0); i < numBytes; i++ {
		if i < 7 {
			dig := int64(binary[6+i])
			dig *= mul
			res += dig
			mul *= 256
		} else if i == 7 {
			bigRes = big.NewInt(res)
			bigMul = big.NewInt(mul)
			bigDig := big.NewInt(int64(binary[6+i]))
			bigDig.Mul(bigDig, bigMul)
			bigRes.Add(bigRes, bigDig)
			bigMul.Mul(bigMul, twoFiveSix)
		} else {
			bigDig := big.NewInt(int64(binary[6+i]))
			bigDig.Mul(bigDig, bigMul)
			bigRes.Add(bigRes, bigDig)
			bigMul.Mul(bigMul, twoFiveSix)
		}
	}

	if binary[5] == 1 {
		if bigRes != nil {
			bigRes.Neg(bigRes)
		} else {
			res *= -1
		}
	}

	var rem []byte

	if int64(len(binary)) == 6+int64(numBytes) {
		rem = []byte{}
	} else {
		rem = binary[5:]
	}

	if bigRes != nil {
		return ErlBigInt{bigRes}, rem, nil
	} else {
		return Int64(res), rem, nil
	}
}
