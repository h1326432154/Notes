# Runtime

## Goroutine 原理

### Goroutine

> Goroutine 是一个与其他 goroutines 并行运行在同一地址空间的 Go 函数或方法。一个运行的程序由一个或更多个 goroutine 组成。它与线程、协程、进程等不同。它是一个 goroutine   —— Rob Pike

Goroutines 在同一个用户地址空间里并行独立执行 functions，channels 则用于 goroutines 间的通信和同步访问控制。

* goroutine 和 thread 的区别？

1. 内存占用: 创建一个 goroutine 的栈内存消耗为 2 KB(Linux AMD64 Go v1.4后)，运行过程中，如果栈空间不够用，会自动进行扩容。

2. 创建/销毁: 线程创建和销毀都会有巨大的消耗，是内核级的交互(trap)

3. 调度切换: 抛开陷入内核，线程切换会消耗 1000-1500 纳秒(上下文保存成本高，较多寄存器，公平性，复杂时间计算统计)，一个纳秒平均可以执行 12-18 条指令。

4. 复杂性: 线程的创建和退出复杂，多个thread间通讯复杂(share memory)。

### GMP 调度模型

* G:goruntine 每次go func()都代表一个G,无限制
* M:工作线程(OS thread)也被称为 Machine，使用 struct runtime.m，所有 M 是有线程栈的。

    如果不对该线程栈提供内存的话，系统会给该线程栈提供内存(不同操作系统提供的线程栈大小不同)。当指定了线程栈，则 M.stack→G.stack，M 的 PC 寄存器指向 G 提供的函数，然后去执行。

* P:“Processor”是一个抽象的概念，并不是真正的物理 CPU。

    它代表了 M 所需的上下文环境，也是处理用户级代码逻辑的处理器。它负责衔接 M 和 G 的调度上下文，将等待执行的 G 与 M 对接。当 P 有任务时需要创建或者唤醒一个 M 来执行它队列里的任务。所以 P/M 需要进行绑定，构成一个执行单元。
    P 决定了并行任务的数量，可通过 runtime.GOMAXPROCS 来设定。在 Go1.5 之后GOMAXPROCS 被默认设置可用的核数，而之前则默认为1。

* M:N 模型

Go 创建 M 个线程(CPU 执行调度的单元，内核的 task_struct)，之后创建的 N 个 goroutine 都会依附在这 M 个线程上执行，即 M:N 模型。

同一个时刻，一个线程只能跑一个 goroutine。当 goroutine 发生阻塞(chan 阻塞、mutex、syscall 等等)时，Go 会把当前的 goroutine 调度走，让其他 goroutine 来继续执行，而不是让线程阻塞休眠，尽可能多的分发任务出去，让 CPU 忙。

### Work-stealing 调度算法

## 内存分配原理

## GC 原理

## Channel 原理
