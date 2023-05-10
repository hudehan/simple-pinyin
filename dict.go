package pinyin

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// InitSupplement pinyin.json为基础数据，用户可以更改，kTGHZ2013.txt为字典，基础数据没有的话通过字典补充
func initSupplement() {
	// 读取本地文件
	pinyinJson, err := os.ReadFile("./pinyin.json")
	if err != nil {
		return
	}
	json.Unmarshal(pinyinJson, &pinyinMap)
	txt, err := os.ReadFile("./kTGHZ2013.txt")
	if err != nil {
		return
	}
	rows := strings.Split(string(txt), "\n")
	for _, row := range rows {
		row = strings.ReplaceAll(row, "  ", " ")
		item := strings.Split(row, " ")
		if len(item) == 4 {
			if _, ok := pinyinMap[item[3]]; !ok {
				fmt.Println()
				var sb strings.Builder
				set := make(map[string]struct{}, 0)
				for _, letter := range item[1] {
					s := phoneticSymbol[string(letter)]
					if s != "" {
						sb.WriteString(s)
					} else {
						sb.WriteString(string(letter))
					}
				}
				if _, ok = set[item[3]]; !ok {
					pinyinMap[item[3]] = append(pinyinMap[item[3]], sb.String())
					set[item[3]] = struct{}{}
				}
			}
		}
	}
	fmt.Sprint("初始化完成")
}
