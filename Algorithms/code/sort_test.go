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
		res := bubbleSort(in)
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
		res := insertionSort(in)
		t.Logf("插入排序结果为: %+v", res)
		fmt.Println()
	}

	// 归并排序
	{
		t.Log("归并排序")
		in := []int{4, 5, 6, 3, 2, 1, 9, 6} // , 10, 20, 100, 7, 101, 102, 103, 104
		t.Logf("输入为: %+v", in)
		res := mergeSort(in, "in")
		t.Logf("归并排序结果为: %+v", res)
		fmt.Println()
	}

}

// bubbleSort 冒泡排序
func bubbleSort(in []int) []int {
	for i := 0; i < len(in); i++ {
		flag := false
		for j := 0; j < len(in)-1-i; j++ {
			if in[j] > in[j+1] {
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

// selectionSort 选择排序
func selectionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
	return arr
}

// insertionSort 插入排序
func insertionSort(in []int) []int {
	for i := 1; i < len(in); i++ {
		value := in[i]
		j := i - 1
		for ; j >= 0; j-- {
			if in[j] > value {
				in[j+1] = in[j]
			} else {
				break
			}
		}
		in[j+1] = value
	}
	return in
}

// func mergeSort(arr []int) []int {
// 	length := len(arr)
// 	if length < 2 {
// 		return arr
// 	}
// 	middle := length / 2
// 	left := arr[0:middle]
// 	right := arr[middle:]
// 	return merge(mergeSort(left), mergeSort(right))
// }

// func merge(left []int, right []int) []int {
// 	var result []int
// 	for len(left) != 0 && len(right) != 0 {
// 		if left[0] <= right[0] {
// 			result = append(result, left[0])
// 			left = left[1:]
// 		} else {
// 			result = append(result, right[0])
// 			right = right[1:]
// 		}
// 	}

// 	for len(left) != 0 {
// 		result = append(result, left[0])
// 		left = left[1:]
// 	}

// 	for len(right) != 0 {
// 		result = append(result, right[0])
// 		right = right[1:]
// 	}
// 	return result
// }

func mergeSort(arr []int) []int {
	lens := len(arr)
	if lens < 2 {
		return arr
	}
	mid := lens / 2
	left := arr[:mid]
	right := arr[mid:]
	return merge(mergeSort(left), mergeSort(right))
}

func merge(left []int, right []int) []int {
	var result = []int{}
	for len(left) > 0 && len(right) > 0 {
		if left[0] <= right[0] {

		}
	}

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
