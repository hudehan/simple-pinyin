package main

import (
	"flag"
	"fmt"
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
		fmt.Fprintln(os.Stderr, "请至少输入一个汉字,: pinyin [-e]  HANS [HANS ...]")
		fmt.Fprintln(os.Stderr, "-h : 查看帮助")
		os.Exit(1)
	}

	//pys := (strings.Join(hans, ""), *heteronym)
	pys := make([][]string, len(hans))
	fmt.Print(*heteronym)
	for _, s := range pys {
		fmt.Print(strings.Join(s, ","), " ")
	}
	if len(pys) > 0 {
		fmt.Println()
	}
}
