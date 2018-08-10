package mhw

type armorSetBonus struct {
	id int

	name                              string
	activatedSkillIdsByComponentCount [armorComponentCount]int // skillId
}

func newArmorSetBonus() *armorSetBonus {
	asb := armorSetBonus{}
	return &asb
}
