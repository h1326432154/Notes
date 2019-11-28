package cache

import (
	"fmt"
	"testing"
	"time"
)

func fibonacci(a int) (result int) {
	//递归需要一个出口
	if a == 1 || a == 2 {
		result = 1
		return result
	}
	//开始递归，前一个数加后一个数，一直调用，直到出口为止
	result = fibonacci(a-1) + fibonacci(a-2)
	return result
}
func TestFibonacci(t *testing.T) {
	var i int = 45
	//开始调用
	start := time.Now()
	num := fibonacci(i)
	fmt.Printf("第%d个斐波那契数列为%d", i, num)
	end := time.Now()
	dur := end.Sub(start)
	fmt.Println("运行时间为:", dur)
}

const LIM = 45

var fibs [LIM]uint64

func fibonacci1(n int) (res uint64) {
	// 缓存: 检查是否已经计算过，如已有值则直接提取。
	if fibs[n] != 0 {
		res = fibs[n]
		return
	}
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci1(n-1) + fibonacci1(n-2)
	}
	//计算完毕，加入缓存中，以备下次之需
	fibs[n] = res
	return
}

func TestFibonacciA(t *testing.T) {
	var result uint64 = 0
	start := time.Now()
	for i := 0; i < LIM; i++ {
		result = fibonacci1(i)
		fmt.Printf("fibonacci(%d) is: %d\n", i, result)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)
}

func TestFibonacciB(t *testing.T) {
	var result uint64 = 0
	start := time.Now()
	for i := 0; i < LIM; i++ {
		result = fibonacci1(i)
		fmt.Printf("fibonacci(%d) is: %d\n", i, result)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)
}

func fibonacciB(i int) {

}
