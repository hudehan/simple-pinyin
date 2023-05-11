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

	// 把map输出为go文件
	var sb strings.Builder
	sb.WriteString("package pinyin\n\n")
	sb.WriteString("var pinyinMap = map[string][]string{\n")
	for k, v := range pinyinMap {
		sb.WriteString("\t\"" + k + "\": {\"" + strings.Join(v, "\", \"") + "\"},\n")
	}
	sb.WriteString("}")
	dictByte := []byte(sb.String())
	os.WriteFile("./dict.go", dictByte, 0666)
	fmt.Sprint("初始化完成")
}
