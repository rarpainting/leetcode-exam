package main

import (
	"errors"
	"fmt"
)

const (
	// for effeciency, define default array-size
	startSize = 10
)

func main() {
	kmp, err := NewKMP("abceabf")
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(kmp.ContainedIn("2abceabf2"))
}

// 思想: 当匹配时出现匹配失败时, KMP算法的想法是, 设法利用已匹配的字符串段的信息
// 不要把 "搜索位置" 移回已经比较过的位置, 而是继续把它向后移
// *注*: 该算法的 部分匹配表 与 [阮一峰](http://www.ruanyifeng.com/blog/2013/05/Knuth%E2%80%93Morris%E2%80%93Pratt_algorithm.html) 的相比, prefix[0] 的值固定为 -1 , 其余值向后推移
// 这种做法 **统一了部分匹配和没有匹配的情况**, 这两种情况下都能用 (m+i-prefix[i]) 来表示
type KMP struct {
	// 匹配的字符串
	pattern string
	// 匹配字符串的长度
	size int
	// 部分匹配表
	prefix []int
}

// compile new prefix-array given argument
func NewKMP(pattern string) (*KMP, error) {
	prefix, err := computePrefix(pattern)
	if err != nil {
		return nil, err
	}
	return &KMP{
			pattern: pattern,
			prefix:  prefix,
			size:    len(pattern)},
		nil
}

func computePrefix(pattern string) ([]int, error) {
	// sanity check
	len_p := len(pattern)
	if len_p < 2 {
		if len_p == 0 {
			return nil, errors.New("'pattern' must contain at least one character")
		}
		return []int{-1}, nil
	}
	t := make([]int, len_p)
	t[0], t[1] = -1, 0

	pos, count := 2, 0
	for pos < len_p {

		// | count |  | pos |  |  |  |
		if pattern[pos-1] == pattern[count] {
			count++
			t[pos] = count
			pos++
		} else {
			if count > 0 {
				count = t[count]
			} else {
				t[pos] = 0
				pos++
			}
		}
	}
	return t, nil
}

// return index of **first** occurence of kmp.pattern in argument 's'
// - if not found, returns -1
func (kmp *KMP) FindStringIndex(s string) int {
	// sanity check
	if len(s) < kmp.size {
		return -1
	}
	// i -- 已匹配的字符数
	// m -- s 的当前 prv 位置
	// (i - kmp.prefix[i]) -- 移动位数
	m, i := 0, 0
	for m+i < len(s) {
		if kmp.pattern[i] == s[m+i] {
			if i == kmp.size-1 {
				return m
			}
			i++
		} else {
			m = m + i - kmp.prefix[i]
			if kmp.prefix[i] > -1 {
				i = kmp.prefix[i]
			} else {
				i = 0
			}
		}
	}
	return -1
}

// returns true if pattern i matched at least once
func (kmp *KMP) ContainedIn(s string) bool {
	return kmp.FindStringIndex(s) >= 0
}
