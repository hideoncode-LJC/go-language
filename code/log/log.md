# log

## 使用logger

```go
func Print(v ...any)
func Printf(format string, v ...any)
func Println(...any)

func Fatal(v ...any)
func Fatalf(format string, v ...any)
func Fatalln(... any)

func Panic(v ...any)
func Panicf(format string ,v ...any)
func Panicln(v ...any)
```
+ Print系列函数直接输出日志。
+ Fatal系列函数先输入日志，再调用 os.Exit(1) 终结程序。
+ Panic系列函数先输出日志，再调用 Panic 函数处理错误。

## 配置logger
```go
func Flags() int 
func SetFlags(flag int)
```
+ Flags()函数返回标准logger的输出配置。
+ SetFlags()用来设置标准logger的输出配置。

### Flag 标准配置

```go
const (
    // 控制输出日志信息的细节，不能控制输出的顺序和格式。
    // 输出的日志在每一项后会有一个冒号分隔：例如2009/01/23 01:23:23.123123 /a/b/c/d.go:23: message
    Ldate         = 1 << iota     // 日期：2009/01/23
    Ltime                         // 时间：01:23:23
    Lmicroseconds                 // 微秒级别的时间：01:23:23.123123（用于增强Ltime位）
    Llongfile                     // 文件全路径名+行号： /a/b/c/d.go:23
    Lshortfile                    // 文件名+行号：d.go:23（会覆盖掉Llongfile）
    LUTC                          // 使用UTC时间
    LstdFlags     = Ldate | Ltime // 标准logger的初始值
)
```

### 配置日志前缀

```go
func Prefix() string 
func SetPrefix(prefix string)
```

+ Prefix 查看 logger 的输出前缀
+ SetPrefix 用来设置 logger 的输出前缀

> 这样我们就能够在代码中为我们的日志信息添加指定的前缀，方便之后对日志信息进行检索和处理。

### 配置日志输出位置

```go
func SetOutput(w io.Writer)
```

+ SetOutput函数用来设置标准logger的输出目的地，默认是标准错误输出。

### 创建新的logger

```go
func New(out io.Writer, prefix string, flag int) *Logger
```

+ log标准库中还提供了一个创建新logger对象的构造函数–New，支持我们创建自己的logger示例。
+ New创建一个Logger对象。其中，参数out设置日志信息写入的目的地。参数prefix会添加到生成的每一条日志前面。参数flag定义日志的属性（时间、文件等等）。