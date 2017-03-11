package erl_ext_test

import (
	"github.com/NobbZ/erl_ext"
	"testing"
)

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
}

func TestReadingIntegers(t *testing.T) {
	for _, test := range table {
		t.Run(test.Name, func(t *testing.T) {
			if val, err := test.Data.Decode(); val != test.Expect {
				if err == nil {
					t.Errorf(`%#v parsed into %#v, expected %#v."`, test.Data, val, test.Expect)
				} else {
					t.Errorf(`%#v encountered error "%v", expected value %#v.`, test.Data, err, test.Expect)
				}
			}
		})
	}
}
