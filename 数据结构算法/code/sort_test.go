package main

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {
	in := []int{4, 5, 6, 3, 2, 1, 9, 6, 10, 20, 100, 7, 101, 102, 103, 104}
	t.Logf("输入为: %+v", in)
	res := BubbleSort(in)
	t.Logf("冒泡排序结果为: %+v", res)

	res1 := InsertionSort(in)
	t.Logf("插入排序结果为: %+v", res1)
}

// 冒泡排序
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

// 插入排序
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
