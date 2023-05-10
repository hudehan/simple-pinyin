package pinyin

import (
	"fmt"
	"testing"
)

func TestPinyin(t *testing.T) {
	type args struct {
		word string
		args Args
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "first",
			args: args{
				word: "长歌行",
				args: Args{
					Style:     Str,
					Heteronym: true,
					Fallback: func(ch string) string {
						return ch
					},
				},
			},
		},
		{
			name: "first",
			args: args{
				word: "长歌行第1季",
				args: Args{
					Style:     Str,
					Heteronym: true,
					Fallback: func(ch string) string {
						return ch
					},
				},
			},
		},
		{
			name: "只是获取前缀多音字",
			args: args{
				word: "长歌行",
				args: Args{
					Style:     PrefixStr,
					Heteronym: true,
					Fallback: func(ch string) string {
						return ch
					},
				},
			},
		},
		{
			name: "noraml开启多音字",
			args: args{
				word: "长歌行",
				args: Args{
					Style:     Normal,
					Heteronym: true,
					Fallback: func(ch string) string {
						return ch
					},
				},
			},
		},
		{
			name: "noraml不开启多音字",
			args: args{
				word: "长歌行",
				args: Args{
					Style:     Normal,
					Heteronym: false,
					Fallback: func(ch string) string {
						return ch
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := Pinyin(tt.args.word, tt.args.args)
			fmt.Printf("res:%v\n", res)
		})
	}
}
