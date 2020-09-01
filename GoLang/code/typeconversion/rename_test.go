package typeconversion

import "testing"

func TestTypeReName(t *testing.T) {

	type MyString = string
	type MyString2 string

	var str1 string
	var str2 MyString
	var str3 MyString2

	var str4 []string
	var str5 []MyString
	var str6 []MyString2

	str1 = "string"
	str2 = str1
	// 此处这样写将会报错cannot use str2 (type string) as type MyString2
	// str3 = str2
	str4 = []string{"a", "b", "c"}
	str5 = str4

	// str6 = str5

	t.Logf("str1 = %s", str1)
	t.Logf("str2 = %s", str2)
	t.Logf("str3 = %s", str3)
	t.Logf("str4 = %s", str4)
	t.Logf("str5 = %s", str5)
	t.Logf("str6 = %s", str6)

}

/*运行结果
$ go test -v ./typeconversion/rename_test.go
=== RUN   TestTypeReName
--- PASS: TestTypeReName (0.00s)
    rename_test.go:27: str1 = string
    rename_test.go:28: str2 = string
    rename_test.go:29: str3 =
    rename_test.go:30: str4 = [a b c]
    rename_test.go:31: str5 = [a b c]
    rename_test.go:32: str6 = []
PASS
*/
