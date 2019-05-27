package main

import (
	"unicode/utf8"
)

func main() {

}

/*
(1) "好后缀"的位置以最后一个字符为准. 假定"ABCDEF"的"EF"是好后缀, 则它的位置以"F"为准, 即5(从0开始计算)

(2) 如果"好后缀"在搜索词中只出现一次, 则它的上一次出现位置为 -1. 比如, "EF"在"ABCDEF"之中只出现一次, 则它的上一次出现位置为-1（即未出现）

(3) 如果"好后缀"有多个, 则除了最长的那个"好后缀", 其他"好后缀"的上一次出现位置必须在头部. 比如, 假定"BABCDAB"的"好后缀"是"DAB"、"AB"、"B", 请问这时"好后缀"的上一次出现位置是什么? 回答是, 此时采用的好后缀是"B", 它的上一次出现位置是头部, 即第 0 位. 这个规则也可以这样表达: 如果最长的那个"好后缀"只出现一次, 则可以把搜索词改写成如下形式进行位置计算"(DA)BABCDAB", 即虚拟加入最前面的"DA"
*/

type SkipTable map[rune]int

func BuildSkipTable(needle string) SkipTable {
	table := make(map[rune]int)

	l := utf8.RuneCountInString(needle)
	runes := []rune(needle)

	for i, v := range runes {
		table[v] = l - i - 1
	}

	return table
}

func (table SkipTable) Search(hayStack, needle string) int {
	i, c := 0, 0
	hrunes := []rune(hayStack)
	nrunes := []rune(needle)
	hl := utf8.RuneCountInString(hayStack)
	nl := utf8.RuneCountInString(needle)

	return 0
}
