package utils

import (
    "strings"
    "regexp"
    "fmt"
)

func IsValidEmail(s string) bool {
	strings.Trim(s, " ")
	fmt.Println(s)
	
	rep := regexp.MustCompile("[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(.[a-zA-Z0-9_-]+)+")
	if rep == nil {
			fmt.Println("构建邮箱正则发生错误")
			return false
	}
	result := rep.FindAllStringSubmatch(s, -1)
	return result != nil
}