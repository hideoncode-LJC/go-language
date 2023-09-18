# Channel

## 为什么需要Channel

如果说goroutine是Go语言程序的并发体的话，那么channels则是它们之间的通信机制。一个channel是一个通信机制，它可以让一个goroutine通过它给另一个goroutine发送值信息。每个channel都有一个特殊的类型，也就是channels可发送数据的类型。一个可以发送int类型数据的channel一般写为chan int。

```go
// 初始化Channel
package main

import (
	"fmt"
)

func main() {
	var chInt chan int

	if chInt == nil {
		fmt.Println("该Channel尚未被初始化")
	}

    // make(chan type, len)
	chInt = make(chan int, 1)

	if chInt != nil {
		fmt.Println("该Channel已经被初始化")
	}
}

![Alt text](image.png)


// 
```
和map类似，channel也对应一个make创建的底层数据结构的引用。当我们复制一个channel或用于函数参数传递时，我们只是拷贝了一个channel引用，因此调用者和被调用者将引用同一个channel对象。和其它的引用类型一样，channel的零值也是nil。

## 声明Channel

