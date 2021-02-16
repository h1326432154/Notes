package strings

import (
	"fmt"
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
