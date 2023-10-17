- About this project
    
    Go by Example is a hands-on introduction to Go using annotated example programs

- Time
    
    2023-09-04 ~ 

- Toc

# Harvest
1. Import
    
    import **"Mod/Package"**


2. e

   以**10**为底, 并不是以e为底 5e3 = 5000


3. switch

    对于多种条件的判断更为灵活


4. slice

    make([]type,num) 使用make声明且最好预先分配内存

    添加元素append

    go 1.21以后可以用slices包


5. values in map

    ```
   if _,prs := m["key"]; prs{ 
        return true 
   }else{
        return false
   }
    ```
   

6. 闭包Closure
   
   闭包在返回函数的同时还记录了上下文环境（把里面的变量封闭在了闭包里面）

   > "The returned function closes over the variable i to form a closure." ---GoByExample
   
   封闭了的包，把外层函数里面的变量环境也存储在了二层函数里面，并且将用到的变量封闭在里面
   。
   
   //Todo
   闭包是返回了二层函数，所以一层函数的作用就是提供变量的生命周期？


7. strings

    Go string 是只读的bytes数组的切片, `string == []byte`

   In other languages, strings are made of “characters”. In Go, the concept of a character is called a rune - it’s an integer that represents a Unicode code point.

   > https://go.dev/blog/strings


8. methods
   
   Go supports **methods defined on struct types.**
   
   也就是说方法只能在自定义的type上


9. Select
   
   _select_ lets you **wait on multiple channel operations.**
   
   select 语句只能用于通道操作，每个 case 必须是一个通道操作，要么是发送要么是接收。

   select 语句会监听所有指定的通道上的操作，一旦其中一个通道准备好就会执行相应的代码块。

10. Timeouts
   > _Timeouts_ are important for programs that connect to external resources or that
   otherwisse need to bound execution time. Implementing timeouts in Go is easy and elegant
   thanks to _channels_ and _select_


11. Non-Blocking Channel Operations

   > Basic sends and receives on channels are blocking. 
   > However, we can use **_select_** with a _**default**_ clause to implement 
   > non-blocking sends, receives, and even non-blocking multi-way _**selects**_.

12. 无缓冲通道

   无缓冲通道指在获取数据之前没有能力保存数据的通道，这种类型的通道要求亮哥Goroutine同时处于执行状态才能完成写入和获取

   如果两个Goroutine没有同时准备，某一个Goroutine执行写入或获取操作将会处于阻塞等待状态，另一个Goroutine无法执行写入或获取操作，
   _程序将会提示异常，这种类型的通道执行写入和获取的交互行为是同步，任意一个操作都无法离开另一个操作单独存在_

   当我们使用无缓冲通道的时候，必须注意通道变量的操作，
   确保程序中有两个或两个以上的Goroutine同时执行通道的读写操作，读写操作必须是一读一写，不能只读不写或只写不读

13. go 闭包在循环中

```go
func main() {
    done := make(chan bool)

    values := []string{"a", "b", "c"}
    for _, v := range values {
        go func() {
            fmt.Println(v)
            done <- true
        }()
    }

    // wait for all goroutines to complete before exiting
    for _ = range values {
        <-done
    }
}
```
期望输出：`a, b, c`
实际有可能的输出：`c, c, c`

这是因为循环的每次迭代都用相同的实例变量v，所以每个闭包共享相同的变量。
当闭包运行时，在`fmt.Println`语句执行时输出变量v，
但是v有可能在goroutine启动时被修改。 想要检测这种问题可以用`go vet`

为了绑定每个闭包在启动时的变量v，我们最有可能在循环内为每次迭代新建一个变量，
其其中一种办法时将变量作为参数传进闭包：
```go
for _,v :=range values {
    go func(u string) {
        fmt.Println(u)
        done <- true
    }(v)
}
```

另一种更简单的办法是用声明的方式创建一个变量，有可能看起来奇怪但是有用：
```go
for _,v := range values {
    v:=v //create a new 'v'
    go func() {
        fmt.Println(v)
        done <- true
    }()
}
```

15. defer

defer 用于确保函数调用再程序执行的一会儿后调用，通常用于清除的场景。

16. recover

go通过 `recover`内置函数可以让程序从panic中恢复，停止panic并执行defer闭包

`**recover` 必须在defer函数中调用，**当发生panic时，defer会激活并且recover调用会捕捉panic，此时程序会将defer执行完后再结束


# Log
| 日期         | 内容                                                                                                                                                                                    | 收获           | Key                 |
|------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|--------------|---------------------|
| 2023-09-04 | "helloworld, Values, Variables, Constants"                                                                                                                                            | harvest 1,2  |                     |
| 2023-09-06 | "For, If/Else, Switch, Arrays, Slices"                                                                                                                                                | harvest 3,4  |                     |
| 2023-09-07 | "Maps, Range, Functions, Multiple Return Values, Variadic Functions"                                                                                                                  | harvest 5,6  | Closure             |
| 2023-09-09 | "Closure, Recursion, Pointers, Strings and Runes, Structs, Methods, Interfaces, Struct Embedding,Generics"                                                                            | harvest 7,8  | Strings & Runes     |                                                                   
| 2023-09-10 | "Errors, Goroutines, Channels, Channel Buffering, ChannelSynchronization, Channel Directions, Select, Timeouts"                                                                       | harvest 9,10 | Timeouts            |
|2023-09-11 | "Timeouts, Non-blocking channel operation, Close channel, channel sychronization"                                                                                                     | harvest 11,12 | unbuffered channel  | 
|2023-09-15 | “Range chan, timers, rickers worker pools, waitGroups”                                                                                                                                | harvest 13 | waitGroups          |
|2023-09-27 | "rateLimiting, atomicCounter, mutexes"                                                                                                                                                | |                     |
|2023-10-02 | "stateful goroutines,sort,sore by functions, panic, defer, recover"                                                                                                                   | | stateful goroutines |
|2023-10-17 | "String Functions","String Formatting","Text Templates","Regular Expressions","JSON","XML","Time","Epoch","Time Formatting / Parsing","Random Numbers","Number Parsing","Url parsing" | | Url parsing         |

