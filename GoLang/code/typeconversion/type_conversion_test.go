package typeconversion

import "testing"

// TestTpypeConversion 类型转换
func TestTpypeConversion(t *testing.T) {
	var srcInt = int16(-255)
	dstInt := int8(srcInt)

	t.Logf("int16(-255) -> int8 = %d", dstInt)

	str := string(-1)
	t.Logf("int(-1) -> string = %s", str)
}

/*输出结果
$ go test -v ./typeconversion/type_conversion_test.go
=== RUN   TestTpypeConversion
--- PASS: TestTpypeConversion (0.00s)
    type_conversion_test.go:10: int16(-255) -> int8 = 1
    type_conversion_test.go:13: int(-1) -> string = �
PASS
*/

// int16(-255) 的源码为0000000011111111 补码为 1111111100000001
// int(int16(-255))的源码为 00000001
// 所以输出为1
