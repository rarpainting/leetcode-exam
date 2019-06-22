package main

import (
	"flag"
	"fmt"
	"math/rand"
	"other/genetic"
	"sort"
	"sync"
	"sync/atomic"
	"time"
)

var (
	Row    = flag.Int("row", 10, "")
	Column = flag.Int("column", 10, "")

	YoungCount   = flag.Int("young-count", 1000, "新生代数量")
	RunCount     = flag.Int("run-count", 1000, "每代运行次数")
	VariateCount = flag.Int("variate-count", 10, "变异数量")
	TotalStep    = flag.Int("total-step", 200, "步数")
)

func main() {
	flag.Parse()
	wg := sync.WaitGroup{}
	mx, results := sync.Mutex{}, []Result{}

	now := time.Now()
	wg.Add(*YoungCount)
	// 先 YoungCount 次
	for i := 0; i < *YoungCount; i++ {
		go func() {
			// init
			now, totalPos, totalScore := time.Now(), *Row**Column-1, int32(0)
			r := rand.New(rand.NewSource(now.Unix()))
			g := genetic.GenerateGenetic(totalPos)
			swg := sync.WaitGroup{}
			swg.Add(*RunCount)

			for j := 0; j < *RunCount; j++ {
				go func() {
					m := genetic.GenerateMap(*Row, *Column)
					score := genetic.Score(0)

					for k, sc, pos := 0, genetic.Score(0), r.Intn(totalPos); k < *TotalStep; k++ {
						sc, pos = m.Do(pos, g.Rule(m, pos))
						score += sc
					}
					atomic.AddInt32(&totalScore, int32(score))

					swg.Done()
				}()
			}
			swg.Wait()

			mx.Lock()
			results = append(results, Result{
				g:     g,
				score: genetic.Score(int(totalScore) / *RunCount),
				time:  time.Now().Sub(now),
			})
			mx.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()

	// 挑选最高得分的两位
	sort.Sort(ResultSlice(results))
	fmt.Printf("first: [%v] second: [%v] last:[%v] | %v\n", results[0].score, results[1].score, results[len(results)-1].score, time.Now().Sub(now))

	for {
		// 杂交
		g1, g2 := results[0], results[1]
		hybrid := genetic.NewHybrid()
		newGens := hybrid.Hybrid(g1.g, g2.g, *VariateCount, *YoungCount, Hybrid)

		// 清空并初始化
		results = results[0:0]
		wg.Add(*YoungCount)

		for i := 0; i < *YoungCount; i++ {
			now = time.Now()
			go func(i int) {
				defer func() {
					if e := recover(); e != nil {
						fmt.Printf("i:[%v] lenOfNewGens: [%v]\n", i, len(newGens))
						panic(e)
					}
				}()

				// init
				now, totalPos, totalScore := time.Now(), *Row**Column-1, int32(0)
				r := rand.New(rand.NewSource(now.Unix()))
				swg := sync.WaitGroup{}
				swg.Add(*RunCount)

				for j := 0; j < *RunCount; j++ {
					go func() {
						m := genetic.GenerateMap(*Row, *Column)
						score := genetic.Score(0)

						for k, sc, pos := 0, genetic.Score(0), r.Intn(totalPos); k < *TotalStep; k++ {
							sc, pos = m.Do(pos, newGens[i].Rule(m, pos))
							score += sc
						}
						atomic.AddInt32(&totalScore, int32(score))
						swg.Done()
					}()
				}
				swg.Wait()

				mx.Lock()
				results = append(results, Result{
					g:     &newGens[i],
					score: genetic.Score(int(totalScore) / *RunCount),
					time:  time.Now().Sub(now),
				})
				mx.Unlock()
				wg.Done()
			}(i)
		}

		wg.Wait()

		sort.Sort(ResultSlice(results))
		fmt.Printf("first: [%v] second: [%v] | %v\n", results[0].score, results[1].score, time.Now().Sub(now))
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
