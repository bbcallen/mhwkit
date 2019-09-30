package mhw

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadArmors(t *testing.T) {
	dm := newDataManager()
	dm.loadSlotCombinations()
	dm.loadSkills()
	dm.loadArmorSetBonuses()
	dm.loadArmors()

	// assert.Equal(t, 359, len(dm.armors), "mismatch count of armors")
	armorName := "戰紋頭盔α"
	armorComponentName := "頭"
	armorComponentId := armorComponentNameToComponentId(armorComponentName)
	assert.NotZero(t, armorComponentId)
	armorResistences := []int{1, 1, -3, 1, -3}
	armorDecorationSlotCombinationId := 1
	armorSkillEnhancements := []skillEnhancement{
		{dm.getSkillIdByName("精神抖擻"), 2},
		{dm.getSkillIdByName("攻擊"), 1},
	}

	armorId := dm.getArmorIdByName(armorName)
	assert.NotZero(t, armorId)
	armor := dm.getArmorById(armorId)
	assert.NotNil(t, armor)
	assert.Equal(t, armorName, armor.name)
	assert.Equal(t, armorComponentId, armor.component)
	assert.Equal(t, armorResistences[:], armor.resistences[:])
	assert.Equal(t, armorDecorationSlotCombinationId, armor.slotCombination.id)
	assert.Equal(t, armorSkillEnhancements[:], armor.skillEnhancements[:])

	for _, a := range dm.armors[1:] {
		assert.NotEmpty(t, a.name)
		assert.NotZero(t, a.component)
		for _, se := range a.skillEnhancements {
			if se.skillId > 0 {
				s := dm.getSkillById(se.skillId)
				assert.True(t, se.enhancedLevel > 0)
				assert.True(t, se.enhancedLevel <= s.maxLevel)
			}
		}
	}
}
