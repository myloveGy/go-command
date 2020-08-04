package cmd

import (
	"cobra/internal/word"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

const (
	ModeUpper                      = iota + 1 // 全部单词转大写
	ModeLower                                 // 全部单词转小写
	ModeUnderscoreToUpperCamelcase            // 下划线单词转为大驼峰单词
	ModeUnderscoreToLowerCamelcase            // 下划线单词转为小驼峰单词
	ModeCamelcaseToUnderscore                 // 驼峰单词转为下划线单词
)

// 说明信息
var desc = strings.Join([]string{
	"该子命令支持各种单词格式转换，模式如下",
	"1: 全部单词转为大写",
	"2: 全部单词转为小写",
	"3: 下划线单词转为大驼峰单词",
	"4: 下划线单词转为小驼峰单词",
	"5: 驼峰单词转为下划线单词",
}, "\n")

var str string
var mode int8

var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "单词格式转换",
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case ModeUpper:
			content = word.ToUpper(str)
		case ModeLower:
			content = word.ToLower(str)
		case ModeUnderscoreToUpperCamelcase:
			content = word.UnderscoreToUpperCamelCase(str)
		case ModeUnderscoreToLowerCamelcase:
			content = word.UnderscoreToLowerCamelCase(str)
		case ModeCamelcaseToUnderscore:
			content = word.CamelCaseToUnderscore(str)
		default:
			log.Fatalf("暂不支持该转换模式，请执行 help word 查看帮助文档")
		}

		log.Printf("输入结果： %s", content)
	},
}

func init() {
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "请输入单词内容")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "请输入单词转换模式")
}
