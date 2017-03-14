package erlgo_test

import (
	"github.com/NobbZ/erlgo"
	"testing"
)

var floatTestTable = []struct {
	Name   string
	Data   erlgo.ErlExtBinary
	Expect erlgo.Term
}{
	{"0.0 (new format)", erlgo.FromBytes([]byte{131, 70, 0, 0, 0, 0, 0, 0, 0, 0}), erlgo.Float(0.0)},
	{"0.0 (old format)", erlgo.FromBytes([]byte{131, 99, 48, 46, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 101, 43, 48, 48, 0, 0, 0, 0, 0}), erlgo.Float(0.0)},
	{"0.0 (old format, compressed)", erlgo.FromBytes([]byte{131, 80, 0, 0, 0, 32, 120, 156, 75, 54, 208, 51, 192, 2, 82, 181, 13, 12, 24, 64, 0, 0, 104, 41, 5, 114}), erlgo.Float(0.0)},
}

func TestReadingFloats(t *testing.T) {
	for _, test := range floatTestTable {
		t.Run(test.Name, func(t *testing.T) {
			if val, err := test.Data.Decode(); err == nil && !val.Matches(test.Expect) {
				t.Errorf(`%#v parsed into %#v, expected %#v."`, test.Data, val, test.Expect)
			} else if err != nil {
				t.Errorf(`%#v encountered error "%v", expected value %#v.`, test.Data, err, test.Expect)
			}
		})
	}
}

func BenchmarkReadingFloats(b *testing.B) {
	for _, data := range floatTestTable {
		b.Run(data.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				data.Data.Decode()
			}
		})
	}
}
