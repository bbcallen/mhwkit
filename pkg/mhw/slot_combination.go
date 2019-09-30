package mhw

import (
	"errors"
	"fmt"
)

const (
	slotCombinationCannotOverride    = 0
	slotCombinationOverrideLeft      = 1
	slotCombinationOverrideRight     = 2
	slotCombinationOverrideEachOther = 3
)

type slotCombination struct {
	id               int
	layout           [3]int
	sizeDistribution [4]int
	slotCount_       int
}

func newSlotCombinationFromLayout(layout [3]int) *slotCombination {
	sc := &slotCombination{}
	sc.id = 0
	sc.sizeDistribution = [4]int{}

	if layout[0] < layout[1] || layout[1] < layout[2] {
		panic(errors.New("invalid layout order "))
	}

	copy(sc.layout[:], layout[:])
	for _, size := range layout {
		if size < 0 || size > 4 {
			panic(errors.New("invalid slot size"))
		}

		if size > 0 {
			sc.sizeDistribution[size-1]++
			sc.slotCount_++
		}
	}

	return sc
}

func (sc *slotCombination) slotCount() int {
	return sc.slotCount_
}

func (sc *slotCombination) hasAnySlot() bool {
	return sc.slotCount() > 0
}

/*
returns:
	slotCombinationCannotOverride
	slotCombinationOverrideLeft
	slotCombinationOverrideRight
	slotCombinationOverrideEachOther
*/
func (sc *slotCombination) compareTo(rhs *slotCombination) int {
	currentResult := slotCombinationOverrideEachOther
	for i, _ := range sc.layout {
		switch {
		case sc.layout[i] > rhs.layout[i]:
			if currentResult == slotCombinationOverrideLeft {
				return slotCombinationCannotOverride
			}
			currentResult = slotCombinationOverrideRight
		case sc.layout[i] < rhs.layout[i]:
			if currentResult == slotCombinationOverrideRight {
				return slotCombinationCannotOverride
			}
			currentResult = slotCombinationOverrideLeft
		}
	}
	return currentResult
}

func (dm *dataManager) loadSlotCombinations() {
	if len(dm.slotCombinations) > 0 {
		panic(errors.New(fmt.Sprintf("duplicated slot combinations loading")))
	}

	slotSizes := []int{0, 1, 2, 3, 4}
	for i, x := range slotSizes {
		for j, y := range slotSizes[:i+1] {
			for _, z := range slotSizes[:j+1] {
				sc := newSlotCombinationFromLayout([3]int{x, y, z})
				dm.registerSlotCombination(*sc)
				// fmt.Printf("sc[%v] = %v\n", sc.id, sc)
			}
		}
	}

	/*
	     for layout vectors [a, b, c] those a >= b >= c
	     we say:
		    layoutX [x0, x1, x2] can override layoutY [y0, y1, y2]
		    only if x0 >= y0 and x1 >= y1 and x2 > = y2
	*/
	overrideCounts := make([]int, len(dm.slotCombinations))
	dm.slotCombinationOverridingMatrix = make([][]int, len(dm.slotCombinations))
	for i, _ := range dm.slotCombinations {
		dm.slotCombinationOverridingMatrix[i] = make([]int, len(dm.slotCombinations))
		for j, _ := range dm.slotCombinations[:i] {
			left := dm.getSlotCombinationById(i)
			right := dm.getSlotCombinationById(j)
			switch left.compareTo(right) {
			case slotCombinationOverrideLeft:
				dm.slotCombinationOverridingMatrix[i][j] = slotCombinationOverrideLeft
				dm.slotCombinationOverridingMatrix[j][i] = slotCombinationOverrideRight
			case slotCombinationOverrideRight:
				dm.slotCombinationOverridingMatrix[i][j] = slotCombinationOverrideRight
				dm.slotCombinationOverridingMatrix[j][i] = slotCombinationOverrideLeft
				overrideCounts[i]++
			case slotCombinationOverrideEachOther:
				dm.slotCombinationOverridingMatrix[i][j] = slotCombinationOverrideEachOther
				dm.slotCombinationOverridingMatrix[j][i] = slotCombinationOverrideEachOther
			case slotCombinationCannotOverride:
				dm.slotCombinationOverridingMatrix[i][j] = slotCombinationCannotOverride
				dm.slotCombinationOverridingMatrix[j][i] = slotCombinationCannotOverride
			default:
				panic(errors.New(fmt.Sprintf("invalid relationship between %v and %v", left.layout, right.layout)))
			}
		}
		dm.slotCombinationOverridingMatrix[i][i] = slotCombinationOverrideEachOther
	}

	/*
	     dm.slotCombinations is topological sorted index in slotCombination.id
		 if layoutA can override layoutB, layoutA.id would have larger id
	*/
}
