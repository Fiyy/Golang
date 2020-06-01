# Go语言学习

## The tour of Go 笔记

- := 在函数外不可用
- go的函数内局部变量必须使用，全局变量可以不使用
- 常量不能用：=声明
- go的函数可以返回多个返回值，声明函数时可以声明返回值的名字，自动返回对应值，应该多用于多个返回值的情况
- for语句不需要括号，for语句的初始化部分和循环部分都可以省略，其分号也可以省略
- if语句不需要括号，if语句也可以增加初始化部分，初始化部分的变量可以在else部分中使用
- import的包必须要使用
- case中得值不一定必须是整数，任何类型都可以使用
- switch语句块的条件可以为空，这种情况下默认为条件真匹配，当case中的值为真时匹配
- go的指针没有算法，不能对指针进行加减计算
- 切片的零值是nil
- 在使用方法的时候，将数据结构当作接收器不能改变结构中变量值，但如果使用指针就可以改变结构中的变量值
- 接口使用时候必须要准确匹配，是指针就是指针，是结构就是结构

## Go语言基础

[github Go学习]: https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/preface.md

### 2.0 Go语言总述

- Go语言的关键字很少，Go语言的编译速度是很快的，关键字共有25个：

```go
break    default      func    interface    select
case     defer        go      map          struct
chan     else         goto    package      switch
const    fallthrough  if      range        type
continue for          import  return       var
```

- Go语言使用UTF-8，发明者相同。

### 2.1 Go中的Package

- Go程序是通过package来组织运行的
- 每个程序的第一行都是 `package <pkgname>` 告诉我们当前文件属于哪个包
- 每个程序想要运行必须要有main包，除了main包以外的每个包都能通过install来生成*.a文件，就是包文件，放在GOPATH/pkg中，GOPATH是自己设置的，而main包会生成对应的一个可执行程序。
- 包的概念有助于将程序模块化，可以将一个程序分为很多个模块，提高可重用性。
- Go中的main函数没有参数也没有返回值。
- 包和包所在的文件夹名可以是不同的。

### 2.2 Go基础

#### 变量

- 用var关键字可以声明变量，变量类型在变量名后面，可以var一次声明多个变量，如果变量类型相同只需要在最后一个变量说明类型
- 在用var声明变量(和常量不同)时可以直接进行初始化

```go
var vname1, vname2, vname3 type= v1, v2, v3
```

​	也可以直接省略类型，将根据所初始化的值对变量进行类型归类

```go
var vname1, vname2, vname3 = v1, v2, v3
```

​	还可以使用`:=`符号来取代 `var` 和 `type`，但是必须要给初始值，而且该语句只能用在函数内部

​	因此`var`多用来声明全局变量

- `_` 可以用来接收值，但是相当于丢弃
- Go语言中声明的变量必须使用，否则编译报错
- Go语言中，如果没有初始化，不管在什么位置声明的，其默认值都是nil或0或者false

#### 常量

- 常量用 `const` 来声明

```go
const constantName = value
//如果需要，也可以明确指定常量的类型：
const Pi float32 = 3.1415926
```

- Go中的常量可以指定相当多的小数位数（例如200位），如果指定给 `float32` 就自动缩短为32bit，同样如果是 `float64`就会自动缩短为64位

#### 内置基础类型

- bool类，true和false，默认值是false
- 整数类型有 `int`、`uint`、`rune`, `int8`, `int16`, `int32`, `int64`和`byte`, `uint8`, `uint16`, `uint32`, `uint64`。其中`rune`是`int32`的别称，`byte`是`uint8`的别称。
- **注意注意，不同类型之间不能相互赋值，不然会报错！**
- 浮点数有 `float32` 和`float64`，默认是`float64`
- 复数类型，默认是`complex128`，虚部和实部各自64位，还有`complex64`

```go
var c complex64 = 5+5i
//output: (5+5i)
fmt.Printf("Value is: %v", c)
```

- 字符串类型string，Go中都使用`UTF-8`,可以用双引号或者反引号（我的键盘是esc下面的1左边的按键）来括起来。字符串的底层物理数据结构还是字节数组，在赋值的时候其实只修改了数据首地址和长度。
- 字符串的值是不能改变的，如果想要改变，首先将字符串转换为[]byte的数组，然后用数组的方式进行下标选择修改，然后在转换为string

```go
s := "hello"
c := []byte(s)  // 将字符串 s 转换为 []byte 类型
c[0] = 'c'
s2 := string(c)  // 再转换回 string 类型
fmt.Printf("%s\n", s2)
```

​	还有一种修改字符串的方式，用切片来操作（不知道为啥感觉有点没有道理，这不就还是强行改变啦字符串的内容？？？

```go
s := "hello"
s = "c" + s[1:] // 字符串虽不能更改，但可进行切片操作
fmt.Printf("%s\n", s)
```

![img](https://chai2010.gitbooks.io/advanced-go-programming-book/content/images/ch1-7-array-4int.ditaa.png)