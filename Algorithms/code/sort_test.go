package main

import (
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {
	// 冒泡排序
	{
		t.Log("冒泡排序")
		in := []int{4, 5, 6, 3, 2, 1, 9, 6} // , 10, 20, 100, 7, 101, 102, 103, 104
		t.Logf("输入为: %+v", in)
		res := BubbleSort(in)
		t.Logf("冒泡排序结果为: %+v", res)
		fmt.Println()
	}

	// 选择排序
	{
		t.Log("选择排序")
		in := []int{4, 5, 6, 3, 2, 1, 9, 6} // , 10, 20, 100, 7, 101, 102, 103, 104
		t.Logf("输入为: %+v", in)
		res := selectionSort(in)
		t.Logf("选择排序结果为: %+v", res)
		fmt.Println()
	}

	// 插入排序
	{
		t.Log("插入排序")
		in := []int{4, 5, 6, 3, 2, 1, 9, 6} // , 10, 20, 100, 7, 101, 102, 103, 104
		t.Logf("输入为: %+v", in)
		res := BubbleSort(in)
		t.Logf("插入排序结果为: %+v", res)
		fmt.Println()
	}

	// 归并排序
	{
		t.Log("归并排序")
		in := []int{4, 5, 6, 3, 2, 1, 9, 6} // , 10, 20, 100, 7, 101, 102, 103, 104
		t.Logf("输入为: %+v", in)
		res := MergeSort(in, "in")
		t.Logf("归并排序结果为: %+v", res)
		fmt.Println()
	}

}

// BubbleSort 冒泡排序
func BubbleSort(in []int) []int {
	for i := 0; i < len(in); i++ {
		flag := false
		for j := 0; j < len(in)-i-1; j++ {
			if in[j] > in[j+1] {
				// in[j] = in[j] ^ in[j+1]
				// in[j+1] = in[j] ^ in[j+1]
				// in[j] = in[j] ^ in[j+1]
				in[j], in[j+1] = in[j+1], in[j]
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
	left := MergeSort(arr[0:i], "left")
	right := MergeSort(arr[i:], "right")
	result := merge(left, right)
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

// selectionSort 选择排序
func selectionSort(arr []int) []int {
	l := len(arr)
	if l == 0 {
		return arr
	}
	for i := 0; i < l; i++ {
		min := i
		for j := i + 1; j < l; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
	return arr
}

// 快排
func quickSort(arr []int) []int {
	return _quickSort(arr, 0, len(arr)-1)
}

func _quickSort(arr []int, left, right int) []int {
	if left < right {
		partitionIndex := partition(arr, left, right)
		_quickSort(arr, left, partitionIndex-1)
		_quickSort(arr, partitionIndex+1, right)
	}
	return arr
}

func partition(arr []int, left, right int) int {
	pivot := left
	index := pivot + 1

	for i := index; i <= right; i++ {
		if arr[i] < arr[pivot] {
			swap(arr, i, index)
			index++
		}
	}
	swap(arr, pivot, index-1)
	return index - 1
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
