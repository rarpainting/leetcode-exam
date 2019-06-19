/*
 */
package main

import (
	"crypto/rand"
	"encoding/binary"
	"flag"
	"fmt"
	"math/big"
	"sort"
)

var (
	randLen = flag.Uint("rand-len", 0, "")
	seedx   = flag.Int64("parts", 1, "")
)

func main() {
	flag.Parse()

	bigInt := big.NewInt(*seedx)
	rand.Int(rand.Reader, bigInt)
}

func create() {
	b := make([]byte, int(*randLen*4))
	res := bytes2ints(generate(b))

	dchan := make(chan uint32)
	go func(b []uint32, ch chan<- uint32) {
		ch <- uint32(len(b))
		for _, bi := range b {
			ch <- bi
		}

		close(ch)

		ge := Generate{G: b}
		sort.Sort(&ge)
	}(res, dchan)

	getLen := <-dchan
	realLen := int(float64(getLen) / *parts)

	for i, max := 0, uint32(0); ; i++ {
		d, ok := <-dchan
		if !ok {
			fmt.Println("gg")
			return
		}

		if i > realLen {
			if max < d {
				fmt.Printf("max of before: [%v], better: [%v]\n", max, d)
				return
			}
		} else {
			if max < d {
				max = d
			}
		}
	}
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
