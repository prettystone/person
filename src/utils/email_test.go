package utils

import "testing"

func TestIsValidEmail(t *testing.T) {
	arr := [...]string {
		"a@126.com",
		"123@163.com",
	}
	for _, v := range arr {
		r := IsValidEmail(v)
		if !r {
			t.Errorf("email %s, expected true.", v)
		}
	}
}