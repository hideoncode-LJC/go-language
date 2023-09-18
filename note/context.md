# context

## 多线程之间的通信

**多线程访问临界资源**

多线程访问临界资源会产生问题。

```go
package main

import "fmt"

func main() {

	for k := 0; k < 200; k++ {
		num := 0
		for i := 0; i < 100; i++ {
			go func() {
				num++
			}()
		}
		fmt.Println("k = ", num)
	}
}

```


**传统的访问临界资源的问题**

其他语言多使用共享内存进行数据交换，通过对互斥量进行加锁、解锁。

```go
package main

import "fmt"

var mu sync.Mutex

func main() {

	for k := 0; k < 200; k++ {
		num := 0
		for i := 0; i < 100; i++ {
			go func() {
				mu.Lock()
				num++
				mu.UnLock()
			}()
		}
		fmt.Println("k = ", num)
	}
}
```

**Go语言访问临界资源**

Go语言的并发模型是CSP（Communicating Sequential Processes），提倡通过通信共享内存而不是通过共享内存而实现通信。

如果说goroutine是Go程序并发的执行体，channel就是它们之间的连接。channel是可以让一个goroutine发送特定值到另一个goroutine的通信机制。

Go 语言中的通道（channel）是一种特殊的类型。通道像一个传送带或者队列，总是遵循先入先出（First In First Out）的规则，保证收发数据的顺序。每一个通道都是一个具体类型的导管，也就是声明channel的时候需要为其指定元素类型。



## 为什么要需要context


**基本示例**

```go
package main

import (
	"fmt"
	"time"
	"sync"
)

var wg sync.WaitGroup

func doTask(n int) {
	time.Sleep(time.Duration(n))
	fmt.Printf("Task %d Done\n", n)
	wg.Done()
}

func main() {
	for i := 0 ; i < 3 ; i++ {
		wg.Add(1)
		go doTask(i)
	}

	wg.Wait()
	fmt.Println("All Task Done")
}
```

> 如何实现子协程退出


**全局通道方式**

```go

```

## Context 接口

```go
type Context interface {
	Deadline() (deadline time.Time, ok bool)
	Done() <-chan struct{}
	Err() error
	Value(key interface{}) interface{}
}
```

+  返回`context.Context`被取消的时间，也就是完成工作的截止日期；