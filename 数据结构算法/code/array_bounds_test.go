package main

import "testing"

// TestArrayBounds 数组边界测试
func TestArrayBounds(t *testing.T) {
	i := 0
	arr := [3]int{0}
	for ; i <= 4; i++ {
		arr[i] = 0
		t.Logf("i = %d | ", i)
		t.Logf("arr[i] = %d | ", arr[i])
	}
}
