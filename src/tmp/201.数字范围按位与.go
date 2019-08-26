/*
 * @lc app=leetcode.cn id=201 lang=golang
 *
 * [201] 数字范围按位与
 */
// 提示: 找公共部分
package tmp

func rangeBitwiseAnd(m int, n int) int {
	i := uint32(0)

	for m != n {
		m = m >> 1
		n = n >> 1
		i++
	}

	return m << i
}
