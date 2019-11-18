package scope

import (
	"fmt"
	"testing"
)

var block = "package"

// 程序实体作用域
func TestScope(t *testing.T) {

	block := "function"
	{
		block := "inner"
		fmt.Printf("The block is %s.\n", block)
	}
	fmt.Printf("The block is %s.\n", block)
}

// 测试结果
// $ go test -v ./scope/
// === RUN   TestScope
// The block is inner.
// The block is function.
// --- PASS: TestScope (0.00s)
