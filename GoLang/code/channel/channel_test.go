package channel

import (
	"testing"
)

func TestChannel(t *testing.T) {
	ch1 := make(chan int, 3)
	ch1 <- 2
	ch1 <- 1
	ch1 <- 3
	elem1 := <-ch1
	t.Logf("The 1st element received from channel ch1: %v\n", elem1)

	elem2 := <-ch1
	t.Logf("The 2nd element received from channel ch1: %v\n", elem2)

	elem3 := <-ch1
	t.Logf("The 3th element received from channel ch1: %v\n", elem3)
}

/*运行结果
$ go test -v ./channel/channel_test.go
=== RUN   TestChannel
--- PASS: TestChannel (0.00s)
    channel_test.go:13: The 1st element received from channel ch1: 2
    channel_test.go:16: The 2nd element received from channel ch1: 1
    channel_test.go:19: The 3th element received from channel ch1: 3
PASS
*/
