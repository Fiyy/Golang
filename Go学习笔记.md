# Go学习笔记

[TOC]

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

该学习笔记基于[教程](https://github.com/astaxie/build-web-application-with-golang) ，推荐所有go语言学习者学习。

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
- 字符串的值是不能改变的，如果想要改变，首先将字符串转换为[]byte的数组，然后用数组的方式进行下标选择修改，然后在转换为string。这里的修改其实是指变量名指向的字符串的修改，而不是字符串变量名称对应的字符串不能改变，你完全可以使用赋值语句重新赋值，这一点和Java中的String相同。

```go
s := "hello"
// 这里可以使用 s = "asdf"来修改字符串
c := []byte(s)  // 将字符串 s 转换为 []byte 类型
c[0] = 'c'
s2 := string(c)  // 再转换回 string 类型
fmt.Printf("%s\n", s2)
```

​	还有一种修改字符串的方式，用切片来操作（其实也就是对这个字符串进行一个重新的赋值，字符串变量名将会指向另一个地址而不是最初的地址）

```go
s := "hello"
s = "c" + s[1:] // 字符串虽不能更改，但可进行切片操作
fmt.Printf("%s\n", s)
```

- 可以使用 `+` 来连接两个字符串
- 如果要声明一个多行的字符串，可以用反引号来声明，反引号括起来的内容是**Raw字符串**，即在代码中的样子就是实际打印的形式，没有字符转义

#### 错误类型

- Go中内置来`error` 类型，可处理错误信息，可以通过`errors` 包中的函数来处理错误，编辑错误信息

```go
err := errors.New("emit macho dwarf: elf header corrupted")
if err != nil {
	fmt.Print(err)
}
```

#### iota 枚举

在使用声明`enum` 时采用，默认开始值为0，好像只能被常量使用？？

```go
const (
	x = iota // x == 0
	y = iota // y == 1
	z = iota // z == 2
	w        // 常量声明省略值时，默认和之前一个值的字面相同。这里隐式地说w = iota，因此w == 3。其实上面y和z可同样不用"= iota"
)

const v = iota // 每遇到一个const关键字，iota就会重置，此时v == 0

const (
	h, i, j = iota, iota, iota //h=0,i=0,j=0 iota在同一行值相同
)

const (
	a       = iota //a=0
	b       = "B"
	c       = iota             //c=2
	d, e, f = iota, iota, iota //d=3,e=3,f=3
	g       = iota             //g = 4
)
```

可以看出两点：

- 每次遇到const关键字时，iota的值归零
- 每一次换行都会增加1，相同一行的iota值是相等的

#### Go的私有和公有

- 大写字母开始的变量是可以导出的，可以被其他包使用，为公有变量，小写字母开始的变量是私有变量
- 同理对函数也是这样，大写字母开始的函数时可以在其他包中使用的

#### 基本的数据结构

##### 数组array

数组的定义方式主要是两种，**都要规定数组的长度，而且长度不再可变**：

```go
var arr [n]type
arr := [10]type{} // 注意这里必须要有中括号，内容可以为空，也可以包含初始化内容
// 如果省去数组初始化长度，就变成了slice，数组是必须有规定长度的
arr1 := []type{1,2,3} // slice
```

- 数组也可以像切片一样省略数组长度，但是括号内`[...]`的方式，然后通过数组内容自动计算数组的长度。

可以用下标对数组进行读取和赋值的功能。

- **注意：**长度也是数组的内容，一个数组的长度不能改变，数组之间的赋值是值得赋值，而不是引用的赋值，例如：**在将一个数组作为参数传入函数的时候，传入的其实是数组的副本，而不是指针或者引用，这一点和C语言相同**， 如果要使用指针的传递，就要使用切片`slice`类型。

数组还支持多维数组，例如：

```go
// 声明了一个二维数组，该数组以两个数组作为元素，其中每个数组中又有4个int类型的元素
doubleArray := [2][4]int{[4]int{1, 2, 3, 4}, [4]int{5, 6, 7, 8}}

// 上面的声明可以简化，直接忽略内部的类型
easyArray := [2][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}
```

##### 切片slice

切片实际上是一个引用类型，一个`slice` 总是指向一个底层数组， 其声明方式和数组的声明基本相同，只是少了数组的长度，切片可以先不声明大小，也不用在声明时候初始化内容，可以在之后使用的过程中再进行赋值

```
arr := [4]int{1, 2, 3}
var f []int
f = arr[0:2]
f[1] = 3
arr[0] = 1
```

切片可以通过指向一个数组的片段来进行赋值，第一个数字是起始的位置，第二个数字是结束位置但是不包含数组元素arr[j]，所以slice指向的数组长度还是 j - i

- 切片的默认开始位置是0，可以这么截取一个数组 `a[:n]` ，等价于`a[0:n]`
- 切片在截取数组时候第二个数字默认值为数组长度，所以`a[x:]`等价于 `a[x:n]`这里的n是数组的长度
- 由前两条可以得出，`a[:]`其实就是对整个数组的引用
- **注意：slice的截取时候不能超出数组的可取范围，否则会报错**

由于slice的实质是引用，所以切片中修改元素的值时，其本质的物理存储数组中也会同时发生改变，数组和其他相同区域的切片中的数组元素的值也会进行修改。

我们也可以将slice看作时一个另类的数据结构，其中包含三个元素:：

- 指针：指向数组中slice开始的位置
- 长度len：slice的长度，即为j-i
- 最大长度cap：slice开始的位置到数组最后的位置之间的长度，是由数组的长度减去i计算得到

如果将slice看作一个结构，我们有对应的几个内置函数可以使用：

- len：切片的长度
- cap：切片的容量
- append：向slice中添加元素，返回一个相同类型的slice，会改变基础数组中的内容，但是如果添加内容过多而slice中的最大长度不足，就会重新动态申请一个存储空间，向新的数组复制原数组的内容，将slice指向这个新的存储空间，这种情况下不会影响最初的基础数组和对应的其他slice
- copy：从源slice的基础数组中复制元素到目标，返回复制元素的个数

现在的slice支持三个参数：

```go
var array [10]int
slice := array[2:4]	// 该切片的容量是10 - 2 = 8
slice = array[2:4:7] // 该切片的容量是7 - 2 = 5
```

前两个参数和之前介绍的相同，第三个参数讲起来可能会稍有点绕：它代指切片的cap(最大长度)所不能指向的第一个元素的位置（如果没有指定第三个参数，其默认值就是基础数组的长度），因此这个切片的容量cap就是7 - 2 = 5，即这个slice不能访问基础数组的最后三个位置。

##### 映射（字典） map

map的声明首先是声明一个变量名，然后必须使用make函数对实际的map进行初始化，这一点和Java很相似，如下：

```go
var numbers map[string]int
numbers = make(map[string]int)
// 也可以将前两句融合为一句 var numbers = make(map[string]int)
numbers["one"] = 1  //赋值
numbers["ten"] = 10 //赋值
numbers["three"] = 3
```

还可以使用另一种方式来声明map，

```go
numbers := make(map[string]int)
```

map也可以不使用make函数进行初始化，但是就必须在声明时使用`｛｝` ，`{}`中可以添加一些键值对`key:val`进行初始化，也可以为空，例如：

```go
var numbers = map[string]int{}
numbers["one"] = 1  //赋值
numbers["ten"] = 10 //赋值
numbers["three"] = 3
fmt.Println(numbers["one"])
```

map 使用注意：

- map中的key值可以为任何定义了`==` 和`!=`操作的类型
- map 是无序的，每次打印的结果都不一样（这里指顺序），只能通过key来获取，不能通过类似数组的方式来用index获取值
- map也有内置的len函数，返回对应的键值对数量
- map的值修改很方便，和数组一样
- map 不是线程安全的，如果有多个runtine存取时必须主动使用mutex lock机制进行控制，其他基本类型不同，它们都是线程安全的。

map的键值对删除可以使用delete函数：

```
delete(numbers "one")  // 删除key为C的元素
```

**map是一种引用类型，当两个map指向来同一个底层map结构时，其中一个map内容修改，另一个map中也会自动修改：**

```go
var numbers = map[string]int{}
numbers["one"] = 1  //赋值
numbers["ten"] = 10 //赋值
numbers["three"] = 3
fmt.Println(numbers["one"]) // 打印结果为1
grade := numbers
grade["one"] = 999
fmt.Println(numbers["one"]) //打印结果为999
```

#### make 和 new 操作

`make`用于内建类型（`map`、`slice` 和`channel`）的内存分配。`new`用于各种类型的内存分配。

- new函数和其他语言中的new函数功能相同，new(T)分配一个用零值填充的T类型的内存空间，返回空间的地址，这里要注意，地址其实就是一个指针值，因此**new函数返回的是一个指针*T，而不是一个类型T的值**。
- make函数只能创建map、slice、channel，返回一个有初始值（非零）的T类型，而不是*T，这一点是和new函数最大的区别。导致这三个类型有所不同的原因就是指向数据结构的引用必须在使用前就被初始化（？？？没有理解，理解之后补充）

#### 零值

下面是对于不同类型的零值表：

```
int     0
int8    0
int32   0
int64   0
uint    0x0
rune    0 //rune的实际类型是 int32
byte    0x0 // byte的实际类型是 uint8
float32 0 //长度为 4 byte
float64 0 //长度为 8 byte
bool    false
string  ""
```

### 2.3 流程和函数

#### 流程控制

- if 语句不用加括号，if语句中可以声明一个变量，该变量只在条件逻辑块内起作用
- goto跳转到标签处，谨慎使用
- for语句，不需要使用括号，可以当做while使用，只有一个条件语句，省去两个`;`

- break 和 continue 语句和C语言中相同
- 可以使用for 配合 range使用来读取slice和map中的数据，返回第一个值为key，第二个为value

```go
for k,v:=range map {
	fmt.Println("map's key:",k)
	fmt.Println("map's val:",v)
}
```

- switch 语句的表达式可以不是整数或者常量，与C语言的switch的区别在于执行一个case后自动跳出，相当于每个case中自带来break语句
- switch语句中可以在case语句中使用`fallthrough`来强行执行后面的case代码，如下：

```go
integer := 6
switch integer {
case 4:
	fmt.Println("The integer was <= 4")
	fallthrough
case 5:
	fmt.Println("The integer was <= 5")
	fallthrough
case 6:
	fmt.Println("The integer was <= 6")
	fallthrough
default:
	fmt.Println("default case")
}
```

- 如果switch没有表达式，则会匹配`true`的情况，类似于条件语句

```go
	x := 5
	switch {
	case x <= 4:
		fmt.Println("The integer was <= 4")
	case x <= 5:
		fmt.Println("The integer was <= 5")
	case x <= 6:
		fmt.Println("The integer was <= 6")
	default:
		fmt.Println("default case")
	}
```

#### 函数

函数是用func来声明的，格式如下：

```go
func funcName(input1 type1, input2 type2) (output1 type1, output2 type2) {
	//这里是处理逻辑代码
	//返回多个值
	return value1, value2
}
```

- 可以有多个传入参数，也可以有多个返回值，都要在参数后带有类型，返回值可以不给变量名只给类型,例如：

```go
func func0(input1, input2 int) (x, y int) {
	x = input1
	y = input2
	return
}
```

- 如果返回值有变量名，可以直接在函数内使用该变量，在返回的时候用`return` 语句可以不说明变量名
- 如果只有一个返回值而且不声明返回值变量，可以直接省略掉括号写法如下：

```go
func func1(input1 int, input2 int) int {
	return input1 + input2
}
```

- 如果几个传入参数的类型一样，只需要在最后一个传入参数处注明类型

```
func func2(input1, input2 int) int {
	return input1 + input2
}
```

#### 可变参数

Go支持可变参数，用 `xxx...type`来构成，代码如下：

```go
func myfunc(arg ...int) {}
for _, n := range arg {
	fmt.Printf("And the number is: %d\n", n)
}
```

#### 函数传参的传值与传指针

- 传值就是传了一个copy值
- 传指针，指针指向存储空间，在函数内直接操作存储空间，在函数中的操作会影响传入参数的值
- go中的指针可以直接使用`.` 来操作结构

**传指针的好处：**

- 可以实际修改对象，操作同一个对象
- 传指针只需要传一个地址值，对于操作对象是一个较大的结构体时，开销较小
- channel、slice、map这三种类型的实现类似指针，因此可以直接传递，实际效果和传递指针相同，**但是如果想要改变slice的长度，还是需要取地址传递指针**

#### defer

defer语句用作延迟效果，如果在一个函数中使用defer语句，会在函数执行到最后时，将这些defer语句逆序执行，然后函数返回。

defer语句作用：当在进行一些资源打开操作时，如果遇到错误，需要提前返回，但是返回前必须要关闭相应的资源，不然会导致资源泄露，所以就需要执行defer语句。

```go
func ReadWrite() bool {
	file.Open("file")
	defer file.Close()
	if failureX {
		return false
	}
	if failureY {
		return false
	}
	return true
}
```

- 注意：defer是在函数返回前**逆序执行**，采用先进后出的栈，这个特性才能释放资源！

#### 函数作为值、类型

go中的函数可以作为一个类型，它的类型就是所有拥有相同的参数（个数和类型相同），相同的返回值（个数和类型相同）的一种类型，但是函数内部不一定相同，那么函数作为类型的作用是什么？就是可以将函数当做值来传递，如下面的代码：

```go
package main
import "fmt"

type testInt func(int) bool // 声明了一个函数类型，类型名字是testInt

func isOdd(integer int) bool { 
	if integer%2 == 0 {
		return false
	}
	return true
}

func isEven(integer int) bool {
	if integer%2 == 0 {
		return true
	}
	return false
}

// 声明的函数类型在这个地方当做了一个参数

func filter(slice []int, f testInt) []int { // 这里的f就是函数，作为值传入该函数
	var result []int
	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}
func main(){
	slice := []int {1, 2, 3, 4, 5, 7}
	fmt.Println("slice = ", slice)
	odd := filter(slice, isOdd)    // 函数当做值来传递了
	fmt.Println("Odd elements of slice are: ", odd)  //打印 1 3 5 7
	even := filter(slice, isEven)  // 函数当做值来传递了
	fmt.Println("Even elements of slice are: ", even) //打印2 4
}
```

那么，我们又为什么要将一个函数作为参数传入另一个函数呢？ 其实主要是用作编写一些通用的函数接口，在可以将函数当做参数传入时，会让编程非常灵活和方便。

此外，也可以不将函数声明成为一种类型再进行参数传递，直接就在函数的传入参数处写该函数的类型，例如将上面的代码可以改为：

```go
func filter(slice []int, f func(int) bool) []int { // 这里的f就是函数，作为值传入该函数
	var result []int
	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}
```

#### Panic和Recover

Go中没有异常机制，不能抛出异常，因此就使用`panic` 和`recover`机制。但是一定要谨慎使用，一般的代码应该不使用或者很少使用。

**panic**

> 是一个内建函数，可以中断原有的控制流程，进入一个`panic`状态中。当函数`F`调用`panic`，函数F的执行被中断，但是`F`中的延迟函数会正常执行，然后F返回到调用它的地方。在调用的地方，`F`的行为就像调用了`panic`。这一过程继续向上，直到发生`panic`的`goroutine`中所有调用的函数返回，此时程序退出。`panic`可以直接调用`panic`产生。也可以由运行时错误产生，例如访问越界的数组。

**Recover**

> 是一个内建的函数，可以让进入`panic`状态的`goroutine`恢复过来。**`recover`仅在延迟函数中有效**。**在正常的执行过程中，调用`recover`会返回`nil`，并且没有其它任何效果**。如果当前的`goroutine`陷入`panic`状态，调用`recover`可以捕获到`panic`的输入值，并且恢复正常的执行。

panic 函数演示如下：

```go
var user = os.Getenv("USER")

func init() {
	if user == "" {
		panic("no value for $USER")
	}
}
```

#### `main`函数和`init`函数

- `init`函数能作用于所有的包，而`main`函数只能用于`main`包中。
- 两个函数都没有参数值和返回值
- 在每个package中可以写任意多个`init`函数，但是建议只写一个
- `init`函数是可选的，main`函数在``main`包中是必须包含的
- go程序会自动调用这两个函数
- 程序的初始化和执行都起始于`main`包，一个包可能在多个包内都进行了`import`，但是实质上只会导入一次供所有包使用
- 一个包被导入的时候，如果该包中还导入了其他包，会首先导入其他包，再对这些包的包级别常量和变量初始化，接着执行`init`函数（如果没有就不执行），等所有被导入的包都加载完了，再开始对`main`包中的包级常量和变量进行初始化，然后执行`main`包中的`init`函数，再执行`main`函数。具体过程如下图：

![img](https://github.com/astaxie/build-web-application-with-golang/raw/master/zh/images/2.3.init.png?raw=true)

#### import

有两种常用模式的import：

1. 相对路径
2. 绝对路径

还有部分是比较特殊的import方式：

1. 点操作

   ```go
    import(
        . "fmt"
    )
   ```

   点操作的作用是在之后调用该包函数是可以省略掉包名，直接用函数，如：

   ```go
   fmt.Println(2)
   Println(2)
   ```

   **注意: 这里要注意一点，在使用点操作引入包后，不能再使用fmt.xxxx函数，而是必须直接使用xxx函数，否则会出现错误，除非你再写一次不带点的导入操作**

2. 别名操作

   类似于import xxx as xxx，可以起一些简单的名字便于记忆

   ```go
    import(
        f "fmt"
    )
   fmt.Println(2)
   f.Println(2)
   ```

   **注意:在使用别名操作引入包后，不能再使用fmt.xxxx函数，而是必须使用x.xxx函数，否则会出现错误，除非你再写一次不带别名的导入操作**

3. `_`操作

   其实就是将一个包引入，但是其包名起别名为`_`，其实也就是不能用了，用了这个操作后，只是引入该包，但是不能主动调用该包内的函数，只是会程序自动调用该包内的`init`函数。

   **注意:同理，不能再使用fmt.xxx了** 

### 2.4 struct类型

`struct`结构可以用`type`声明为类型

```go
type person struct {
	name string
	age int
}
```

struct结构的成员可以用`.`访问和操作，struct对应变量的声明和普通变量相同，但是初始化有好几种方式

```go
var p person
p.name = "asdfad"
p.age = 4
// 还可以这么声明
var p person = person{"asd", 4}
```

如果是在函数内部，还可以使用其他声明方式：

- 按照顺序提供初始值

  ```go
  p := person{"asfd",1}
  ```

- 按照`field:value`初始化

  ```go
  P := person{age:24, name:"Tom"}
  ```

- 