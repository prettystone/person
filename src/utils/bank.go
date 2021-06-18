package bank

import (
    "strings"
    "regexp"
    "fmt"
)

var ARR = [...]int{0, 2, 4, 6, 8, 1, 3, 5, 7, 9}
func ValidLuhn(s string) bool {
    odd := len(s) & 1
    var sum int
    for i, c := range s {
        if c < '0' || c > '9' {
            return false
        }
        if i&1 == odd {
            sum += ARR[c-'0']
        } else {
            sum += int(c - '0')
        }
    }
    return sum%10 == 0
}

func IsValid(s string) bool {
    strings.Trim(s, " ")
    strings.Replace(s, " ", "", -1)
    strings.Replace(s, "-", "", -1)
    fmt.Println(s)
    
    rep := regexp.MustCompile("([1-9]{1}[0-9]{14,18})")
    if rep == nil {
        fmt.Println("构建银行号正则发生错误")
        return false
    }
    result := rep.FindAllStringSubmatch(s, -1)
    if result == nil {
        return false
    }
    return ValidLuhn(s)
}