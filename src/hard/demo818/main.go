package main

func racecar(target int) int {

}

func commandA(pos, spd int) (nextPos, nextSpd int) {
	return pos + spd, spd * 2
}

func commandR(pos, spd int) (nextPos, nextSpd int) {
	if spd > 0 {
		return pos, -1
	} else {
		return pos, 1
	}
}

func resolve(src uint32) (pos, spd uint16) {
	return uint16(src >> 16), uint16(src)
}

func combine(pos, spd uint16) (res uint32) {
	return uint32(pos)<<16 | uint32(spd)
}
