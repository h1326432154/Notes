package leetcode

import (
	"fmt"
	"math"
	"strings"
	"testing"
)

func TestMyAtoi(t *testing.T) {

	var strs = []string{"-42", "4193 with words", "words and 987", "91283472332"}
	for _, str := range strs {
		out := myAtoi(str)
		t.Log(out)
	}
}

// Atoi
func myAtoi(str string) int {
	str = strings.TrimSpace(str)
	if str == "" {
		return 0
	}
	u := 1
	if str[0] == '-' || str[0] == '+' {
		if str[0] == '-' {
			u = -1
		}
		str = str[1:]
	}
	for i, b := range str {
		if b < '0' || b > '9' {
			str = str[:i]
			break
		}
	}
	out := 0
	for _, v := range str {
		out = out*10 + int(v-'0')
	}

	if u*out > math.MaxInt32 {
		fmt.Println("++++", out)
		return math.MaxInt32
	}
	if u*out < math.MinInt32 {
		fmt.Println("----", out)
		return math.MinInt32
	}
	return u * out
}
