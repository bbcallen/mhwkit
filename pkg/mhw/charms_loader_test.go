package mhw

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadCharms(t *testing.T) {
	dm := newDataManager()
	dm.loadSkills()
	dm.loadCharms()

	assert.Equal(t, 103, len(dm.charms), "mismatch count of charms")
	skillName := "體力增強"
	charmName := "體力護石"
	charmMaxLevel := 3
	charmEnhancedSkillIds := []int{dm.getSkillIdByName("體力增強"), 0}

	charmId := dm.getCharmIdByName(charmName)
	assert.NotZero(t, charmId)
	skillId := dm.getSkillIdByName(skillName)
	assert.NotZero(t, skillId)
	charm := dm.getCharmById(charmId)
	assert.NotNil(t, charm)
	assert.Equal(t, charmName, charm.name)
	assert.Equal(t, charmMaxLevel, charm.maxLevel)
	assert.Equal(t, charmEnhancedSkillIds[:], charm.enhancedSkillIds[:])

	for _, c := range dm.charms[1:] {
		assert.NotEmpty(t, c.name)
		assert.NotZero(t, c.maxLevel)
		assert.NotEmpty(t, c.enhancedSkillIds)
	}
}
