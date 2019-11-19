package typeconversion

import "testing"

func TestTypeReName(t *testing.T) {

	type MyString = string
	type MyString2 string

	var str1 string
	var str2 MyString
	var str3 MyString2

	str1 = "string"

	str2 = str1

	// 此处这样写将会报错cannot use str2 (type string) as type MyString2
	// str3 = str2

	t.Logf("str1 = %s", str1)

	t.Logf("str2 = %s", str2)

	t.Logf("str3 = %s", str3)

}

/*运行结果
go test -v ./typeconversion/rename_test.go
=== RUN   TestTypeReName
--- PASS: TestTypeReName (0.00s)
    rename_test.go:20: str1 = string
    rename_test.go:22: str2 = string
    rename_test.go:24: str3 =
PASS
*/
