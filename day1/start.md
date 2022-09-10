## for 循环

```go
for initialization; condition; post {
    // zero or more statements
}
// a traditional "while" loop
for condition {
    // ...
}
// a traditional infinite loop
for {
    // ...
}
```

for 循环的另一种形式，在某种数据类型的区间（range）上遍历，如字符串或切片。echo 的第二版本展示了这种形式：

类似 for of 的感觉

```go
// Echo2 prints its command-line arguments.
package main

import (
    "fmt"
    "os"
)

func main() {
    s, sep := "", ""
    for _, arg := range os.Args[1:] {
        s += sep + arg
        sep = " "
    }
    fmt.Println(s)
}
```
