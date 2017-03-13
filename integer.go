package erlgo

import "math/big"

type Int interface {
	Int64() (int64, bool)
	BigInt() *big.Int

	//Abs(ErlInteger) ErlInteger
	//Add(ErlInteger) ErlInteger
}
