package mhw

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadSlotCombinations(t *testing.T) {
	dm := newDataManager()
	dm.loadSlotCombinations()

	assert.Equal(t, 20, len(dm.slotCombinationOverridingMatrix), "mismatch count of slot combinations")

	if testing.Verbose() {
		for i, _ := range dm.slotCombinationOverridingMatrix {
			sc := dm.getSlotCombinationById(i)
			fmt.Printf("%2v%v: ", i, sc.layout)
			for _, op := range dm.slotCombinationOverridingMatrix[i] {
				switch op {
				case slotCombinationCannotOverride:
					fmt.Printf("?")
				case slotCombinationOverrideLeft:
					fmt.Printf("<")
				case slotCombinationOverrideRight:
					fmt.Printf(">")
				case slotCombinationOverrideEachOther:
					fmt.Printf("=")
				}
			}
			fmt.Println("")
		}
	}
}
