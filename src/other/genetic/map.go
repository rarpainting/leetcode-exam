package genetic

import (
	"math/rand"
)

type MapPoint byte

const (
	MapFood MapPoint = iota
	MapNotThing
	MapWall
)

type Map struct {
	Rows    int
	Columns int
	m       []MapPoint
}

func Generate(rows, columns int) *Map {
	m := Map{Rows: rows, Columns: columns}
	m.m = make([]MapPoint, rows*columns)

	food := 0
	foodLimit := rows * columns / 2
	for i := rows; i < rows*(columns-1); i++ {
		if i/rows == 0 || (i+1)/rows == 0 {
			m.m[i] = MapWall
		} else {
			if food == foodLimit {
				m.m[i] = 1
				continue
			}

			ri := MapPoint(rand.Intn(2))
			if ri == MapFood {
				food++
			}
			m.m[i] = ri
		}
	}

	// 把墙堵上
	for i, lastRows := 0, rows*(columns-1); i < rows; i++ {
		m.m[i] = MapWall
		m.m[i+lastRows] = MapWall
	}

	return &m
}

func (m *Map) Do(pos int, op Operate) (score Score, newPos int) {
	switch op {
	case TurnLeftOP:
		if pos%m.Rows < 2 { // 墙边
			score, newPos = HitBackSC, pos
		} else {
			score, newPos = LeftSC, pos-1
		}
	case TurnRightOP:
		if m.Rows-pos%m.Rows > 1 {
			score, newPos = HitBackSC, pos
		} else {
			score, newPos = RightSC, pos+1
		}
	case TurnUpOP:
		if rePos := pos - m.Rows; rePos < m.Rows {
			score, newPos = HitBackSC, pos%m.Rows+m.Rows
		} else {
			score, newPos = UpSC, pos-m.Rows
		}
	case TurnDownOP:
		if rePos := pos - m.Rows*(m.Columns-2) - 1; rePos > 0 {
			score, newPos = HitBackSC, pos%m.Rows+m.Rows*(m.Columns-2)
		} else {
			score, newPos = DownSC, pos+m.Rows
		}
	case EatOP:
		if m.m[pos] == MapFood {
			score, newPos = EatThingSC, pos
		} else {
			score, newPos = EatNothingSC, pos
		}
	}
	return
}
