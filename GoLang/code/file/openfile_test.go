package file

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

// 读文件
func TestOpenFile(t *testing.T) {
	fileName := "./aaa"
	file, err := os.OpenFile(fileName, os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("Open file error!", err)
		return
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}

	var size = stat.Size()
	fmt.Println("file size=", size)

	buf := bufio.NewReader(file)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		fmt.Println(line)
		if err != nil {
			if err == io.EOF {
				fmt.Println("File read ok!")
				break
			} else {
				fmt.Println("Read file error!", err)
				return
			}
		}
	}
}

// 测试结果
// $ go test -v ./file/
// === RUN   TestOpenFile
// file size= 13
// 123123,456456
// 789789,123123
// File read ok!
// --- PASS: TestOpenFile (0.00s)
// PASS
// ok  	github.com/h1326432154/common/file	0.003s
