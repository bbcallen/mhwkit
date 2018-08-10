package mhw

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadSkills(t *testing.T) {
	dm := newDataManager()
	dm.loadSkills()

	assert.Equal(t, 136, len(dm.skills), "mismatch count of skills")
	skillName := "體力增強"
	skillMaxLevel := 3

	skillId := dm.getSkillIdByName(skillName)
	assert.NotZero(t, skillId)
	skill := dm.getSkillById(skillId)
	assert.NotNil(t, skill)
	assert.Equal(t, skillName, skill.name)
	assert.Equal(t, skillMaxLevel, skill.maxLevel)

	for _, s := range dm.skills[1:] {
		assert.NotEmpty(t, s.name)
		assert.NotZero(t, s.maxLevel)
	}
}
