/*
 */
package main

import (
	"crypto/rand"
	"encoding/binary"
	"flag"
	"fmt"
	"sort"
	"sync"
	"sync/atomic"
)

var (
	randLen = flag.Uint("rand-len", 0, "")
	percent = flag.Uint64("percent", 10, "允许范围, percent%")
)

func main() {
	flag.Parse()

}

func create() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	threshold := uint32(0) // 认为合格的临界值

	b := make([]byte, int(*randLen*4))
	res := bytes2ints(generate(b))

	dchan := make(chan uint32)
	go func(b []uint32, ch chan<- uint32) {
		realLen, copyb := uint32(uint64(len(b)) / *percent), append([]uint32{}, b...)
		ch <- realLen

		for _, bi := range b {
			ch <- bi
		}

		ge := Generate{G: copyb}
		sort.Sort(&ge)
		atomic.SwapUint32(&threshold, ge.G[uint32(len(copyb))-realLen])

		close(ch)
		wg.Done()
	}(res, dchan)

	for i, max := 0, uint32(0); ; i++ {
		realLen := int(<-dchan)
		d, ok := <-dchan
		if !ok {
			fmt.Println("gg")
			break
		}

		if i > realLen {
			if max < d {
				fmt.Printf("max of before: [%v], better: [%v]\n", max, d)
				// 把剩余的倒掉
				for range dchan {
				}
				break
			}
		} else {
			if max < d {
				max = d
			}
		}
	}

	wg.Wait()
}

type A struct {
}

func generate(b []byte) []byte {
	rand.Read(b)
	return b
}

func bytes2ints(b []byte) (res []uint32) {
	for len(b)/4 > 0 {
		res = append(res, binary.BigEndian.Uint32(b))
		b = b[4:]
	}
	return
}

type Generate struct {
	G []uint32
}

func (g *Generate) Len() int {
	return len(g.G)
}

func (g *Generate) Less(i, j int) bool {
	if g.G[i] < g.G[j] {
		return true
	}

	return false
}

func (g *Generate) Swap(i, j int) {
	g.G[i], g.G[j] = g.G[j], g.G[i]
}
