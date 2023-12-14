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
	neighbors := seq.MeetNeighbors()
	min := &State{OverallValue: MaxInt}
	for i := 0; i < NeighborMeetCount; i++ {
		if neighbors[i].OverallValue < min.OverallValue {
			min = neighbors[i]
		}
	}
	return min
}

func (seq *State) CreateOneNeighbor() *State {
	copySeq := *seq
	minmax := copySeq.MinMax()
	minIdx := minmax[0]
	maxIdx := minmax[1]
	randomExchange := rand.Intn(copySeq.StateDetails.PartCounts[maxIdx])

	var maxIdxCounter int
	var toBeExchangedIdx int
	for i := 0; i < HeritageDataCount; i++ {
		if copySeq.Division[i] == maxIdx {
			maxIdxCounter++
			if maxIdxCounter == randomExchange {
				toBeExchangedIdx = i
				break
			}
		}
	}

	copySeq.StateDetails.PartCounts[maxIdx]--
	copySeq.StateDetails.PartCounts[minIdx]++
	copySeq.Division[toBeExchangedIdx] = minIdx

	toBeExchanged := (*copySeq.Data)[toBeExchangedIdx]
	copySeq.StateDetails.IndividualValues[minIdx] += toBeExchanged
	copySeq.StateDetails.IndividualValues[maxIdx] -= toBeExchanged

	copySeq.OverallValue = 0
	for i := 0; i < 3; i++ {
		copySeq.OverallValue += Abs(copySeq.StateDetails.IndividualValues[i] - copySeq.PerfectData.Values[i])
	}

	return &copySeq
}

func InitialState(data *HeritageData, perfect *PerfectHeritageData) *State {
	var seq = State{}
	for i := 0; i < HeritageDataCount; i++ {
		randomSibling := rand.Intn(3)
		seq.Division[i] = randomSibling

		switch randomSibling {
		case 0:
			seq.StateDetails.PartCounts[0]++
			seq.StateDetails.IndividualValues[0] += (*data)[i]
		case 1:
			seq.StateDetails.PartCounts[1]++
			seq.StateDetails.IndividualValues[1] += (*data)[i]
		case 2:
			seq.StateDetails.PartCounts[2]++
			seq.StateDetails.IndividualValues[2] += (*data)[i]
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

func (seq *State) MeetNeighbors() [NeighborMeetCount]*State {
	neighbors := [NeighborMeetCount]*State{}
	for i := 0; i < NeighborMeetCount; i++ {
		neighbors[i] = seq.CreateOneNeighbor()
	}
	return neighbors
}

func (seq *State) Value() int {
	return seq.OverallValue
}

func (seq *State) MinMax() [2]int {
	max := MinInt
	min := MaxInt
	minIdx, maxIdx := 0, 0

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

	return [2]int{minIdx, maxIdx}
}
