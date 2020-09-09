package goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestGoroutine(t *testing.T) {
	num := 10

	var wg sync.WaitGroup
	for i := 0; i < num; i++ {
		wg.Add(1)

		go func(i int) {

			fmt.Println(i)

			wg.Done()
		}(i)

	}

	wg.Wait()
}

// TestGoroutine1 开始执行
func TestA(t *testing.T) {
	ThreadPool()
}

func ThreadPool() {
	num := 100
	loopNum := 3
	//Job为100个可以传递int类型的channel
	jobs := make(chan int, num)
	results := make(chan int, loopNum)

	//开启三个线程，说明线程池中只有三个线程， 在实际情况下可以动态设置开启线程数量
	for w := 1; w <= loopNum; w++ {
		go worker(w, jobs, results)

		// 添加9个任务
		for j := 1; j <= num; j++ {
			jobs <- j
			//向Jobs添加任务： 向Channel中写入数据， 传递的数据类型为int
		}

		//关闭Channel
		close(jobs)
		for a := 1; a <= num; a++ {
			res := <-results
			//从Channel中读取数据, 输出的数据类型为 int
			fmt.Println("ID = ", a, "Results = ", res)
			//注意 a 与 res 不对应是由于处理器调度的结果
		}
	}
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", "processing job", j)

		// time.Sleep(time.Second)
		results <- j
	}
}

func TestGoroutine1(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	time.Sleep(time.Millisecond * 50)
}
