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
   


# Log

2023-09-04 : src/abcd , harvest 1,2

2023-09-06 : src/efghi , harvest 3,4

2023-09-07 : src/jklmno , harvest  5,6 , TODO: closure

2023-09-09 : src/pqrstuvw, harvest 7,8 