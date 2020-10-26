package leetcode

import (
	"fmt"
	"strconv"
	"testing"
)

func TestGetEnv(t *testing.T) {
	a := "6395884‬" // “6395884‬”
	// a := "6395884"
	b, _ := strconv.ParseInt(a, 10, 64)
	fmt.Println(b)
}
