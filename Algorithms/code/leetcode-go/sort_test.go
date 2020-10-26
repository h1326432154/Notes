package leetcode

import (
	"fmt"
	"sort"
	"testing"
	"time"
)

func TestSort(t *testing.T) {

	a := time.Now()
	b := time.Now().AddDate(0, 0, 1)
	fmt.Println(b.Sub(a).Hours() / 24)
	if b.Before(a) {
		fmt.Println(a, b)
	} else {

		fmt.Println("11111")
	}
	// 2020-04-08 11:54:01.0160429 +0800 CST m=+0.001021601 2020-04-09 11:54:01.0160429 +0800 CST
	// ret := a.Sub(time.Now()).String()

	if false {

		orders := make(map[int][]*OrderTimeInfo)
		o1 := &OrderTimeInfo{ID: 8, UID: 1, Time: 999999}
		o2 := &OrderTimeInfo{ID: 10, UID: 1, Time: 123}
		o3 := &OrderTimeInfo{ID: 3, UID: 1, Time: 4567}
		o4 := &OrderTimeInfo{ID: 7, UID: 1, Time: 444}
		o5 := &OrderTimeInfo{ID: 5, UID: 1, Time: 0}

		orders[111] = append(orders[111], o5)
		orders[111] = append(orders[111], o4)
		orders[111] = append(orders[111], o3)
		orders[111] = append(orders[111], o2)
		orders[111] = append(orders[111], o1)
		orders[222] = append(orders[222], o5)
		orders[222] = append(orders[222], o4)
		orders[222] = append(orders[222], o3)
		orders[222] = append(orders[222], o2)
		orders[222] = append(orders[222], o1)
		doSort(orders)
	}

}

// OrderTimeInfo 订单信息
type OrderTimeInfo struct {
	ID   int   `gorm:"column:o_id"`
	UID  int   `gorm:"column:o_u_id"`
	Time int64 `gorm:"column:o_creattime"`
	// Date int64 `gorm:"column:o_date"`
}

func doSort(orders map[int][]*OrderTimeInfo) {

	for _, order := range orders {
		sort.Sort(SortBy(order))
	}
	fmt.Printf("%+v", orders[111][0])
	fmt.Printf("%+v", orders[111][1])
	fmt.Printf("%+v", orders[111][2])
	fmt.Printf("%+v", orders[111][3])
	fmt.Printf("%+v", orders[111][4])
}

// SortBy 切片排序
type SortBy []*OrderTimeInfo

func (a SortBy) Len() int           { return len(a) }
func (a SortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortBy) Less(i, j int) bool { return a[i].Time < a[j].Time }
