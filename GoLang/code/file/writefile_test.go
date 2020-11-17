package file

import (
	"bufio"
	"fmt"
	"os"
)

// 写文件
func WriteMaptoFile(m map[string]string, filePath string) error {

	f, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("create map file error: %v\n", err)
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	for k, v := range m {
		lineStr := fmt.Sprintf("%s^%s", k, v)
		fmt.Fprintln(w, lineStr)
	}
	return w.Flush()

}
