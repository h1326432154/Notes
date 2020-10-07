package maps

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	// maina1()
	maina2()
	// maina3()
	// maina4()
	// maina5()
	// maina6()
	// maina7()
}

//JSON的序列化
//将结构体构成数据，并转JSON
//将map[string]interface{}构成数据，并转JSON
//使用map切片构成数据,并转JSON

//json.cn
//将结构体构成数据，并转JSON
//使用Man 结构体构成数据，并转为json
type Man struct {
	Name  string
	Age   int
	Hobby []string
	Sex   bool
}

func maina1() {
	person := Man{"liyi", 20, []string{"唱歌", "跳舞"}, true}
	bytes, err := json.Marshal(person)
	if err != nil {
		fmt.Println("序列化失败，err=", err)
	} else {
		fmt.Println("序列化成功，err=", err)
		fmt.Println(string(bytes))
	}
}

//将map[string]interface{}构成数据，并转JSON
//这里interface{}代表任意数据类型
func maina2() {
	dataMap := make(map[string]interface{})
	dataMap["name"] = "liyi"
	dataMap["age"] = 20
	dataMap["hobby"] = []string{"唱歌", "跳舞"}
	dataMap["sex"] = true
	bytes, err := json.Marshal(dataMap)
	if err != nil {
		fmt.Println("序列化失败，err=", err)
	} else {
		fmt.Println("序列化成功，err=", err)
		fmt.Println(string(bytes))
	}
}

//使用map切片构成数据,并转JSON
func maina3() {
	dataMap1 := make(map[string]interface{})
	dataMap1["name"] = "liyi"
	dataMap1["age"] = 20
	dataMap1["hobby"] = []string{"唱歌", "跳舞"}
	dataMap1["sex"] = true

	dataMap2 := make(map[string]interface{})
	dataMap2["name"] = "linging"
	dataMap1["age"] = 21
	dataMap1["hobby"] = []string{"打炮"}
	dataMap1["sex"] = true

	dataMap3 := make(map[string]interface{})
	dataMap3["name"] = "yiyi"
	dataMap3["age"] = 18
	dataMap3["hobby"] = []string{"学习"}
	dataMap3["sex"] = true

	dataSlice := make([]map[string]interface{}, 0)
	dataSlice = append(dataSlice, dataMap1, dataMap2, dataMap3)
	bytes, err := json.Marshal(dataSlice)
	if err != nil {
		fmt.Println("序列化失败，err=", err)
	} else {
		fmt.Println("序列化成功")
		fmt.Println(string(bytes))
	}
}

//JSON的反序列化
//将json数据转为map
//将json数据转为结构体
//将json数据转换为map切片
//将json数据转为结构体切片

//将json数据转为map
func maina4() {
	//json数据
	jsonStr := `{"Name":"liyi","Age":18,"Hobby":["抽烟","喝酒"],"sex":true}`
	//将json数据转为字节数据
	jsonbytes := []byte(jsonStr)
	dataMap := make(map[string]interface{})

	err := json.Unmarshal(jsonbytes, &dataMap)
	if err != nil {
		fmt.Println("反序列化失败,err=", err)
	}
	fmt.Println(dataMap)
}

//将json数据转为结构体
func maina5() {
	type Man struct {
		Name  string
		Age   int
		Hobby []string
		Sex   bool
	}
	//json数据
	jsonStr := `{"Name":"liyi","Age":18,"Hobby":["抽烟","喝酒"],"sex":true}`
	//将json数据转为字节数据
	jsonbytes := []byte(jsonStr)
	//创建一个结构体
	Man1 := new(Man)
	err := json.Unmarshal(jsonbytes, &Man1)
	if err != nil {
		fmt.Println("反序列化失败,err=", err)
	}
	fmt.Println(*Man1)
}

//将json数据转换为map切片
func maina6() {
	jsonStr := `[{"age":21,"hobby":["打炮"],"name":"liyi","sex":true},{"name":"linging"},{"age":18,"hobby":["学习"],"name":"yiyi","sex":true}]`
	jsonBytes := []byte(jsonStr)
	dataSlice := make([]map[string]interface{}, 0)
	err := json.Unmarshal(jsonBytes, &dataSlice)
	if err != nil {
		fmt.Println("反序列化失败,err=", err)
	}
	fmt.Println(dataSlice)
}

//将json数据转为结构体切片
func maina7() {
	type Man struct {
		Name  string
		Age   int
		Hobby []string
		Sex   bool
	}
	jsonStr := `[{"age":21,"hobby":["打炮"],"name":"liyi","sex":true},{"name":"linging"},{"age":18,"hobby":["学习"],"name":"yiyi","sex":true}]`
	jsonBytes := []byte(jsonStr)
	Mans := make([]Man, 0)
	err := json.Unmarshal(jsonBytes, &Mans)
	if err != nil {
		fmt.Println("反序列化失败,err=", err)
	}
	fmt.Println(Mans)
}
