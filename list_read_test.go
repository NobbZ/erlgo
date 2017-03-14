package erlgo_test

import (
	"github.com/NobbZ/erlgo"
	"testing"
)

var listTestTable = []struct {
	Name   string
	Data   erlgo.ErlExtBinary
	Expect erlgo.Term
}{
	{"empty list", erlgo.FromBytes([]byte{131, 106}), erlgo.NewListFromTerms([]erlgo.Term{})},
	{"short byte list", erlgo.FromBytes([]byte{131, 107, 0, 1, 130}), erlgo.NewListFromTerms([]erlgo.Term{erlgo.Int64(130)})},
	{"medium byte list", erlgo.FromBytes([]byte{131, 107, 125, 0, 130, 130, 130, 130, 130, 130, 130, 130, 130, 130, 130, 130, 130, 130, 130, 130, 130, 130, 130, 130, 130, 130, 130, 130, 130}), erlgo.ProperList{}},
	{"long byte list", erlgo.FromBytes([]byte{131, 107, 255, 254, 130, 130, 130, 130, 130, 130, 130, 130, 130, 130, 130, 130, 130, 130, 130, 130, 130, 130, 130, 130, 130, 130, 130, 130, 130}), erlgo.ProperList{}},
}

func TestReadingLists(t *testing.T) {
	for _, test := range listTestTable {
		t.Run(test.Name, func(t *testing.T) {
			if val, err := test.Data.Decode(); err == nil && !val.Matches(test.Expect) {
				t.Errorf(`%#v parsed into %#v, expected %#v."`, test.Data, val, test.Expect)
			} else if err != nil {
				t.Errorf(`%#v encountered error "%v", expected value %#v.`, test.Data, err, test.Expect)
			}
		})
	}
}

func BenchmarkReadingLists(b *testing.B) {
	for _, data := range floatTestTable {
		b.Run(data.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				data.Data.Decode()
			}
		})
	}
}
