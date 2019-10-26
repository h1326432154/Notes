package code

import "testing"

// TestArrayBounds 数组边界测试
func TestArrayBounds(t *testing.T) {
	i := 0
	arr := [3]int{0}
	for ; i < 3; i++ {
		arr[i] = 0
		t.Logf("====%d", arr[i])
	}
}
