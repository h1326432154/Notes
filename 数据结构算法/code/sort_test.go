package main

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {
	in := []int{4, 5, 6, 3, 2, 1, 9, 6, 10, 20, 100, 7, 101, 102, 103, 104}
	t.Logf("输入为: %+v", in)
	res := BubbleSort(in)
	t.Logf("冒泡排序结果为: %+v", res)
}

func BubbleSort(in []int) []int {
	for i := 0; i < len(in); i++ {
		flag := false
		for j := 0; j < len(in)-i-1; j++ {
			if in[j] > in[j+1] {
				in[j] = in[j] ^ in[j+1]
				in[j+1] = in[j] ^ in[j+1]
				in[j] = in[j] ^ in[j+1]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
	return in
}
