/*
CPU 有多少就吃多少型遗传算法
*/
package genetic

import (
	"math/rand"
	"time"
)

// 遗传算法
type Genetic struct {
	G []Operate
}

func GenerateGenetic(grLen int) *Genetic {
	realLen := grLen * 3 // 总共有 pos * 3
	r, remain := rand.New(rand.NewSource(time.Now().Unix())), EatOP+1
	g := make([]Operate, realLen)

	for i := 0; i < realLen; i++ {
		g[i] = Operate(r.Intn(int(remain)))
	}

	return &Genetic{
		G: g,
	}
}

func (g *Genetic) Rule(m *Map, primPos int) Operate {
	return g.G[primPos*int(m.M[primPos])]
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
// varCount: 变异的数量
// gCount: 杂交的次一代的数量
func (hb *HybridMachan) Hybrid(g1 *Genetic, g2 *Genetic, varCount int, gCount int, bybrid func(g1, g2 *Genetic) (g12, g21 *Genetic)) []Genetic {
	gs := make([]Genetic, gCount)
	lenOfGen, remain := len(g1.G), int(EatOP+1)

	g12, g21 := bybrid(g1, g2)
	for i := 0; i < gCount/2; i++ {
		newG12 := Genetic{
			G: append([]Operate{}, g12.G...),
		}
		newG21 := Genetic{
			G: append([]Operate{}, g21.G...),
		}
		for j := 0; j < varCount; j++ {
			newG12.G[hb.r.Intn(lenOfGen)] = Operate(hb.r.Intn(remain))
			newG21.G[hb.r.Intn(lenOfGen)] = Operate(hb.r.Intn(remain))
		}
		gs[2*i], gs[2*i+1] = newG12, newG21
	}

	return gs
}

func (g *Genetic) String() string {
	str := ""
	for _, v := range g.G {
		switch v {
		case TurnLeftOP:
			str += " <-"
		case TurnRightOP:
			str += " ->"
		case TurnUpOP:
			str += " v"
		case TurnDownOP:
			str += " ^"
		case EatOP:
			str += " E"
		}
	}

	return str
}
