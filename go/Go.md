# GO

## GO函数 一等公民

### 与其他语言的差异

1. 可以有多个返回值
2. 所有参数都是值传递
3. 函数可以作为变量的值
4. 函数可以作为参数和返回值
[text](./go函数差异.png)

#### defer函数 延迟执行函数

## 接口interface

### 与其他编程语言的差异

1. 接口为非入侵性，实现不依赖于接口定义
2. 所以接口的定义可以包含在接口使用者包内

## 错误机制

### panic

* 用于不可恢复的错误
* panic退出前会执行defer指定的内容

### exit

* 退出前不会执行defer指定的内容

### recover

* 错误恢复

## 构建可复用模块

### init方法

* 在main被执行之前，所有依赖的package的init方法都被别执行
* 不同包的init函数按照包导入顺序的依赖关系决定执行顺序
* 每个包可以有多个init
* 包的每个源文件也可以有多个init函数

### Go并发编程

#### 协程使用

    * 例：

    var wg sync.WaitGroup // wg
    for i := 1; i <= 3; i++ {
        wg.Add(1)   // 新增一个协程
        go func(i int) {    // 协程标识 go func
            fmt.Printf("%d", i)
            wg.Done()   // 主体执行完，关闭协程wg.Done() 等同于 wg.Add(-1)
        }(i)
    }
    wg.Wait()   // 协程执行中，等待
