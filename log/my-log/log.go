package log

import (
	"log"
	"os"
)

// log levels
const (
	InfoLevel = iota
	ErrorLevel
	Disabled
)

// log methods
// var (
// 	Error  = errorLog.Println
// 	Errorf = errorLog.Printf
// 	Info   = infoLog.Println
// 	Infof  = infoLog.Printf
// )

func CreateMyLog() {

	logFlag := log.Lshortfile | log.Lmicroseconds | log.Ldate

	openFileFlag := os.O_CREATE|os.O_WRONLY|os.O_APPEND

	errFile, err := os.OpenFile("./errlog.log", openFileFlag, 0644)

	if err != nil {
		log.Println(err)
	}

	infoFile, err := os.OpenFile("./infolog.log", openFileFlag, 0644)
	
	if err != nil {
		log.Println(err)
	}

	errPrefix := "\033[31m[error]\033[0m "
	infoPrefix := "\033[34m[info ]\033[0m "
	
	errorLog := log.New(errFile, errPrefix, logFlag)
	infoLog := log.New(infoFile, infoPrefix, logFlag)

	loggers := []*log.Logger{errorLog, infoLog}



}



// const (
// 	Ldate = 1 << iota
// 	Ltime
// 	Lmicroseconds
// 	Llongfile
// 	Lshortfile
// 	LUTC
// 	Lmsgprefix
// 	LstdFlags = Ldate | Ltime
// )

// func PrintLogConst() {
// 	fmt.Println("Ldate = ", Ldate)
// 	fmt.Println("Ltime = ", Ltime)
// 	fmt.Println("Lmicroseconds = ", Lmicroseconds)
// 	fmt.Println("Llongfile = ", Llongfile)
// 	fmt.Println("Lshortfile = ", Lshortfile)
// 	fmt.Println("LUTC = ", LUTC)
// 	fmt.Println("Lmsgprefix = ", Lmsgprefix)
// 	fmt.Println("LstdFlags = ", LstdFlags)
// }

// type Logger struct {

// }


// func OutPut() {
// 	v := "测试 Fatal"

// 	log.Fatal(v)
// 	log.Fatalf("%s\n", v)
// 	log.Fatalln(v)
// }