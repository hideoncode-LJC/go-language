package channel

import (
	"fmt"
)


func DefineChannel() {

	var chInt chan int

	if chInt == nil {
		fmt.Println("该Channel尚未被初始化")
	}

	chInt = make(chan int, 1)

	if chInt != nil {
		fmt.Println("该Channel已经被初始化")
	}

}

func CompareChannel() {
	chInt1 := make(chan int, 1)
	
	var chInt2 chan int

	if chInt1 == chInt2 {
		fmt.Println("chInt1 == chInt2")
	}

	chInt2 = make(chan int, 1)

	if chInt1 == chInt2 {
		fmt.Println("chInt1 == chInt2")
	}

}