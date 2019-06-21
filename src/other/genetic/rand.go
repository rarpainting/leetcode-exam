package genetic

/*
1. 生成 n 张随机地图, 原始人位置为 m
2. 设置 最大步数为 step, 每个操作为 1 个步数
3.
*/
func Run(m *Map, stepCount int, primFirstPos int, rule func(m *Map, primPos int) (op Operate)) (totalCount int) {
	totalCount = 0
	primPos, count := primFirstPos, Score(0)
	for i := 0; i < stepCount; i++ {
		count, primPos = m.Do(primPos, rule(m, primPos))
		totalCount += int(count)
	}
	return
}
