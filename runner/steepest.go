package runner

import (
	"math/rand"
)

type StateDetails struct {
	PartCounts       [3]int
	IndividualValues [3]int
}

type State struct {
	Division     [HeritageDataCount]int
	Data         *HeritageData
	PerfectData  *PerfectHeritageData
	StateDetails StateDetails
	OverallValue int
}

func (seq *State) ChooseNeighbor() *State {
	neighbors := seq.CreateMultipleNeighbors()
	min := &State{OverallValue: MaxInt}
	for index := range neighbors {
		if neighbors[index].OverallValue < min.OverallValue {
			min = neighbors[index]
		}
	}
	return min
}

func InitialState(data *HeritageData, perfect *PerfectHeritageData) *State {
	var seq = State{}
	for i := 0; i < HeritageDataCount; i++ {
		randomSibling := rand.Intn(100)

		switch {
		case randomSibling > -1 && randomSibling < 41:
			seq.StateDetails.PartCounts[0]++
			seq.StateDetails.IndividualValues[0] += (*data)[i]
			seq.Division[i] = 0
		case randomSibling > 40 && randomSibling < 81:
			seq.StateDetails.PartCounts[1]++
			seq.StateDetails.IndividualValues[1] += (*data)[i]
			seq.Division[i] = 1
		case randomSibling > 80:
			seq.StateDetails.PartCounts[2]++
			seq.StateDetails.IndividualValues[2] += (*data)[i]
			seq.Division[i] = 2
		default:
			panic("invalid random number...")
		}
	}

	seq.Data = data
	seq.PerfectData = perfect

	seq.OverallValue = 0
	for i := 0; i < 3; i++ {
		seq.OverallValue += Abs(seq.StateDetails.IndividualValues[i] - perfect.Values[i])
	}

	return &seq
}

func (seq *State) CreateMultipleNeighbors() []*State {
	minIdx, maxIdx := seq.MinMax()
	maxDivision := make([]int, seq.StateDetails.PartCounts[maxIdx])
	minDivision := make([]int, seq.StateDetails.PartCounts[minIdx])
	maxCount := seq.StateDetails.PartCounts[maxIdx]
	minCount := seq.StateDetails.PartCounts[minIdx]
	for i,j,k:=0,0,0; i<HeritageDataCount; i++{
		switch seq.Division[i]{
		case maxIdx:
			maxDivision[j] = i
			j++
		case minIdx:
			minDivision[k] = i
			k++
		}
	}
	neighbors := make([]*State, maxCount * minCount)

	var neighborIdx int
	for i:=0; i<maxCount; i++{
		for j:=0; j<minCount; j++{
			neighbors[neighborIdx] = seq.CreateOneNeighbor(minDivision[j], maxDivision[i], minIdx, maxIdx)
			neighborIdx++
		}
	}
	return neighbors
}

func (seq *State) Value() int {
	return seq.OverallValue
}

func (seq *State) MinMax() (int, int) {
	max := MinInt
	min := MaxInt
	var minIdx, maxIdx int

	for i := 0; i < 3; i++ {
		value := seq.StateDetails.IndividualValues[i]
		perfectValue := seq.PerfectData.Values[i]
		if value-perfectValue < min {
			min = value - perfectValue
			minIdx = i
		}
		if value-perfectValue > max {
			max = value - perfectValue
			maxIdx = i
		}
	}

	return minIdx, maxIdx
}

func (seq *State) CreateOneNeighbor(taker, giver, minIdx, maxIdx int) *State {
	copySeq := *seq

	takerData := copySeq.Data[taker]
	giverData := copySeq.Data[giver]

	copySeq.Division[taker], copySeq.Division[giver] = maxIdx, minIdx
	copySeq.StateDetails.IndividualValues[minIdx] += -1*takerData + giverData
	copySeq.StateDetails.IndividualValues[maxIdx] += -1*giverData + takerData

	copySeq.OverallValue = 0
	for i := 0; i < 3; i++ {
		copySeq.OverallValue += Abs(copySeq.StateDetails.IndividualValues[i] - copySeq.PerfectData.Values[i])
	}

	return &copySeq
}
