package leetcode

import (
	"testing"
)

func TestLongestString(t *testing.T) {
	var str string = "abcabcbb"

	l := lengthOfLongestSubstring(str)

	t.Logf("最长不重复长度为 %d", l)
}

func lengthOfLongestSubstring(s string) int {
	var (
		// ASCII
		bStr = []byte(s)
		// K = 字符串中的字符 V = 最后一次出现的位置+1
		bMap = make(map[byte]int)
		// 字符串总长
		lens = len(s)
		// max 记录最大长度，num记录当前循环中的长度
		max, num int
		// 如果字符已经出现过，则i为它上次出现的位置
		i int
	)
	for j := 0; i < lens && j < lens; j++ {
		if bMap[bStr[j]] > i {
			i = bMap[bStr[j]]
		}
		num = j - i + 1
		if num > max {
			max = num
		}
		bMap[bStr[j]] = j + 1
	}
	return max
}
