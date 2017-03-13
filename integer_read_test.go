package erlgo_test

import (
	"github.com/NobbZ/erlgo"
	"math/big"
	"testing"
)

var plusNine, _ = big.NewInt(0).SetString("18446744073709551616", 10)
var plusTen, _ = big.NewInt(0).SetString("4722366482869645213696", 10)
var plusVeryBig, _ = big.NewInt(0).SetString(
	"4931183787736664932360058088481132806464249064592816777363639133838600942820417921935608125537553934278674005267623599165972"+
		"8331223283265831128162210767033570298579967195123431015316391585772868035976621069439038508288907840911493166867209378778336"+
		"2893396695740300064741326536430985501229973638902647863548613194784388249853831252667031319724958132568898411896638150110768"+
		"6008635362008714927712797983425463367606140704111001183715568718307746262268630617253614384647693738511782868915581833149250"+
		"99540247780495920664946518646198552749613009880449926596639031121858756000207590413184793166384097191709192063287296", 10)
var plusVeryBigMulTwoFiveSix = big.NewInt(0).Mul(big.NewInt(256), plusVeryBig)
var minusNine = big.NewInt(0).Neg(plusNine)
var minusTen = big.NewInt(0).Neg(plusTen)
var minusVeryBig = big.NewInt(0).Neg(plusVeryBig)
var minusVeryBigMulTwoFiveSix = big.NewInt(0).Neg(plusVeryBigMulTwoFiveSix)

var table = []struct {
	Name   string
	Data   erlgo.ErlExtBinary
	Expect erlgo.Term
}{
	{"0", erlgo.ErlExtBinary{131, 97, 0}, erlgo.Int64(0)},
	{"1", erlgo.ErlExtBinary{131, 97, 1}, erlgo.Int64(1)},
	{"2", erlgo.ErlExtBinary{131, 97, 2}, erlgo.Int64(2)},
	{"10", erlgo.ErlExtBinary{131, 97, 10}, erlgo.Int64(10)},
	{"100", erlgo.ErlExtBinary{131, 97, 100}, erlgo.Int64(100)},
	{"256", erlgo.ErlExtBinary{131, 98, 0, 0, 1, 0}, erlgo.Int64(256)},
	{"65536", erlgo.ErlExtBinary{131, 98, 0, 1, 0, 0}, erlgo.Int64(65536)},
	{"16777216", erlgo.ErlExtBinary{131, 98, 1, 0, 0, 0}, erlgo.Int64(16777216)},
	{"4294967296", erlgo.ErlExtBinary{131, 110, 5, 0, 0, 0, 0, 0, 1}, erlgo.Int64(4294967296)},
	{"1099511627776", erlgo.ErlExtBinary{131, 110, 6, 0, 0, 0, 0, 0, 0, 1}, erlgo.Int64(1099511627776)},
	{"281474976710656", erlgo.ErlExtBinary{131, 110, 7, 0, 0, 0, 0, 0, 0, 0, 1}, erlgo.Int64(281474976710656)},
	{"72057594037927936", erlgo.ErlExtBinary{131, 110, 8, 0, 0, 0, 0, 0, 0, 0, 0, 1}, erlgo.Int64(72057594037927936)},
	{"18446744073709551616", erlgo.ErlExtBinary{131, 110, 9, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, erlgo.IntBig{plusNine}},
	{"4722366482869645213696", erlgo.ErlExtBinary{131, 110, 10, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, erlgo.IntBig{plusTen}},
	{"veryBig", erlgo.ErlExtBinary{131, 110, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, erlgo.IntBig{plusVeryBig}},
	{"veryBig * 256", erlgo.ErlExtBinary{131, 111, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, erlgo.IntBig{plusVeryBigMulTwoFiveSix}},
	{"-1", erlgo.ErlExtBinary{131, 98, 255, 255, 255, 255}, erlgo.Int64(-1)},
	{"-2", erlgo.ErlExtBinary{131, 98, 255, 255, 255, 254}, erlgo.Int64(-2)},
	{"-10", erlgo.ErlExtBinary{131, 98, 255, 255, 255, 246}, erlgo.Int64(-10)},
	{"-100", erlgo.ErlExtBinary{131, 98, 255, 255, 255, 156}, erlgo.Int64(-100)},
	{"-256", erlgo.ErlExtBinary{131, 98, 255, 255, 255, 0}, erlgo.Int64(-256)},
	{"-65536", erlgo.ErlExtBinary{131, 98, 255, 255, 0, 0}, erlgo.Int64(-65536)},
	{"-16777216", erlgo.ErlExtBinary{131, 98, 255, 0, 0, 0}, erlgo.Int64(-16777216)},
	{"-4294967296", erlgo.ErlExtBinary{131, 110, 5, 1, 0, 0, 0, 0, 1}, erlgo.Int64(-4294967296)},
	{"-1099511627776", erlgo.ErlExtBinary{131, 110, 6, 1, 0, 0, 0, 0, 0, 1}, erlgo.Int64(-1099511627776)},
	{"-281474976710656", erlgo.ErlExtBinary{131, 110, 7, 1, 0, 0, 0, 0, 0, 0, 1}, erlgo.Int64(-281474976710656)},
	{"-72057594037927936", erlgo.ErlExtBinary{131, 110, 8, 1, 0, 0, 0, 0, 0, 0, 0, 1}, erlgo.Int64(-72057594037927936)},
	{"-18446744073709551616", erlgo.ErlExtBinary{131, 110, 9, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1}, erlgo.IntBig{minusNine}},
	{"-4722366482869645213696", erlgo.ErlExtBinary{131, 110, 10, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, erlgo.IntBig{minusTen}},
	{"veryBig (negative)", erlgo.ErlExtBinary{131, 110, 255, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, erlgo.IntBig{minusVeryBig}},
	{"veryBig * 256 (negative)", erlgo.ErlExtBinary{131, 111, 0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, erlgo.IntBig{minusVeryBigMulTwoFiveSix}},
}

func TestReadingIntegers(t *testing.T) {
	for _, test := range table {
		t.Run(test.Name, func(t *testing.T) {
			if val, err := test.Data.Decode(); err == nil && !val.Matches(test.Expect) {
				t.Errorf(`%#v parsed into %#v, expected %#v."`, test.Data, val, test.Expect)
			} else if err != nil {
				t.Errorf(`%#v encountered error "%v", expected value %#v.`, test.Data, err, test.Expect)
			}
		})
	}
}

func BenchmarkReadingIntegers(b *testing.B) {
	for _, data := range table {
		b.Run(data.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				data.Data.Decode()
			}
		})
	}
}
