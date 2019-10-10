/*
堆排序
*/
package main

import (
	"fmt"
)

// 构建最小堆
func main() {
	heapSort := []int{9, 6, 5, 4, 3, 2, 1}

	header := heapSort
	for len(header) != 0 {
		sortonce(header)
		header = header[1:]
	}

	fmt.Println(heapSort)
}

func sortonce(list []int) {
	// 最后一个非叶子节点
	// 构建最小堆
	for i, length := len(list)>>1-1, len(list); i >= 0; i-- {
		// heapSort > 左节点
		if list[i] > list[i<<1+1] {
			list[i], list[i<<1+1] = list[i<<1+1], list[i]
		}
		if length > i<<1+2 && list[i] > list[i<<1+2] {
			list[i], list[i<<1+2] = list[i<<1+2], list[i]
		}
	}
}
