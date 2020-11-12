package leetcode

import (
	"fmt"
	"testing"
)

func TestMedian(t *testing.T) {
	var (
		nums1 = []int{1, 3, 6, 7}
		nums2 = []int{2, 8, 9}
	)
	res := findMedianSortedArrays(nums1, nums2)
	t.Logf("中位数为%f", res)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	if m > n {
		nums1, nums2 = nums2, nums1
		// m = min(m,n)
		m, n = n, m
	}
	fmt.Printf("m = %d, n = %d \n", m, n)
	iMin, iMax, halfLen := 0, m, (m+n+1)/2
	fmt.Printf("halfLen = %d \n", halfLen)
	for iMin <= iMax {
		i := (iMax + iMin) / 2
		j := halfLen - i
		fmt.Printf("i = %d \n", i)
		fmt.Printf("j = %d \n", j)
		// i < 较短数组
		fmt.Printf("nums2[j-1] = %d ,nums1[i] = %d \n", nums2[j], nums1[i])
		fmt.Printf("nums2[j] = %d ,nums1[i-1] = %d \n", nums2[j], nums1[i-1])
		if i < iMax && nums2[j-1] > nums1[i] {
			iMin = i + 1
		} else if i > iMin && nums2[j] < nums1[i-1] {
			iMax = i - 1
		} else {
			maxLeft := 0
			if i == 0 {
				maxLeft = nums2[j-1]
			} else if j == 0 {
				maxLeft = nums1[i-1]
			} else {
				maxLeft = max(nums2[j-1], nums1[i-1])
			}
			if (m+n)%2 == 1 { //奇数，中位数正好是maxLeft
				return float64(maxLeft)
			}
			minRight := 0
			if i == m {
				minRight = nums2[j]
			} else if j == n {
				minRight = nums1[i]
			} else {
				minRight = min(nums1[i], nums2[j])
			}
			return float64(minRight+maxLeft) / 2.0
		}
		fmt.Println("-----------------------")
	}
	return 0.0
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
