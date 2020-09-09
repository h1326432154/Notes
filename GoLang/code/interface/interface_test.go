package interfacetest

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

type Pet interface {
	// SetName(name string)
	Name() string
	Category() string
}

type Dog struct {
	name string // 名字。
}

func (dog *Dog) SetName(name string) {
	dog.name = name
}

func (dog Dog) Name() string {
	return dog.name
}

func (dog Dog) Category() string {
	return "dog"
}

func TestInterface(t *testing.T) {

	// 示例1。
	dog := Dog{"little pig"}

	_, ok := interface{}(dog).(Pet)

	t.Logf("Dog implements interface Pet: %v\n", ok)

	_, ok = interface{}(&dog).(Pet)
	t.Logf("*Dog implements interface Pet: %v\n", ok)
	fmt.Println()

	// 示例2。
	var pet Pet = &dog
	t.Logf("This pet is a %s, the name is %q.\n",
		pet.Category(), pet.Name())

}

func TestInterFace1(t *testing.T) {
	jsonStr := []byte(`{"age":1,"zz":333}`)
	fmt.Println("jsonStr=", jsonStr)
	fmt.Println("jsonStr=", string(jsonStr))
	var value map[string]interface{}
	json.Unmarshal(jsonStr, &value)
	fmt.Println("value=", value)
	age := value["age"]
	fmt.Println("age=", age)
	fmt.Println(reflect.TypeOf(age))
}
