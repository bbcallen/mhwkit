package mhw

const (
	armorComponentCount = 7
	nilArmor            = 0
	headArmor           = 1
	chestArmor          = 2
	armsArmor           = 3
	waistArmor          = 4
	legArmor            = 5
	charmAsArmor        = 6

	maxArmorEnhancedSkillCount = 2

	resistenceTypeCount = 5
	fileResistence      = 0
	waterResistence     = 1
	thunderResistence   = 2
	iceResistence       = 3
	dragonResistence    = 4
)

var (
	armorComponentNames []string = []string{"无", "頭", "身", "腕", "腰", "腳", "護石"}
	resistenceNames     []string = []string{"火", "水", "雷", "冰", "龙"}
)

type armor struct {
	id int

	name              string
	component         int
	setBonusId        int
	resistences       [resistenceTypeCount]int
	skillEnhancements [maxArmorEnhancedSkillCount]skillEnhancement

	slotCombination slotCombination
}

func newArmor() *armor {
	a := armor{}
	return &a
}

func (a *armor) decorationSlotCount() int {
	return a.slotCombination.slotCount()
}

func (a *armor) hasAnyDecorationSlot() bool {
	return a.slotCombination.hasAnySlot()
}

func (a *armor) doesEnhanceSkillId(skillId int) bool {
	for _, se := range a.skillEnhancements {
		if se.skillId == skillId {
			return true
		}
	}
	return false
}

func resistenceIdToName(id int) (name string) {
	return resistenceNames[id]
}

func armorComponentIdToComponentName(id int) (name string) {
	return armorComponentNames[id]
}

func armorComponentNameToComponentId(name string) (id int) {
	for i, n := range armorComponentNames {
		if n == name {
			return i
		}
	}
	return 0
}
