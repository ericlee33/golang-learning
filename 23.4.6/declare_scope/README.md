# golang 可以跨文件共享函数 / 变量

一个 Go 语言编写的程序对应一个或多个以.go 为文件后缀名的源文件。每个源文件中以包的声明语句开始，说明该源文件是属于哪个包。包声明语句之后是 import 语句导入依赖的其它包，然后是包一级的类型、变量、常量、函数的声明语句，包一级的各种类型的声明语句的顺序无关紧要（译注：函数内部的名字则必须先声明之后才能使用）。例如，下面的例子中声明了一个常量、一个函数和两个变量：

# 操作

```shell
1.
go run *.go

2.
go build
./declare_scope
```
