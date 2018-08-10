package mhw

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadDecorations(t *testing.T) {
	dm := newDataManager()
	dm.loadSkills()
	dm.loadDecorations()

	assert.Equal(t, 99, len(dm.decorations), "mismatch count of decorations")
	skillName := "體力增強"
	decorationName := "體力珠【1】"
	decorationSize := 1
	decorationRarity := 6

	decorationId := dm.getDecorationIdByName(decorationName)
	assert.NotZero(t, decorationId)
	skillId := dm.getSkillIdByName(skillName)
	assert.NotZero(t, skillId)
	decorationId, ok := dm.lookupDecorationIdBySkillId(skillId)
	assert.True(t, ok)
	assert.NotZero(t, decorationId)
	decoration := dm.getDecorationById(decorationId)
	assert.NotNil(t, decoration)
	assert.Equal(t, decorationName, decoration.name)
	assert.Equal(t, decorationSize, decoration.size)
	assert.Equal(t, decorationRarity, decoration.rarity)
	assert.Equal(t, skillId, decoration.skillId)

	for _, d := range dm.decorations[1:] {
		assert.NotEmpty(t, d.name)
		assert.NotZero(t, d.size)
		assert.NotZero(t, d.rarity)
		assert.NotZero(t, d.skillId)
	}
}
