package genetic

type Operate byte

const (
	TurnLeftOP Operate = iota
	TurnRightOP
	TurnUpOP
	TurnDownOP
	EatOP
)

type Score int

const (
	LeftSC  Score = 0
	RightSC Score = 0
	UpSC    Score = 0
	DownSC  Score = 0

	HitBackSC    Score = -5
	EatNothingSC Score = -2
	EatThingSC   Score = 10
)
