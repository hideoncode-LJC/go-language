# context

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