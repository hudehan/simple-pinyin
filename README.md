# simple-pinyin
适合go简单的拼音库，可以自己配置基数据，也可以使用第三方字典，支持多格式输出


go-pinyin
=========

[![Build Status](https://github.com/mozillazg/go-pinyin/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/mozillazg/go-pinyin/actions/workflows/ci.yml)
[![Coverage Status](https://coveralls.io/repos/mozillazg/go-pinyin/badge.svg?branch=master)](https://coveralls.io/r/mozillazg/go-pinyin?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/mozillazg/go-pinyin)](https://goreportcard.com/report/github.com/mozillazg/go-pinyin)
[![GoDoc](https://godoc.org/github.com/mozillazg/go-pinyin?status.svg)](https://godoc.org/github.com/mozillazg/go-pinyin)

汉语拼音转换工具 Go 版。


Installation
------------

```
go get -u github.com/mozillazg/go-pinyin
```

install CLI tool:

```
go get -u github.com/mozillazg/go-pinyin/cmd/pinyin
$ pinyin 中国人
zhōng guó rén
```


Documentation
--------------

API documentation can be found here:
https://godoc.org/github.com/mozillazg/go-pinyin


Usage
------

```go
package main

import (
	"fmt"
	"github.com/hudehan/simple-pinyin"
)

func main() {
    hans := "长歌行"

    // Normal输出
    a := pinyin.Args{
      Style: pinyin.Normal,
    }
    fmt.Println(pinyin.Pinyin(hans, a))
    // [chang ge xing]

    // Normal输出，包含多音字
    a := pinyin.Args{
      Style:     pinyin.Str,
      Heteronym: false,
    }
    fmt.Println(pinyin.Pinyin(hans, a)) 
    // [changgexing]

    // Normal输出，包含多音字
    a := pinyin.Args{
      Style:     pinyin.Str,
      Heteronym: true,
    }
    fmt.Println(pinyin.Pinyin(hans, a))
    // [changgexing changgehang zhanggexing zhanggehang]


}
```更多待你探索

注意：

* 默认情况下会忽略没有拼音的字符（可以通过自定义 `Fallback` 参数的值来自定义如何处理没有拼音的字符，


Related Projects
-----------------

* [hotoo/pinyin](https://github.com/hotoo/pinyin): 汉语拼音转换工具 Node.js/JavaScript 版。
* [mozillazg/python-pinyin](https://github.com/mozillazg/python-pinyin): 汉语拼音转换工具 Python 版。
* [mozillazg/rust-pinyin](https://github.com/mozillazg/rust-pinyin): 汉语拼音转换工具 Rust 版。


pinyin data
-----------------

* 使用 [pinyin-data](https://github.com/mozillazg/pinyin-data) 的拼音数据


License
---------

Under the MIT License.
