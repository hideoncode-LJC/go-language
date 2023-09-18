package main

import (
	// "fmt"
	"log"
	"os"
)


func main() {

	// log.SetFlags(log.Lshortfile | log.Lmicroseconds | log.Ldate)

	// str := "普通日志"

	// log.Println(str)

	// fmt.Println("log.Flags() = ", log.Flags())

	// prefix := log.Prefix()

	// fmt.Println("prefix = ", prefix)


	// log.SetPrefix("\033[31m[error]\033[0m ")

	// log.Println(str)

	// prefix = log.Prefix()

	// fmt.Println("prefix = ", prefix)

	// log.SetPrefix("\033[34m[info ]\033[0m ")

	// log.Println(str)

	// prefix = log.Prefix()

	// fmt.Println("prefix = ", prefix)

	OutPutLog()
}

func OutPutLog() {
	logfile, err := os.OpenFile("./standard-log.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		log.Panicln(err)
	}

	log.SetOutput(logfile)

	log.SetFlags(log.Lshortfile | log.Lmicroseconds | log.Ldate)

	log.SetPrefix("\033[34m[info ]\033[0m ")


	for i := 0 ; i < 10 ; i++ {
		log.Println(i, "输出日志")
	}
}