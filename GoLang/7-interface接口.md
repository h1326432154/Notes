#

## 接口 interface

* 在面向对象的语言中，接口是用来限制实现类行为的。

> GO语言有接口类型（interface{}），它与面向对象的接口含义不同，GO语言的接口类型与数组（array）、切片（slice）、集合（map）、结构体（struct）是同等地位的。

interface{}是一个`任意类型`,或者说是万能类型,定义一个变量为interface{}类型，可以把任意的值赋给这个变量

[示例代码](./code/interface/interface_test.go)

```go
type Pet interface {
  SetName(name string)
  Name() string
  Category() string
}
```

一个接口类型Pet，它包含了 3 个方法定义，方法名称分别为SetName、Name和Category。这 3 个方法共同组成了接口类型Pet的方法集合。

只要一个数据类型的方法集合中有这 3 个方法，那么它就一定是Pet接口的实现类型。这是一种无侵入式的接口实现方式。这种方式还有一个专有名词，叫“Duck typing”，中文常译作“鸭子类型”

* 判定一个数据类型的某一个方法实现的就是某个接口类型中的某个方法

    这有两个充分必要条件，一个是“两个方法的签名需要完全一致”，另一个是“两个方法的名称要一模一样”

### interface类型与struct类型

struct里面的成员是变量，而interface里面的成员是函数，即我们可以使用interface定义接口。
