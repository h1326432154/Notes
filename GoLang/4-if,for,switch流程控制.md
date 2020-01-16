#

## 使用携带range子句的for语句时需要注意的细节

```go
numbers1 := []int{1, 2, 3, 4, 5, 6}
for i := range numbers1 {
  if i == 3 {
    numbers1[i] |= i
  }
}
fmt.Println(numbers1)
```
