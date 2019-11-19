package assert

import (
	"testing"
)

var container = []string{"zero", "one", "two"}

// TestAssert 断言
func TestAssert(t *testing.T) {
	container := map[int]string{0: "zero", 1: "one", 2: "two"}

	t.Logf("The element is %q.\n", container[1])

	value, ok := interface{}(container).(map[int]string)
	if !ok {
		t.Error("assert failed")
	}

	t.Log(value)
}

/*输出结果
go test -v ./assert/assert_test.go
=== RUN   TestAssert
--- PASS: TestAssert (0.00s)
    assert_test.go:13: The element is "one".
    assert_test.go:20: map[0:zero 1:one 2:two]
PASS
*/
