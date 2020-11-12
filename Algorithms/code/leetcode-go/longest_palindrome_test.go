package leetcode

import (
	"fmt"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	strs := []string{"cccc", "cbbd", "bb", "a", "ac", "babad", "abcba", "abba"}
	// strs := []string{"cbbd"}
	for _, s := range strs {
		res := longestPalindrome(s)
		fmt.Printf("最长回文子串为: %s \n", res)
	}
}

func longestPalindrome(s string) string {
	if s == "" {
		return s
	}
	m := make(map[int]map[int]bool)
	loopMap := make(map[int]bool)
	l := len(s)
	b := 0
	e := 0
	for i := 0; i < l; i++ {
		im := make(map[int]bool)
		m[i] = im
	}
	for loop := 0; loop < l; loop++ {
		if loop > 3 && !loopMap[loop-2] {
			continue
		}
		for i := 0; i < l-loop; i++ {
			if loop == 0 || loop == 1 && string(s[i]) == string(s[i+1]) || m[i+1][i+loop-1] && string(s[i]) == string(s[i+loop]) {
				m[i][i+loop] = true
				b = i
				e = i + loop
				loopMap[loop] = true
			}
		}
	}
	return s[b : e+1]
}

func longestPalindrome1(s string) string {
	var (
		byt    = []byte(s)
		ls     = len(byt)
		maxStr = []byte{}
		res    = []byte{}
	)
	if len(byt) <= 1 {
		return string(byt)
	}
	for i := 0; i < ls-1; i++ {
		if byt[i] == byt[i+1] {
			fmt.Println("双", i)
			res = byt[i : i+2]
			for j := 0; j <= i && j < ls-i-1; j++ {
				if byt[i] == byt[i+j+1] {
					maxStr = byt[i : i+j+2]
					fmt.Println(len(maxStr), len(res))
					if len(maxStr) > len(res) {
						res = maxStr
					}
				}
				if byt[i-j] == byt[i+j+1] {
					maxStr = byt[i-j : i+j+2]
					if len(maxStr) > len(res) {
						res = maxStr
					}
				}
			}
		} else {
			fmt.Println("单", i)
			for j := 0; j <= i && j <= ls-i-1; j++ {
				if byt[i-j] == byt[i+j] {
					maxStr = byt[i-j : i+j+1]
				}
				if len(maxStr) > len(res) {
					res = maxStr
				}
			}
		}
	}
	return string(res)
}
