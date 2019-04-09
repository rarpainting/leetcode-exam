/*
 * 格雷编码是一个二进制数字系统，在该系统中，两个连续的数值仅有一个位数的差异。
 *
 * 给定一个代表编码总位数的非负整数 n，打印其格雷编码序列。格雷编码序列必须以 0 开头。
 *
 * 示例 1:
 *
 * 输入: 2
 * 输出: [0,1,3,2]
 * 解释:
 * 00 - 0
 * 01 - 1
 * 11 - 3
 * 10 - 2
 *
 * 对于给定的 n，其格雷编码序列并不唯一。
 * 例如，[0,2,3,1] 也是一个有效的格雷编码序列。
 *
 * 00 - 0
 * 10 - 2
 * 11 - 3
 * 01 - 1
 *
 * 示例 2:
 *
 * 输入: 0
 * 输出: [0]
 * 解释: 我们定义格雷编码序列必须以 0 开头。
 * 给定编码总位数为 n 的格雷编码序列，其长度为 2^n。当 n = 0 时，长度为 2^0 = 1。
 * 因此，当 n = 0 时，其格雷编码序列为 [0]。
 *
 *
 */
package main

import (
	"flag"
	"fmt"
)

var (
	n = flag.Int("n", 0, "")
)

func main() {
	flag.Parse()
	res := grayCode2(*n)
	fmt.Println(*n, len(res), res)
}

/*
0. 格雷码结果不唯一
1. 格雷码序列必须以 0 开头
2. 当格雷码 [位数] 为 n 时, 长度为 2^n
3. 格雷码前后两数只相差 1 位
*/

/*
数字转格雷码
*/
func grayCode(n int) []int {
	if n == 0 {
		return []int{0}
	}

	bits := 1 << uint(n)

	res := make([]int, bits)

	for i := 1; i < bits; i++ {
		res[i] = i ^ (i >> 1)
	}

	return res
}

// 利用格雷码的镜像特性
func grayCode2(n int) []int {
	if n == 0 {
		return []int{0}
	}

	res, i, j, mask, resLen := []int{0}, 0, 0, 0, 0

	for i = 0; i < n; i++ {
		resLen = len(res)
		mask = 1 << uint(i)
		for j = resLen - 1; j >= 0; j-- {
			res = append(res, res[j]|mask)
		}
	}

	return res
}

/*
TODO: 直接排列
1. i=0, n[i] = 0
2. i=1, 改变最右的位元
3. i=2, 改变右起的第一个为 1 的位元的左边位元
4. 重复前两步
*/
func grayCode3(n int) []int {
	if n == 0 {
		return []int{0}
	}

	bits := uint(1 << uint(n))

	res := []int{0}
	for i := uint(1); i < bits; i++ {
		pre := uint(res[len(res)-1])
		if i%2 == 1 {
			// 改变最右的位元
			// len-2 是除了第一位, 其余有效位皆为 1 === 0xfffffffe
			pre = pre&(bits-2) | ^pre&1
		} else {
			// 改变右起的第一个为 1 的位元的左边位元
			t, cnt := uint(1), pre
			for (t & 1) != 1 {
				cnt++
				t >>= 1
			}

			// 取反某位
			if (pre & (1 << cnt)) == 0 {
				pre |= (1 << cnt)
			} else {
				pre &= ^(1 << cnt)
			}
		}
	}

	return res
}
