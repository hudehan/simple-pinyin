package main

import (
	"flag"
	"fmt"
	pinyin "github.com/hudehan/simple-pinyin"
	"github.com/mattn/go-isatty"
	"io"
	"os"
	"strings"
)

func main() {
	heteronym := flag.Bool("e", false, "启用多音字模式")
	needHelp := flag.Bool("h", false, "获取帮助")
	flag.Parse()
	hans := flag.Args()
	var stdin []byte
	if !isatty.IsTerminal(os.Stdin.Fd()) {
		stdin, _ = io.ReadAll(os.Stdin)
	}
	if len(stdin) > 0 {
		hans = append(hans, string(stdin))
	}

	if *needHelp {
		fmt.Fprintln(os.Stderr, "-e : 开启多音字")
		fmt.Fprintln(os.Stderr, "-h : 查看帮助")
		os.Exit(1)
	}

	if len(hans) == 0 {
		fmt.Fprintln(os.Stderr, "请至少输入一个汉字,: pinyin [-e]  HANS ")
		fmt.Fprintln(os.Stderr, "-h : 查看帮助")
		os.Exit(1)
	}

	res := pinyin.Pinyin(strings.Join(hans, ""), pinyin.Args{
		Style:     pinyin.Str,
		Heteronym: *heteronym,
	})

	for _, s := range res {
		fmt.Println(s, " ")
	}
}
