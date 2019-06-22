package genetic

import (
	"math/rand"
	"time"
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
	M       []MapPoint
}

func GenerateMap(rows, columns int) *Map {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	m := Map{Rows: rows, Columns: columns}
	m.M = make([]MapPoint, rows*columns)

	food := 0
	foodLimit := rows * columns / 3
	for i := rows; i < rows*(columns-1); i++ {
		if i%rows == 0 || (i+1)%rows == 0 {
			m.M[i] = MapWall
		} else {
			if food == foodLimit {
				m.M[i] = MapNotThing
				continue
			}

			ri := MapPoint(r.Intn(int(MapWall)))
			if ri == MapFood {
				food++
			}
			m.M[i] = ri
		}
	}

	// 把墙堵上
	for i, lastRows := 0, rows*(columns-1); i < rows; i++ {
		m.M[i] = MapWall
		m.M[i+lastRows] = MapWall
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
		if m.M[pos] == MapFood {
			score, newPos = EatThingSC, pos
			m.M[pos] = MapNotThing
		} else {
			score, newPos = EatNothingSC, pos
		}
	}
	return
}

func (m *Map) String() string {
	str := ""
	for i, v := range m.M {
		switch v {
		case MapFood:
			str += "*"
		case MapNotThing:
			str += " "
		case MapWall:
			str += "W"
		}
		if (i+1)%m.Rows == 0 {
			str += "\n"
		}
	}
	return str
}
