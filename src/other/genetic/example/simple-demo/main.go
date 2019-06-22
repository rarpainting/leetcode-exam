package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"other/genetic"
	"sort"
	"sync"
	"time"
)

var (
	Row    = flag.Int("row", 10, "")
	Column = flag.Int("column", 10, "")

	YoungCount   = flag.Int("young-count", 1000, "新生代数量")
	RunCount     = flag.Int("run-count", 1000, "每代运行次数")
	VariateCount = flag.Int("variate-count", 20, "变异数量")
	TotalStep    = flag.Int("total-step", 200, "步数")
)

func main() {
	go func() {
		panic(http.ListenAndServe(":6060", nil))
	}()

	flag.Parse()
	wg := sync.WaitGroup{}
	mx, results := sync.Mutex{}, []Result{}

	now := time.Now()
	wg.Add(*YoungCount)
	// 先 YoungCount 次
	for i := 0; i < *YoungCount; i++ {
		go func() {
			// init
			totalPos := *Row**Column - 1
			g := genetic.GenerateGenetic(totalPos)

			// 时间阻塞在 Done 方法上, 得不偿失
			// swg := sync.WaitGroup{}
			// swg.Add(*RunCount)

			res := generateResult(g)

			mx.Lock()
			results = append(results, *res)
			mx.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()

	// 挑选最高得分的两位
	sort.Sort(ResultSlice(results))
	fmt.Printf("first: [%v] second: [%v] last:[%v] | %v\n", results[0].score, results[1].score, results[len(results)-1].score, time.Now().Sub(now))

	for genNum := 0; ; genNum++ {
		now = time.Now()
		// 杂交
		g1, g2 := results[0], results[1]
		hybrid := genetic.NewHybrid()
		newGens := hybrid.Hybrid(g1.g, g2.g, *VariateCount, *YoungCount, Hybrid)

		// 清空并初始化
		results = results[0:0]
		wg.Add(*YoungCount)

		for i := 0; i < *YoungCount; i++ {
			go func(i int) {
				res := generateResult(&newGens[i])
				mx.Lock()
				results = append(results, *res)
				mx.Unlock()
				wg.Done()
			}(i)
		}

		wg.Wait()

		sort.Sort(ResultSlice(results))
		fmt.Printf("{%v} first: [%v] second: [%v] last:[%v] | %v\n", genNum,
			results[0].score, results[1].score, results[len(results)-1].score,
			time.Now().Sub(now))
	}
}

type Result struct {
	g     *genetic.Genetic
	score genetic.Score
	time  time.Duration
}

type ResultSlice []Result

func (r ResultSlice) Len() int {
	return len(r)
}

func (r ResultSlice) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

// 反过来
func (r ResultSlice) Less(i, j int) bool {
	if r[i].score > r[j].score {
		return true
	}
	return false
}

func Hybrid(g1, g2 *genetic.Genetic) (g12, g21 *genetic.Genetic) {
	halfOfLen := len(g1.G) / 2
	g12 = &genetic.Genetic{
		G: append([]genetic.Operate{}, g1.G[:halfOfLen]...),
	}
	g12.G = append(g12.G, g2.G[halfOfLen:]...)

	g21 = &genetic.Genetic{
		G: append([]genetic.Operate{}, g2.G[:halfOfLen]...),
	}
	g21.G = append(g21.G, g1.G[halfOfLen:]...)

	return
}

func generateResult(g *genetic.Genetic) *Result {
	// init
	now, totalPos, totalScore := time.Now(), *Row**Column-1, genetic.Score(0)
	r := rand.New(rand.NewSource(now.Unix()))

	for j := 0; j < *RunCount; j++ {
		m := genetic.GenerateMap(*Row, *Column)
		score := genetic.Score(0)

		for k, sc, pos := 0, genetic.Score(0), r.Intn(totalPos); k < *TotalStep; k++ {
			sc, pos = m.Do(pos, g.Rule(m, pos))
			score += sc
		}
		totalScore += score
	}

	return &Result{
		g:     g,
		score: totalScore / genetic.Score(*RunCount),
		time:  time.Now().Sub(now),
	}
}
