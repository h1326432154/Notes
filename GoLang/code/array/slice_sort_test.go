package array

import (
	"fmt"
	"sort"
	"testing"
)

func TestSliceSort(t *testing.T) {

	o := []*sortTest{
		{1, 4},
		{1, 5},
		{1, 3},
	}

	fmt.Printf("%+v  %+v  %+v", o[0], o[1], o[2])
	sort.Sort(SortByTime(o))
	fmt.Printf("%+v  %+v  %+v", o[0], o[1], o[2])
}

// sortTest 订单基础信息
type sortTest struct {
	ID    int
	OTime int
}

// SortByTime 根据订单下单时间排序
type SortByTime []*sortTest

func (a SortByTime) Len() int           { return len(a) }
func (a SortByTime) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortByTime) Less(i, j int) bool { return a[i].OTime < a[j].OTime }
