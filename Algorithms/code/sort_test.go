package main

import (
	"fmt"
	"testing"
	"time"
)

func TestSort(t *testing.T) {
	in := []int{4, 5, 6, 3, 2, 1, 9, 6} // , 10, 20, 100, 7, 101, 102, 103, 104
	t.Logf("输入为: %+v", in)

	// 冒泡排序
	// {
	// 	st := time.Now()
	// 	res := []int{}
	// 	// for i := 0; i < 1000; i++ {
	// 	res = BubbleSort(in)
	// 	// }
	// 	et := time.Now()
	// 	t.Logf("冒泡排序结果为: %+v", res)
	// 	t.Logf("冒泡排序用时为: %d", et.S··········· 45ub(st).Nanoseconds())
	// }

	// // 插入排序
	// {
	// 	st := time.Now()
	// 	res := []int{}
	// 	// for i := 0; i < 1000; i++ {
	// 	res = BubbleSort(in)
	// 	// }
	// 	et := time.Now()
	// 	t.Logf("插入排序结果为: %+v", res)
	// 	t.Logf("插入排序用时为: %d", et.Sub(st).Nanoseconds())
	// }

	// 归并排序
	{
		st := time.Now()
		res := []int{}
		// for i := 0; i < 1000; i++ {
		res = MergeSort(in, "in")
		// }
		et := time.Now()
		t.Logf("归并排序结果为: %+v", res)
		t.Logf("归并排序用时为: %d", et.Sub(st).Nanoseconds())
	}
}

// BubbleSort 冒泡排序
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

// InsertionSort 插入排序
func InsertionSort(in []int) []int {

	if len(in) <= 1 {
		return nil
	}

	for i := 1; i < len(in); i++ {
		value := in[i]
		j := i - 1
		// 查找插入的位置
		for ; j >= 0; j-- {
			if in[j] > value {
				in[j+1] = in[j] // 数据移动
			} else {
				break
			}
		}
		in[j+1] = value // 插入数据
	}

	return in
}

// MergeSort 归并排序
func MergeSort(arr []int, a string) []int {
	if len(arr) < 2 {
		return arr
	}
	i := len(arr) / 2
	fmt.Printf("task = %s  i = %d\n", a, i)
	left := MergeSort(arr[0:i], "left")
	fmt.Println("left", left)
	right := MergeSort(arr[i:], "right")
	fmt.Println("right", right)
	result := merge(left, right)
	fmt.Println("result", result)
	return result
}

func merge(left, right []int) []int {
	result := make([]int, 0)
	m, n := 0, 0 // left和right的index位置
	l, r := len(left), len(right)
	for m < l && n < r {
		if left[m] > right[n] {
			result = append(result, right[n])
			n++
			continue
		}
		result = append(result, left[m])
		m++
	}
	result = append(result, right[n:]...) // 这里竟然没有报数组越界的异常？
	result = append(result, left[m:]...)
	return result
}
