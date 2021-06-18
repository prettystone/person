package utils

import "testing"

func TestIsValid(t *testing.T) {
	arr := [...]string {
		"6259064220951256",
		"5242064266521528",
		"5242064279349313",
		"4693804290023810",
		"6259064298790537",
		"6259074247313794",
		"6283880273521591",
		"6283881047522543",
		"6283883536667359",
		"6283885660841668",
		"6283885802798230",
		"6283887300768517",
		"5566324213235665",
		"6259064283831130",
		"6259064225837757",
		"6259064227598704",
		"6259064250019644",
		"6227654224101384",
		"6259064277325040",
		"6259064229288023",
	}
	for _, v := range arr {
		r := IsValidBank(v)
		if !r {
			t.Errorf("bankid %s, expected true.", v)
		}
	}
}

func TestIsValidFalse(t *testing.T) {
	arr := [...]string {
		"6259064277325042",
		"6259064229288025",
	}
	for _, v := range arr {
		r := IsValid(v)
		if r {
			t.Errorf("bankid %s, expected false.", v)
		}
	}
}