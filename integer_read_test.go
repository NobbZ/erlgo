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
var minusNine = big.NewInt(0).Neg(plusNine)
var minusTen = big.NewInt(0).Neg(plusTen)
var minusVeryBig = big.NewInt(0).Neg(plusVeryBig)

var table = []struct {
	Name   string
	Data   erl_ext.ErlExtBinary
	Expect erl_ext.ErlType
}{
	{"0", erl_ext.ErlExtBinary{131, 97, 0}, erl_ext.ErlInt(0)},
	{"1", erl_ext.ErlExtBinary{131, 97, 1}, erl_ext.ErlInt(1)},
	{"2", erl_ext.ErlExtBinary{131, 97, 2}, erl_ext.ErlInt(2)},
	{"10", erl_ext.ErlExtBinary{131, 97, 10}, erl_ext.ErlInt(10)},
	{"100", erl_ext.ErlExtBinary{131, 97, 100}, erl_ext.ErlInt(100)},
	{"256", erl_ext.ErlExtBinary{131, 98, 0, 0, 1, 0}, erl_ext.ErlInt(256)},
	{"65536", erl_ext.ErlExtBinary{131, 98, 0, 1, 0, 0}, erl_ext.ErlInt(65536)},
	{"16777216", erl_ext.ErlExtBinary{131, 98, 1, 0, 0, 0}, erl_ext.ErlInt(16777216)},
	{"4294967296", erl_ext.ErlExtBinary{131, 110, 5, 0, 0, 0, 0, 0, 1}, erl_ext.ErlInt(4294967296)},
	{"1099511627776", erl_ext.ErlExtBinary{131, 110, 6, 0, 0, 0, 0, 0, 0, 1}, erl_ext.ErlInt(1099511627776)},
	{"281474976710656", erl_ext.ErlExtBinary{131, 110, 7, 0, 0, 0, 0, 0, 0, 0, 1}, erl_ext.ErlInt(281474976710656)},
	{"72057594037927936", erl_ext.ErlExtBinary{131, 110, 8, 0, 0, 0, 0, 0, 0, 0, 0, 1}, erl_ext.ErlInt(72057594037927936)},
	{"18446744073709551616", erl_ext.ErlExtBinary{131, 110, 9, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, erl_ext.ErlBigInt{plusNine}},
	{"4722366482869645213696", erl_ext.ErlExtBinary{131, 110, 10, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, erl_ext.ErlBigInt{plusTen}},
	{"veryBig", erl_ext.ErlExtBinary{131, 110, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, erl_ext.ErlBigInt{plusVeryBig}},
	{"-1", erl_ext.ErlExtBinary{131, 98, 255, 255, 255, 255}, erl_ext.ErlInt(-1)},
	{"-2", erl_ext.ErlExtBinary{131, 98, 255, 255, 255, 254}, erl_ext.ErlInt(-2)},
	{"-10", erl_ext.ErlExtBinary{131, 98, 255, 255, 255, 246}, erl_ext.ErlInt(-10)},
	{"-100", erl_ext.ErlExtBinary{131, 98, 255, 255, 255, 156}, erl_ext.ErlInt(-100)},
	{"-256", erl_ext.ErlExtBinary{131, 98, 255, 255, 255, 0}, erl_ext.ErlInt(-256)},
	{"-65536", erl_ext.ErlExtBinary{131, 98, 255, 255, 0, 0}, erl_ext.ErlInt(-65536)},
	{"-16777216", erl_ext.ErlExtBinary{131, 98, 255, 0, 0, 0}, erl_ext.ErlInt(-16777216)},
	{"-4294967296", erl_ext.ErlExtBinary{131, 110, 5, 1, 0, 0, 0, 0, 1}, erl_ext.ErlInt(-4294967296)},
	{"-1099511627776", erl_ext.ErlExtBinary{131, 110, 6, 1, 0, 0, 0, 0, 0, 1}, erl_ext.ErlInt(-1099511627776)},
	{"-281474976710656", erl_ext.ErlExtBinary{131, 110, 7, 1, 0, 0, 0, 0, 0, 0, 1}, erl_ext.ErlInt(-281474976710656)},
	{"-72057594037927936", erl_ext.ErlExtBinary{131, 110, 8, 1, 0, 0, 0, 0, 0, 0, 0, 1}, erl_ext.ErlInt(-72057594037927936)},
	{"-18446744073709551616", erl_ext.ErlExtBinary{131, 110, 9, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1}, erl_ext.ErlBigInt{minusNine}},
	{"-4722366482869645213696", erl_ext.ErlExtBinary{131, 110, 10, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, erl_ext.ErlBigInt{minusTen}},
	{"veryBig (negative)", erl_ext.ErlExtBinary{131, 110, 255, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, erl_ext.ErlBigInt{minusVeryBig}},
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
