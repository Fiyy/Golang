

## Go命令

- go build 可以跟后缀文件名，否则默认文件下所有go文件

  - 会主动忽略掉以_或者.开头的go文件
  - 将文件名后加_系统名，会只在该系统中运行，如xxx_linux只在Linux系统中运行
  - -o 指定生成文件的名字
  - `-p n` 指定可以并行可运行的编译数目，默认是CPU数目

- go get 可以直接获取远程包，并且将其install在GOPATH/pkg中，代码保存在GOPATH/src中，相当于clone + install

  - `-d` 只下载不安装
  - `-f` 只有在你包含了`-u`参数的时候才有效，不让`-u`去验证import中的每一个都已经获取了，这对于本地fork的包特别有用
  - `-fix` 在获取源码之后先运行fix，然后再去做其他的事情
  - `-t` 同时也下载需要为运行测试所需要的包
  - `-u` 强制使用网络去更新包和它的依赖包
  - `-v` 显示执行的命令

- go clean 移除当前源码包和相关源码包中编译生成的文件，文件包括：

  - `-i` 清除关联的安装的包和可运行文件，也就是通过go install安装的文件
  - `-n` 把需要执行的清除命令打印出来，但是不执行，这样就可以很容易的知道底层是如何运行的
  - `-r` 循环的清除在import中引入的包
  - `-x` 打印出来执行的详细命令，其实就是`-n`打印的执行版本

  ```
  _obj/            旧的object目录，由Makefiles遗留
  _test/           旧的test目录，由Makefiles遗留
  _testmain.go     旧的gotest文件，由Makefiles遗留
  test.out         旧的test记录，由Makefiles遗留
  build.out        旧的test记录，由Makefiles遗留
  *.[568ao]        object文件，由Makefiles遗留
  
  DIR(.exe)        由go build产生
  DIR.test(.exe)   由go test -c产生
  MAINFILE(.exe)   由go build MAINFILE.go产生
  *.so             由 SWIG 产生
  ```

- go fmt xxx.go 自动将代码格式化为标准格式

  - `gofmt -w -l src`可以格式化整个项目
  - `-l` 显示那些需要格式化的文件
  - `-w` 把改写后的内容直接写入到文件中，而不是作为结果打印到标准输出。
  - `-r` 添加形如“a[b:len(a)] -> a[b:]”的重写规则，方便我们做批量替换
  - `-s` 简化文件中的代码
  - `-d` 显示格式化前后的diff而不是写入文件，默认是false
  - `-e` 打印所有的语法错误到标准输出。如果不使用此标记，则只会打印不同行的前10个错误。
  - `-cpuprofile` 支持调试模式，写入相应的cpufile到指定的文件

- go install 实际操作分为两步

  - 第一步是生成可执行文件或者.a包
  - 第二步将生成的包移到GOPATH/bin或者GOPATH/pkg下
  - 编译参数 `-v` 可以随时查看底层的执行信息

- go test 将会自动读取目录下名为`xxx_test.go` 的可执行文件，会自动将源代码下的所有test文件都进行测试

  - `-bench regexp` 执行相应的benchmarks，例如 `-bench=.`
  - `-cover` 开启测试覆盖率
  - `-run regexp` 只运行regexp匹配的函数，例如 `-run=Array` 那么就执行包含有Array开头的函数
  - `-v` 显示测试的详细命令

- go tool 里有几个工具命令

  - `go tool fix . `修复老版本的代码到新版本，例如将go1.13的代码转换为go1.14
  - `go tool vet directory|files ` 用来分析当前目录的代码是否都是正确的代码

- go generate 通过分析源码中特殊的注释，执行相应的命令。这个命令是给写代码的人用的，不是给用这个包的人用的，是用来方便你生成一些代码的

  - 例如我们使用yacc来生成代码，会用到这样的命令

    ```
    go tool yacc -o gopher.go -p parser gopher.y
    ```

    -o 指定了输出的文件名， -p指定了package的名称，这是一个单独的命令，如果我们想让`go generate`来触发这个命令，那么就可以在当前目录的任意一个`xxx.go`文件里面的任意位置增加一行如下的注释：

    ```
    //go:generate go tool yacc -o gopher.go -p parser gopher.y
    ```

    这里我们注意了，`//go:generate`是没有任何空格的，这其实就是一个固定的格式，在扫描源码文件的时候就是根据这个来判断的。所以我们可以通过如下的命令来生成，编译，测试。如果`gopher.y`文件有修改，那么就重新执行`go generate`重新生成文件就好。

- go doc 生成文档并查看

  如何查看相应package的文档呢？ 例如builtin包，那么执行`godoc builtin` 如果是http包，那么执行`godoc net/http` 查看某一个包里面的函数，那么执行`godoc fmt Printf` 也可以查看相应的代码，执行`godoc -src fmt Printf`

  通过命令在命令行执行 godoc -http=:端口号 比如`godoc -http=:8080`。然后在浏览器中打开`127.0.0.1:8080`，你将会看到一个golang.org的本地copy版本，通过它你可以查询pkg文档等其它内容。如果你设置了GOPATH，在pkg分类下，不但会列出标准包的文档，还会列出你本地`GOPATH`中所有项目的相关文档，这对于经常被墙的用户来说是一个不错的选择。

- go version 查看go当前的版本
- go env 查看当前go的环境变量
- go list 列出当前全部安装的package
- go run 编译并运行Go程序