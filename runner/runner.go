package runner

import (
	"fmt"
)

type PerfectHeritageData struct {
	Sum    int
	Values [3]int
}

type Runner struct {
	IterationCount  int
	PerfectData     *PerfectHeritageData
	Data            *HeritageData
	CurrentSequence *State
	BestSequence    *State
}

func (rn *Runner) Run() {
	var stepInShoulder = 0

	for i := 0; i < rn.IterationCount; i++ {
		stepInShoulder = 0
		rn.CurrentSequence = InitialState(rn.Data, rn.PerfectData)
		for {
			bestNeighbor := rn.CurrentSequence.ChooseNeighbor()
			if bestNeighbor.Value() < rn.CurrentSequence.Value() {
				rn.CurrentSequence = bestNeighbor
				if rn.CurrentSequence.Value() < rn.BestSequence.Value() {
					rn.BestSequence = bestNeighbor
					fmt.Println(rn.BestSequence.OverallValue)
				}
			} else if bestNeighbor.Value() == rn.CurrentSequence.Value() {
				stepInShoulder++
				rn.CurrentSequence = bestNeighbor
				if stepInShoulder == 10 {
					break
				}
			} else {
				break
			}
		}
	}
	fmt.Printf("%#v", rn.BestSequence)
}

func CalculatePerfect(data *HeritageData) PerfectHeritageData {
	perfect := PerfectHeritageData{}
	for i := 0; i < HeritageDataCount; i++ {
		perfect.Sum += (*data)[i]
	}

	perfect.Values[0] = int(float64(perfect.Sum) * 0.4)
	perfect.Values[1] = int(float64(perfect.Sum) * 0.4)
	perfect.Values[2] = int(float64(perfect.Sum) * 0.2)

	return perfect
}
