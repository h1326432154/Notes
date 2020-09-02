package array

import (
	"testing"
)

// TestArraySilce 数组切片的长度与容量
func TestArraySilce(t *testing.T) {
	s1 := make([]int, 2)
	t.Logf("The length of s1: %d\n", len(s1))
	t.Logf("The capacity of s1: %d\n", cap(s1))
	t.Logf("The value of s1: %d\n", s1)
	s2 := make([]int, 5, 8)
	t.Logf("The length of s2: %d\n", len(s2))
	t.Logf("The capacity of s2: %d\n", cap(s2))
	t.Logf("The value of s2: %d\n", s2)
}

/*测试结果
$ go test -v ./array/array_slice_test.go
=== RUN   TestArraySilce
--- PASS: TestArraySilce (0.00s)
    array_slice_test.go:10: The length of s1: 5
    array_slice_test.go:11: The capacity of s1: 5
    array_slice_test.go:12: The value of s1: [0 0 0 0 0]
    array_slice_test.go:14: The length of s2: 5
    array_slice_test.go:15: The capacity of s2: 8
    array_slice_test.go:16: The value of s2: [0 0 0 0 0]
PASS
*/

// TestSlice 切片与数组关系
// s1 = s[m,n] 其中 m>=0 n <= len(s)
// 下标m开始截取，截取长度为n-m(注：此处并非截取到下标n的位置)
// s1长度为 n-m  容量为len(s)-m
func TestSlice(t *testing.T) {
	s3 := make([]int, 9, 12)
	s3 = []int{1, 2, 3, 4, 5, 6, 7, 8}
	s4 := s3[3:6]
	t.Logf("The length of s4: %d\n", len(s4))
	t.Logf("The capacity of s4: %d\n", cap(s4))
	t.Logf("The value of s4: %v\n", s4)
	t.Logf("The value of s4: %d\n", s4[0:cap(s4)])
}

/*运行结果
$ go test -v ./array/array_slice_test.go --test.run TestSlice
=== RUN   TestSlice
--- PASS: TestSlice (0.00s)
    array_slice_test.go:36: The length of s4: 3
    array_slice_test.go:37: The capacity of s4: 5
	array_slice_test.go:38: The value of s4: [4 5 6]
	array_slice_test.go:39: The value of s4: [4 5 6 7 8]
PASS
*/

func TestSlice2(t *testing.T) {
	// 原长度为10，容量为20
	s8 := make([]int, 10, 20)
	t.Logf("s8: len: %d, cap: %d\n", len(s8), cap(s8))

	// 新增较小长度(长度相加小于原容量) 长度将相加，容量不变
	s8a := append(s8, make([]int, 9, 20)...)
	t.Logf("s8a: len: %d, cap: %d\n", len(s8a), cap(s8a))

	// 新增长度与原长度相加 大于原容量且小于2倍原容量 长度将相加，容量将为原长度的2倍
	s8b := append(s8a, make([]int, 11)...)
	t.Logf("s8b: len: %d, cap: %d\n", len(s8b), cap(s8b))

	// 新增长度与原长度相加 大于2倍原容量 长度将相加，容量 > 长度相加 容量等于长度相加后首个大于该长度的8的公倍数
	// 例如 30+58时 cap = 88  30+59时 cap = 96
	s8c := append(s8b, make([]int, 59)...)
	t.Logf("s8c: len: %d, cap: %d\n", len(s8c), cap(s8c))

}

/* 运行结果
$ go test -v ./array/array_slice_test.go --test.run TestSlice2
=== RUN   TestSlice2
    array_slice_test.go:60: s8: len: 10, cap: 20
    array_slice_test.go:64: s8a: len: 19, cap: 20
    array_slice_test.go:68: s8b: len: 30, cap: 40
    array_slice_test.go:72: s8c: len: 97, cap: 112
--- PASS: TestSlice2 (0.00s)
*/

// TestSlice 切片与数组关系
func TestSlice3(t *testing.T) {
	s3 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	s4 := s3[3:6]
	s5 := s3[5:8]

	t.Logf("The value of s: %d\n", s4)
	t.Logf("The value of s4: %d\n", s5)

	s4[2] = 100

	t.Logf("The value of s3: %d\n", s3)
	t.Logf("The value of s4: %d\n", s4)
	t.Logf("The value of s5: %d\n", s5)
}

/*运行结果
$ go test -v ./array/array_slice_test.go --test.run TestSlice3
=== RUN   TestSlice3
--- PASS: TestSlice3 (0.00s)
    array_slice_test.go:78: The value of s: [4 5 6]
    array_slice_test.go:79: The value of s4: [6 7 8]
    array_slice_test.go:83: The value of s3: [1 2 3 4 5 100 7 8]
    array_slice_test.go:84: The value of s4: [4 5 100]
    array_slice_test.go:85: The value of s5: [100 7 8]
PASS
*/

func TestSlice4(t *testing.T) {
	a := [3]int{1, 2, 3}
	for k, v := range a {
		if k == 0 {
			a[0], a[1] = 100, 200
			t.Log(a)
		}
		a[k] = 100 + v
	}
	t.Log(a)

	b := []int{1, 2, 3}
	for k, v := range b {
		if k == 0 {
			b[0], b[1] = 100, 200
			t.Log(b)
		}
		b[k] = 100 + v
		t.Log("1111111====", b)
	}
	t.Log(b)
}
