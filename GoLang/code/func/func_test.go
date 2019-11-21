package function

import (
	"fmt"
	"testing"
)

type Printer func(contents string) (n int, err error)

func printToStd(contents string) (bytesNum int, err error) {
	return fmt.Println(contents)
}

func TestFunc(t *testing.T) {
	var p Printer
	p = printToStd
	p("something")
}
