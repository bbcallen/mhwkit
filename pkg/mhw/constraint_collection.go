package mhw

import (
	"errors"
	"fmt"
)

type decorationLimitation struct {
	decorationId int
	maxCount     int
}

func newDecorationLimitation() *decorationLimitation {
	dl := &decorationLimitation{}
	return dl
}

type constraintCollection struct {
	dataManager               *dataManager
	requiredSkillEnhancements []skillEnhancement
	decorationLimitations     []decorationLimitation
	extraSlotCountGroupBySize [4]int
}

func newConstraintCollection(dm *dataManager) *constraintCollection {
	cc := &constraintCollection{}
	cc.dataManager = dm
	return cc
}

func (cc *constraintCollection) addRequiredSkillByName(name string, level int) {
	se := skillEnhancement{}
	se.skillId = cc.dataManager.getSkillIdByName(name)
	for i, v := range cc.requiredSkillEnhancements {
		if v.skillId == se.skillId {
			cc.requiredSkillEnhancements[i].enhancedLevel += level
			return
		}
	}
	se.enhancedLevel = level
	cc.requiredSkillEnhancements = append(cc.requiredSkillEnhancements, se)
}

func (cc *constraintCollection) addDecorationLimitationByName(name string, maxCount int) {
	dl := newDecorationLimitation()
	dl.decorationId = cc.dataManager.getDecorationIdByName(name)
	dl.maxCount = maxCount
	cc.decorationLimitations = append(cc.decorationLimitations, *dl)
}

func (cc *constraintCollection) addExtraSlot(size int, count int) {
	if size <= 0 || size > 4 {
		panic(errors.New(fmt.Sprintf("invalid decorations size: %v", size)))
	}
	cc.extraSlotCountGroupBySize[size-1] += count
}
