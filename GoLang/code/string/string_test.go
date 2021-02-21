package strings

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"unsafe"
)

func TestStrMemory(t *testing.T) {
	// var m runtime.MemStats
	// runtime.ReadMemStats(&m)
	// fmt.Printf("%+v\n", m.Sys)

	s := "666666666666666666666666666666666666666666666666666666666"
	fmt.Printf("%d\n", unsafe.Sizeof(s))
	s = ""
	fmt.Printf("%d\n", unsafe.Sizeof(s))

}

func BenchmarkStringPlus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringPlus()
	}
}

func StringPlus() string {
	var s string
	s += "666666666"
	s += "7777777777777777777777"
	s += "888888888888888888888"
	return s
}

func BenchmarkStringFmt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringFmt()
	}
}

func StringFmt() string {
	return fmt.Sprint("666666666", "7777777777777777777777", "888888888888888888888")
}

func StringJoin() string {
	s := []string{"666666666", "7777777777777777777777", "888888888888888888888"}
	return strings.Join(s, "")
}

func BenchmarkStringJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringJoin()
	}
}

func StringBuffer() string {
	var b bytes.Buffer
	b.WriteString("666666666")
	b.WriteString("7777777777777777777777")
	b.WriteString("888888888888888888888")
	return b.String()
}

func BenchmarkStringBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringBuffer()
	}
}

func StringBuilder() string {
	var b strings.Builder
	b.WriteString("666666666")
	b.WriteString("7777777777777777777777")
	b.WriteString("888888888888888888888")
	return b.String()
}

func BenchmarkStringBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringBuilder()
	}
}
