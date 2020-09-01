package typeconversion

import (
	"testing"
)

func TestString(t *testing.T) {
	var str string
	str = "你好"
	t.Log(str)
	byt := []byte(str)
	t.Logf("string to []byte %+v", byt)
	run := []rune(str)
	t.Logf("string to []rune %+v", run)
}

/*输出结果
$ go test -v ./typeconversion/string_test.go
=== RUN   TestString
--- PASS: TestString (0.00s)
    string_test.go:10: 你好
    string_test.go:12: string to []byte [228 189 160 229 165 189]
    string_test.go:14: string to []rune [20320 22909]
PASS
*/
