package mhw

const (
	maxCharmEnhancedSkillCount = 2
)

type charm struct {
	id int

	name             string
	rarity           int
	maxLevel         int
	enhancedSkillIds [maxCharmEnhancedSkillCount]int
}

func newCharm() *charm {
	d := charm{}
	return &d
}
