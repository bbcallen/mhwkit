package mhw

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadArmorSetBonuses(t *testing.T) {
	dm := newDataManager()
	dm.loadSkills()
	dm.loadArmorSetBonuses()

	assert.Equal(t, 19, len(dm.armorSetBonuses), "mismatch count of armor set bonuses")
}
