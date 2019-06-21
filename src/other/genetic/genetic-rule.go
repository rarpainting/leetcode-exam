package genetic

import (
	"math/rand"
	"time"
)

// 遗传算法
type Genetic struct {
	g []Operate
}

func GenerateGenetic(grLen int) *Genetic {
	r, remain := rand.New(rand.NewSource(time.Now().Unix())), EatOP+1
	g := make([]Operate, grLen)

	for i := 0; i < grLen; i++ {
		g[i] = Operate(r.Intn(int(remain)))
	}

	return &Genetic{
		g: g,
	}
}

func (g *Genetic) Rule(m *Map, primPos int) Operate {
	return g.g[primPos]
}

type HybridMachan struct {
	r *rand.Rand
}

func NewHybrid() *HybridMachan {
	return &HybridMachan{
		r: rand.New(rand.NewSource(time.Now().Unix())),
	}
}

// 杂交
func (hb *HybridMachan) Hybrid(g1 *Genetic, g2 *Genetic, varCount int, gCount int) []Genetic {
	g := make([]Genetic, gCount)
	lenOfGen, halfOfGen, remain := len(g1.g), len(g1.g)/2, int(EatOP+1)

	g12 := append([]Operate{}, g1.g[:halfOfGen]...)
	g12 = append(g12, g2.g[halfOfGen:]...)
	g21 := append([]Operate{}, g2.g[:halfOfGen]...)
	g21 = append(g21, g1.g[halfOfGen:]...)
	for i := 0; i < gCount/2; i++ {
		lenOfCut := hb.r.Intn(lenOfGen)
		 := hb.r.Intn(remain)
	}

	return g
}
