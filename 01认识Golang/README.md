# 认识golang
- go语言是什么
- 编译器和集成开发环境的安装
- hello world程序

# Go语言是什么

[golang官方网站](https://golang.google.cn/)

我们想要一个**安全的、静态编译的、高性能的**、类似C++和Java这样的语言，但是得更轻量级并且要像Python这种动态解释型语言这样有趣。 —— Rob Pike（Unix作者、UTF8设计者、Golang作者...）

Go 编程语言是一个**开源**项目，旨在提高程序员的工作效率。

Go 语言**表现力强、简洁、干净、高效**。它的并发机制使编写程序变得容易，从而**最大限度地利用多核**，而其新颖的类型系统则实现了灵活的模块化程序构建。

Go 可以**快速编译**成机器代码，同时还具有**垃圾回收**的便利性和运行时反射的强大功能。它是一种快速、静态类型的编译语言，感觉就像一种动态类型的解释型语言。

—— Go官方文档



# 安装Go编译器和集成开发环境

## 编译器下载和配置

下载地址：[download page](https://golang.google.cn/dl/)

配置环境
```shell
# 国内的需要设置 go 代理
go env -w GOPROXY=https://goproxy.cn,direct
go env -w GO111MODULE=on
```

## 开发环境安装

选择自己喜欢的集成开发环境，随便一个就好。

### Goland

商业软件，付费的软件， 学生和教育工作者及开源软件作者可以申请免费许可，或一些其他的方法。

下载地址 [download page](https://www.jetbrains.com/zh-cn/go/)

### VSCode
开源软件，免费

需要安装一些插件，来支持Golang语言，可以编写多种语言。

下载地址 [download page](https://code.visualstudio.com/Download)

### 其他
NeoVim、Vim、sublime、Atom等等。

# Hello World
> 包、函数以及main函数
```go
package main

import "fmt"

func main() {
	fmt.Println("Hello world!")
}
```

# Run&Build
```shell
# 快速运行项目
go run main.go

# 构建项目
go build main.go
```

# Makefile
使用 makefile 工具来管理项目
