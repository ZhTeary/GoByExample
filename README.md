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

    Go string 是只读的bytes数组的切片, `srting == []byte` 

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

# Log
| 日期         | 内容                                                                                                              | 收获           | Key                |
|------------|-----------------------------------------------------------------------------------------------------------------|--------------|--------------------|
| 2023-09-04 | "helloworld, Values, Variables, Constants"                                                                      | harvest 1,2  |                    |
| 2023-09-06 | "For, If/Else, Switch, Arrays, Slices"                                                                          | harvest 3,4  |                    |
| 2023-09-07 | "Maps, Range, Functions, Multiple Return Values, Variadic Functions"                                            | harvest 5,6  | Closure            |
| 2023-09-09 | "Closure, Recursion, Pointers, Strings and Runes, Structs, Methods, Interfaces, Struct Embedding,Generics"      | harvest 7,8  | Strings & Runes    |                                                                   
| 2023-09-10 | "Errors, Goroutines, Channels, Channel Buffering, ChannelSynchronization, Channel Directions, Select, Timeouts" | harvest 9,10 | Timeouts           |
|2023-09-11 | "Timeouts, Non-blocking channel operation, Close channel, channel sychronization"                               | harvest 11,12 | unbuffered channel | 
