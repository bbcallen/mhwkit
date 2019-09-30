package mhw

import (
	"errors"
	"fmt"
)

/*
 constant data manager
*/
type dataManager struct {
	skills           []skill
	decorations      []decoration
	armorSetBonuses  []armorSetBonus
	armors           []armor
	charms           []charm
	slotCombinations []slotCombination

	skillNameToSkillIdMap                 map[string]int
	decorationNameToDecorationIdMap       map[string]int
	armorSetBonusNameToArmorSetBonusIdMap map[string]int
	armorNameToArmorIdMap                 map[string]int
	charmNameToCharmIdMap                 map[string]int

	skillIdToDecorationIdMap map[int]int

	armorIdListsByComponent [armorComponentCount][]int

	slotCombinationOverridingMatrix             [][]int
	slotCombinationLayoutToSlotCombinationIdMap [5][5][5]int
}

func newDataManager() *dataManager {
	dm := &dataManager{}

	dm.skillNameToSkillIdMap = make(map[string]int)
	dm.decorationNameToDecorationIdMap = make(map[string]int)
	dm.armorSetBonusNameToArmorSetBonusIdMap = make(map[string]int)
	dm.armorNameToArmorIdMap = make(map[string]int)
	dm.charmNameToCharmIdMap = make(map[string]int)

	dm.skillIdToDecorationIdMap = make(map[int]int)

	return dm
}

func loadDataManager() *dataManager {
	dm := newDataManager()
	dm.loadSlotCombinations()
	dm.loadSkills()
	dm.loadArmorSetBonuses()
	dm.loadArmors()
	dm.loadDecorations()
	dm.loadCharms()
	return dm
}

func (dm *dataManager) registerSkill(s skill) {
	if _, ok := dm.skillNameToSkillIdMap[s.name]; ok {
		panic(errors.New(fmt.Sprintf("try to register duplicated skill %v", s.name)))
	}
	s.id = len(dm.skills)
	dm.skills = append(dm.skills, s)
	dm.skillNameToSkillIdMap[s.name] = s.id
}

func (dm *dataManager) registerArmorSet(b armorSetBonus) {
	if _, ok := dm.armorSetBonusNameToArmorSetBonusIdMap[b.name]; ok {
		panic(errors.New(fmt.Sprintf("try to register duplicated armor set %v", b.name)))
	}
	b.id = len(dm.armorSetBonuses)
	dm.armorSetBonuses = append(dm.armorSetBonuses, b)
	dm.armorSetBonusNameToArmorSetBonusIdMap[b.name] = b.id
}

func (dm *dataManager) registerArmor(a armor) {
	if _, ok := dm.armorNameToArmorIdMap[a.name]; ok {
		panic(errors.New(fmt.Sprintf("try to register duplicated armor %v", a.name)))
	}
	a.id = len(dm.armors)
	dm.armors = append(dm.armors, a)
	dm.armorIdListsByComponent[a.component] = append(dm.armorIdListsByComponent[a.component], a.id)
	dm.armorNameToArmorIdMap[a.name] = a.id
}

func (dm *dataManager) registerDecoration(d decoration) {
	if _, ok := dm.decorationNameToDecorationIdMap[d.name]; ok {
		panic(errors.New(fmt.Sprintf("try to register duplicated decoration %v", d.name)))
	}
	d.id = len(dm.decorations)
	dm.decorations = append(dm.decorations, d)
	dm.decorationNameToDecorationIdMap[d.name] = d.id
	if d.enhancedLevel <= 1 {
		if _, ok := dm.skillIdToDecorationIdMap[d.skillId]; ok {
			s := dm.skills[d.skillId]
			panic(errors.New(fmt.Sprintf("try to register decoration %v with duplicated skill %v", d.name, s.name)))
		}
		dm.skillIdToDecorationIdMap[d.skillId] = d.id
	}
}

func (dm *dataManager) registerCharm(c charm) {
	if _, ok := dm.charmNameToCharmIdMap[c.name]; ok {
		panic(errors.New(fmt.Sprintf("try to register duplicated charm %v", c.name)))
	}
	c.id = len(dm.charms)
	dm.charms = append(dm.charms, c)
	dm.charmNameToCharmIdMap[c.name] = c.id

	// register charms as armor
	if c.maxLevel > 0 {
		armor := newArmor()
		armor.name = c.name
		armor.component = charmAsArmor
		for i, skillId := range c.enhancedSkillIds {
			armor.skillEnhancements[i].skillId = skillId
			armor.skillEnhancements[i].enhancedLevel = c.maxLevel
		}
		dm.registerArmor(*armor)
	}
}

func (dm *dataManager) registerSlotCombination(sc slotCombination) {
	sc.id = len(dm.slotCombinations)
	dm.slotCombinations = append(dm.slotCombinations, sc)
	dm.slotCombinationLayoutToSlotCombinationIdMap[sc.layout[0]][sc.layout[1]][sc.layout[2]] = sc.id
}

func (dm *dataManager) getSkillById(id int) *skill {
	return &dm.skills[id]
}

func (dm *dataManager) getSkillIdByName(name string) (id int) {
	if id, ok := dm.skillNameToSkillIdMap[name]; ok {
		return id
	} else {
		panic(errors.New(fmt.Sprintf("can't find skill %v", name)))
	}
}

func (dm *dataManager) getDecorationById(id int) *decoration {
	return &dm.decorations[id]
}

func (dm *dataManager) getDecorationIdByName(name string) (id int) {
	if id, ok := dm.decorationNameToDecorationIdMap[name]; ok {
		return id
	} else {
		panic(errors.New(fmt.Sprintf("can't find decoration %v", name)))
	}
}

func (dm *dataManager) lookupDecorationIdBySkillId(skillId int) (decorationId int, ok bool) {
	decorationId, ok = dm.skillIdToDecorationIdMap[skillId]
	return decorationId, ok
}

func (dm *dataManager) getArmorById(id int) *armor {
	return &dm.armors[id]
}

func (dm *dataManager) getArmorIdByName(name string) (id int) {
	if id, ok := dm.armorNameToArmorIdMap[name]; ok {
		return id
	} else {
		panic(errors.New(fmt.Sprintf("can't find armor %v", name)))
	}
}

func (dm *dataManager) getArmorSetBonusById(id int) *armorSetBonus {
	return &dm.armorSetBonuses[id]
}

func (dm *dataManager) getArmorSetIdByName(name string) (id int) {
	if id, ok := dm.armorSetBonusNameToArmorSetBonusIdMap[name]; ok {
		return id
	} else {
		panic(errors.New(fmt.Sprintf("can't find armor %v", name)))
	}
}

func (dm *dataManager) getCharmById(id int) *charm {
	return &dm.charms[id]
}

func (dm *dataManager) getCharmIdByName(name string) (id int) {
	if id, ok := dm.charmNameToCharmIdMap[name]; ok {
		return id
	} else {
		panic(errors.New(fmt.Sprintf("can't find charm %v", name)))
	}
}

func (dm *dataManager) getSlotCombinationById(id int) *slotCombination {
	return &dm.slotCombinations[id]
}

func (dm *dataManager) getSlotCombinationIdByLayout(layout [3]int) int {
	return dm.slotCombinationLayoutToSlotCombinationIdMap[layout[0]][layout[1]][layout[2]]
}

/*
 could be one of:
	slotCombinationCannotOverride
	slotCombinationOverrideLeft
	slotCombinationOverrideRight
	slotCombinationOverrideEachOther
*/
func (dm *dataManager) getSlotCombinationIdsOverrideRelation(leftId int, rightId int) int {
	return dm.slotCombinationOverridingMatrix[leftId][rightId]
}
