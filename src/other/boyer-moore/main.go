package main

import (
	"github.com/cubicdaiya/bms"
	"unicode/utf8"
)

func main() {
	println(bms.Search("HERE IS A SIMPLE EXAMPLE", "EXAMPLE"))
}

/*
(1) "好后缀"的位置以最后一个字符为准. 假定"ABCDEF"的"EF"是好后缀, 则它的位置以"F"为准, 即5(从0开始计算)

(2) 如果"好后缀"在搜索词中只出现一次, 则它的上一次出现位置为 -1. 比如, "EF"在"ABCDEF"之中只出现一次, 则它的上一次出现位置为-1（即未出现）

(3) 如果"好后缀"有多个, 则除了最长的那个"好后缀", 其他"好后缀"的上一次出现位置必须在头部. 比如, 假定"BABCDAB"的"好后缀"是"DAB"、"AB"、"B", 请问这时"好后缀"的上一次出现位置是什么? 回答是, 此时采用的好后缀是"B", 它的上一次出现位置是头部, 即第 0 位. 这个规则也可以这样表达: 如果最长的那个"好后缀"只出现一次, 则可以把搜索词改写成如下形式进行位置计算"(DA)BABCDAB", 即虚拟加入最前面的"DA"
*/

type SkipTable struct {
	Table  map[rune]int
	Origin string
}

func BuildSkipTable(needle string) SkipTable {
	table := make(map[rune]int)

	l := utf8.RuneCountInString(needle)
	runes := []rune(needle)

	for i, v := range runes {
		table[v] = l - i - 1
	}

	return SkipTable{
		Table:  table,
		Origin: needle,
	}
}

// search a needle in haystack and return count of needle.
// table is build by BuildSkipTable.
func (sk *SkipTable) Search(hayStack string) []int {
	// i: 当前 haystack 下标; c: 已匹配成功数
	i, result := 0, []int{}
	hrunes := []rune(hayStack)
	nrunes := []rune(sk.Origin)
	hl := utf8.RuneCountInString(hayStack)
	nl := utf8.RuneCountInString(sk.Origin)

	if hl == 0 || nl == 0 || hl < nl {
		return []int{}
	}

	if hl == nl && hayStack == sk.Origin {
		return []int{}
	}

loop:
	/*
		SIMPLE DXAMPLE
		 DXAMPLE
	*/
	for i+nl <= hl {
		// j -- 下标
		for j := nl - 1; j >= 0; j-- {
			if hrunes[i+j] != nrunes[j] {
				if _, ok := sk.Table[hrunes[i+j]]; !ok {
					if j == nl-1 {
						i += nl
					} else {
						i += nl - j - 1
					}
				} else {
					// nl 是长度, nl-1 才是下标
					// 表示 两个字符串不匹配的字符在 nrunes 中 [下标差]
					// table[hrunes[i+j]] --> (nl - k - 1)
					// n --> j - k // 下 - 上
					n := sk.Table[hrunes[i+j]] - (nl - j - 1)
					// n==0 的情况有可能吗 ?
					// n<0 : hrunes[i+j] 在 nrunes[j] 后, 同时前面已经匹配成功了的一个上字符
					// n>0 : nrunes[j] 在 hrunes[i+j] 后
					if n <= 0 { // n<=0 的操作可能只是一个保守做法 ?
						i++
					} else {
						i += n
					}
				}
				goto loop
			}
		}

		result = append(result, i)
		// 为下一次找匹配
		if _, ok := sk.Table[hrunes[i+nl-1]]; ok {
			i += sk.Table[hrunes[i+nl-1]]
		} else {
			i += nl
		}
	}

	return result
}
