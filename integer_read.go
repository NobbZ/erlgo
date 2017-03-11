package erl_ext

import "fmt"

type ErlInt int64

func (ei ErlInt) ToInteger() (int64, error) {
	return int64(ei), nil
}

func (ei ErlInt) IsInteger() bool {
	return true
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
