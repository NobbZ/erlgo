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
	case IntBig:
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

type IntBig struct {
	*big.Int
}

func (ei IntBig) ToInteger() (Int, error) {
	return ei, nil
}

func (ei IntBig) IsInteger() bool {
	return true
}

func (ei IntBig) Matches(other Term) bool {
	switch o := other.(type) {
	case Int64:
		return o.Matches(ei)
	case IntBig:
		return ei.Int.Cmp(o.Int) == 0
	default:
		return false
	}
}

func (ebi IntBig) Int64() (int64, bool) {
	if ebi.Cmp(int64MaxBig) == 1 || ebi.Cmp(int64MinBig) == -1 {
		return 0, false
	}
	return ebi.Int.Int64(), true
}

func (ebi IntBig) BigInt() *big.Int {
	return ebi.Int
}

func readInt32(b ErlExtBinary) (int32, error) {
	result := int32(0)

	if b1, err := b.bs.ReadByte(); err != nil {
		return 0, err
	} else if b2, err := b.bs.ReadByte(); err != nil {
		return 0, err
	} else if b3, err := b.bs.ReadByte(); err != nil {
		return 0, err
	} else if b4, err := b.bs.ReadByte(); err != nil {
		return 0, err
	} else {
		result = int32(b1)<<24 | int32(b2)<<16 | int32(b3)<<8 | int32(b4)
	}

	return result, nil
}

func decodeSmallInteger(b ErlExtBinary) (Term, error) {
	if tag, err := b.bs.ReadByte(); err != nil {
		return nil, err
	} else if tag != smallIntegerExt {
		return nil, fmt.Errorf("%v is not tagging a small integer", tag)
	}

	if byte, err := b.bs.ReadByte(); err != nil {
		return nil, err
	} else {
		return Int64(byte), nil
	}
}

func decodeInteger(b ErlExtBinary) (Term, error) {
	if tag, err := b.bs.ReadByte(); err != nil {
		return nil, err
	} else if tag != integerExt {
		return nil, fmt.Errorf("%v is not tagging a integer", tag)
	}

	res, err := readInt32(b)

	return Int64(res), err
}

func decodeSmallBigInteger(b ErlExtBinary) (Term, error) {
	if tag, err := b.bs.ReadByte(); err != nil {
		return nil, err
	} else if tag != smallBigIntegerExt {
		return nil, fmt.Errorf("%v is not tagging a small big integer", tag)
	}

	if byteCount, err := b.bs.ReadByte(); err != nil {
		return nil, err
	} else {
		res := int64(0)
		mul := int64(1)

		signum, err := b.bs.ReadByte()
		if err != nil {
			return nil, err
		}

		var bigRes *big.Int // := big.NewInt(0)
		var bigMul *big.Int // := big.NewInt(1)

		for i := 0; i < int(byteCount); i++ {
			if i < 7 {
				if dig, err := b.bs.ReadByte(); err != nil {
					return nil, err
				} else {
					dig := int64(dig)
					dig *= mul
					res += dig
					mul *= 256
				}
			} else if i == 7 {
				if dig, err := b.bs.ReadByte(); err != nil {
					return nil, err
				} else {
					bigRes = big.NewInt(res)
					bigMul = big.NewInt(mul)
					dig := big.NewInt(int64(dig))
					dig.Mul(dig, bigMul)
					bigRes.Add(bigRes, dig)
					bigMul.Mul(bigMul, twoFiveSix)
				}
			} else {
				if dig, err := b.bs.ReadByte(); err != nil {
					return nil, err
				} else {
					dig := big.NewInt(int64(dig))
					dig.Mul(dig, bigMul)
					bigRes.Add(bigRes, dig)
					bigMul.Mul(bigMul, twoFiveSix)
				}
			}
		}

		if signum == 1 {
			if bigRes != nil {
				bigRes.Neg(bigRes)
			} else {
				res *= -1
			}
		}

		if bigRes != nil {
			return IntBig{bigRes}, nil
		} else {
			return Int64(res), nil
		}
	}
}

func decodeLargeBigInteger(b ErlExtBinary) (Term, error) {
	if tag, err := b.bs.ReadByte(); err != nil {
		return nil, err
	} else if tag != largeBigIntegerExt {
		return nil, fmt.Errorf("%v is not tagging a large big integer", tag)
	}

	byteCount, err := readInt32(b)
	if err != nil {
		return nil, err
	}

	signum, err := b.bs.ReadByte()
	if err != nil {
		return nil, err
	}

	var bigRes *big.Int // := big.NewInt(0)
	var bigMul *big.Int // := big.NewInt(1)

	res := int64(0)
	mul := int64(1)

	for i := uint32(0); i < uint32(byteCount); i++ {
		if i < 7 {
			if dig, err := b.bs.ReadByte(); err != nil {
				return nil, err
			} else {
				dig := int64(dig)
				dig *= mul
				res += dig
				mul *= 256
			}
		} else if i == 7 {
			if dig, err := b.bs.ReadByte(); err != nil {
				return nil, err
			} else {
				bigRes = big.NewInt(res)
				bigMul = big.NewInt(mul)
				dig := big.NewInt(int64(dig))
				dig.Mul(dig, bigMul)
				bigRes.Add(bigRes, dig)
				bigMul.Mul(bigMul, twoFiveSix)
			}
		} else {
			if dig, err := b.bs.ReadByte(); err != nil {
				return nil, err
			} else {
				dig := big.NewInt(int64(dig))
				dig.Mul(dig, bigMul)
				bigRes.Add(bigRes, dig)
				bigMul.Mul(bigMul, twoFiveSix)
			}
		}
	}

	if signum == 1 {
		if bigRes != nil {
			bigRes.Neg(bigRes)
		} else {
			res *= -1
		}
	}

	if bigRes != nil {
		return IntBig{bigRes}, nil
	} else {
		return Int64(res), nil
	}
}
