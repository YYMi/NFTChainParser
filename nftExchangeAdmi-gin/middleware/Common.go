package middleware

import (
	"regexp"
	"strings"
)

// wildcardToRegex 将通配符模式转换为正则表达式
func wildcardToRegex(pattern string) string {
	// 转义特殊字符
	regex := regexp.QuoteMeta(pattern)
	// 将通配符替换为正则表达式
	regex = strings.ReplaceAll(regex, `\*\*`, `.*`)  // ** 替换为 .*
	regex = strings.ReplaceAll(regex, `\*`, `[^/]*`) // * 替换为除 / 外的任意字符
	regex = strings.ReplaceAll(regex, `\?`, `.`)     // ? 替换为任意单个字符
	// 添加起始和结束锚定
	return `^` + regex + `$`
}
