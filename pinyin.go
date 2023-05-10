package pinyin

import (
	"strings"
	"sync"
)

// Args 配置信息
type Args struct {
	Style     int  // 拼音风格（默认： Normal)
	Heteronym bool // 是否启用多音字模式（默认：禁用）

	// 处理没有拼音的字符（默认忽略没有拼音的字符）
	// 如果是空则代表不处理
	Fallback func(ch string) string
}

const (
	Normal    = iota // 普通风格，返回类型为[chang ge xing] 开启多音字后为 [chang|zhange ge  xing|hang]
	Str              // 字符串格式 [changgexing] 开启多音字后为 [changgexing zhanggexing changgehang zhanggehang]
	Prefix           // 首字母风格，返回类型为[c g x]	开启多音字后为 [c|z g|x]
	PrefixStr        // 首字母字符串格式 [cgx]	开启多音字后为 [cgx zgx cgh zgh]
)

var pinyinMap = map[string][]string{}
var once = sync.Once{}

func init() {
	once.Do(
		func() {
			initSupplement()
		})
}

func Pinyin(word string, args Args) (res []string) {
	switch args.Style {
	case Normal:
		res = TransferPinYinNormal(word, args.Fallback, args.Heteronym)
	case Str:
		res = TransferPinYinStr(word, args.Fallback, args.Heteronym)
	case Prefix:
		res = TransferPinYinPrefix(word, args.Fallback, args.Heteronym)
	case PrefixStr:
		res = TransferPinYinPrefixStr(word, args.Fallback, args.Heteronym)
	}

	return res
}

func TransferPinYinNormal(word string, fallback func(ch string) string, heteronym bool) (res []string) {
	queue := []string{""}
	for _, w := range word {
		// 通过bfs算法，找出所有的组合
		if arr, ok := pinyinMap[string(w)]; ok {
			if heteronym {
				var sb []string
				for k := 0; k < len(arr); k++ {
					sb = append(sb, arr[k])
				}
				queue = append(queue, strings.Join(sb, "|"))
			} else {
				queue = append(queue, arr[0])
			}
		} else {
			for j := 0; j < len(queue); j++ {
				queue = append(queue, fallback(string(w)))
			}
		}
	}
	return queue[1:]
}

func TransferPinYinPrefix(word string, fallback func(ch string) string, heteronym bool) (res []string) {
	queue := []string{""}
	for _, w := range word {
		// 通过bfs算法，找出所有的组合
		if arr, ok := pinyinMap[string(w)]; ok {
			if heteronym {
				var sb []string
				for k := 0; k < len(arr); k++ {
					sb = append(sb, arr[k][:1])
				}
				queue = append(queue, strings.Join(sb, "|"))
			} else {
				queue = append(queue, arr[0][:1])
			}
		} else {
			for j := 0; j < len(queue); j++ {
				queue = append(queue, fallback(string(w)))
			}
		}
	}
	return queue[1:]
}

func TransferPinYinStr(word string, fallback func(ch string) string, heteronym bool) (res []string) {
	queue := []string{""}
	for _, w := range word {
		var tmp []string
		// 通过bfs算法，找出所有的组合
		if arr, ok := pinyinMap[string(w)]; ok {
			for j := 0; j < len(queue); j++ {
				if heteronym {
					for k := 0; k < len(arr); k++ {
						tmp = append(tmp, queue[j]+arr[k])
					}
				} else {
					tmp = append(tmp, queue[j]+arr[0])
				}
			}
		} else {
			for j := 0; j < len(queue); j++ {
				tmp = append(tmp, queue[j]+fallback(string(w)))
			}
		}
		queue = tmp
	}
	return queue
}

func TransferPinYinPrefixStr(word string, fallback func(ch string) string, heteronym bool) (res []string) {
	queue := []string{""}
	for _, w := range word {
		var tmp []string
		// 通过bfs算法，找出所有的组合
		if arr, ok := pinyinMap[string(w)]; ok {
			for j := 0; j < len(queue); j++ {
				if heteronym {
					for k := 0; k < len(arr); k++ {
						tmp = append(tmp, queue[j]+arr[k][:1])
					}
				} else {
					tmp = append(tmp, queue[j]+arr[0])
				}
			}
		} else {
			for j := 0; j < len(queue); j++ {
				tmp = append(tmp, queue[j]+fallback(string(w)))
			}
		}
		queue = tmp
	}
	return queue
}
