package runner

type HeritageData = [HeritageDataCount]int

const (
	HeritageDataCount = 100

	MaxUint = ^uint(0)
	MinUint = 0
	MaxInt  = int(MaxUint >> 1)
	MinInt  = -MaxInt - 1
)
