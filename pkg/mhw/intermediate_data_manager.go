package mhw

import (
	"errors"
	"fmt"
	"sort"
)

var (
	sizeToMatchingScorePerSkillLevel []int = []int{
		10000 * 10000, // no decoration
		10000,         // 1
		10500,         // 2
		10505,         // 3
	}
)

/*
 constraints related data manager
*/
type intermediateSkill struct {
	skill
	requiredLevel            int
	decorationId             int
	decorationSize           int
	availableDecorationCount int
}

func newIntermediateSkill() *intermediateSkill {
	skill := &intermediateSkill{}
	return skill
}

func (skill *intermediateSkill) isRequired() bool {
	return skill.requiredLevel > 0
}

func (skill *intermediateSkill) matchingScorePerLevel() int {
	if !skill.isRequired() {
		return 0
	}
	if skill.availableDecorationCount < skill.requiredLevel {
		return sizeToMatchingScorePerSkillLevel[0]
	}
	return sizeToMatchingScorePerSkillLevel[skill.decorationSize]
}

type intermediateArmor struct {
	armor
	doesEnhanceRequiredSkill bool
	mayHaveEquivalentArmors  bool
	matchingScore            int
}

func newIntermediateArmor() *intermediateArmor {
	armor := &intermediateArmor{}
	return armor
}

type intermediateComponentDimension struct {
	componentId                                 int
	slotCombinationIdToEquivalentArmorIdListMap map[int][]int
	candidateArmorIdList                        []int
	maxMatchingScore                            int
	necessaryMatchingScore                      int
}

func newIntermediateComponentDimension() *intermediateComponentDimension {
	component := &intermediateComponentDimension{}
	component.slotCombinationIdToEquivalentArmorIdListMap = make(map[int][]int)
	component.candidateArmorIdList = component.candidateArmorIdList[:]
	component.candidateArmorIdList = append(component.candidateArmorIdList, 0)
	return component
}

//----------------------------------------

type intermediateStatistic struct {
	armorIdListPrunedDueToNoHelp        []int
	armorIdListPrunedDueToEquivalent    []int
	armorIdListPrunedDueToMatchingScore []int

	armorIdListHaveRequiredSkill []int

	armorCombinationPrunedCounterByMatchingScore int
	armorCombinationBacktracingCounterByDepth [armorComponentCount]int

	armorCombinationEvaluationCounter int
	armorChangingCounter              int

	basicSolutionCounter int
}

type intermediateDataManager struct {
	dataManager          *dataManager
	constraintCollection constraintCollection

	armors []intermediateArmor
	skills []intermediateSkill

	componentDimensions         [armorComponentCount]intermediateComponentDimension
	totalRequisiteMatchingScore int

	statistic intermediateStatistic
}

func newIntermediateDataManager(dm *dataManager) *intermediateDataManager {
	idm := &intermediateDataManager{}
	idm.dataManager = dm
	return idm
}

func loadIntermediateDataManager(dm *dataManager) *intermediateDataManager {
	idm := newIntermediateDataManager(dm)
	idm.reloadData()
	return idm
}

func (idm *intermediateDataManager) getSkillById(id int) *intermediateSkill {
	return &idm.skills[id]
}

func (idm *intermediateDataManager) getArmorById(id int) *intermediateArmor {
	return &idm.armors[id]
}

func (idm *intermediateDataManager) getArmorMatchingScoreByArmorId(id int) int {
	return idm.armors[id].matchingScore
}

func (idm *intermediateDataManager) requisiteSkillEnhancements() []skillEnhancement {
	// TODO: reserve extra skills
	return idm.constraintCollection.requiredSkillEnhancements
}

func (idm *intermediateDataManager) extraSlots() [3]int {
	var extraSlots [3]int
	copy(extraSlots[:], idm.constraintCollection.extraSlotCountGroupBySize[:])
	return extraSlots
}

func (idm *intermediateDataManager) getEquivalentArmorIdListByArmorId(id int) []int {
	a := idm.getArmorById(id)
	if a.doesEnhanceRequiredSkill {
		return nil
	}
	return idm.componentDimensions[a.component].slotCombinationIdToEquivalentArmorIdListMap[a.slotCombination.id]
}

func (idm *intermediateDataManager) getArmorListByComponentId(id int) []int {
	return idm.componentDimensions[id].candidateArmorIdList
}

func (idm *intermediateDataManager) reloadData() {
	dm := idm.dataManager

	idm.skills = nil
	for _, skill := range dm.skills {
		s := newIntermediateSkill()
		s.skill = skill
		if decorationId, ok := dm.lookupDecorationIdBySkillId(skill.id); ok {
			d := dm.getDecorationById(decorationId)
			s.decorationId = decorationId
			s.decorationSize = d.size
			s.availableDecorationCount = s.maxLevel
		}
		idm.skills = append(idm.skills, *s)
	}

	idm.armors = nil
	for _, armor := range dm.armors {
		a := newIntermediateArmor()
		a.armor = armor
		idm.armors = append(idm.armors, *a)
	}

	for i, _ := range idm.componentDimensions {
		idm.componentDimensions[i] = *newIntermediateComponentDimension()
		idm.componentDimensions[i].componentId = i
	}

	idm.totalRequisiteMatchingScore = 0
	idm.statistic = intermediateStatistic{}
}

func (idm *intermediateDataManager) evaluateConstraints(cc constraintCollection) {
	idm.constraintCollection = cc
	idm.reloadData()
	idm.evaluateRequiredSkillEnhancements()
	idm.evaluateDecorationLimitations()
	idm.evaluateCandidateArmorList()
}

func (idm *intermediateDataManager) evaluateRequiredSkillEnhancements() {
	dm := idm.dataManager
	cc := idm.constraintCollection

	for i, _ := range idm.skills {
		idm.skills[i].requiredLevel = 0
	}

	for _, se := range cc.requiredSkillEnhancements {
		s := idm.getSkillById(se.skillId)
		s.requiredLevel = se.enhancedLevel
		idm.totalRequisiteMatchingScore += se.enhancedLevel * s.matchingScorePerLevel()
	}

	for i, count := range cc.extraSlotCountGroupBySize {
		idm.totalRequisiteMatchingScore -= count * sizeToMatchingScorePerSkillLevel[i+1]
	}

	for i, _ := range idm.armors {
		a := &idm.armors[i]
		for _, se := range a.skillEnhancements {
			if se.skillId > 0 {
				skill := idm.getSkillById(se.skillId)
				if skill.isRequired() {
					a.doesEnhanceRequiredSkill = true
					a.matchingScore += min(se.enhancedLevel, skill.requiredLevel) * skill.matchingScorePerLevel()
				}
			}
		}
		if a.setBonusId > 0 {
			armorSetBonus := dm.getArmorSetBonusById(a.setBonusId)
			for componentCount, skillId := range armorSetBonus.activatedSkillIdsByComponentCount {
				if skillId > 0 {
					skill := idm.getSkillById(skillId)
					if skill.isRequired() {
						a.doesEnhanceRequiredSkill = true
						a.matchingScore += skill.matchingScorePerLevel() / (componentCount - 1)
					}
				}
			}
		}

		for i, count := range a.slotCombination.sizeDistribution {
			a.matchingScore += count * sizeToMatchingScorePerSkillLevel[i+1]
		}

		if a.doesEnhanceRequiredSkill {
			idm.componentDimensions[a.component].maxMatchingScore = max(idm.componentDimensions[a.component].maxMatchingScore, a.matchingScore)
		} else if a.hasAnyDecorationSlot() {
			a.mayHaveEquivalentArmors = true
			eqMap := &idm.componentDimensions[a.component].slotCombinationIdToEquivalentArmorIdListMap
			(*eqMap)[a.slotCombination.id] = append((*eqMap)[a.slotCombination.id], a.id)
		}
	}

	maxTotalMatchingScore := 0
	for _, dim := range idm.componentDimensions {
		maxTotalMatchingScore += dim.maxMatchingScore
	}
	for i, _ := range idm.componentDimensions {
		dim := &idm.componentDimensions[i]
		maxOthersMatchingScore := maxTotalMatchingScore - dim.maxMatchingScore
		dim.necessaryMatchingScore = max(0, idm.totalRequisiteMatchingScore-maxOthersMatchingScore)
	}
}

func (idm *intermediateDataManager) evaluateDecorationLimitations() {
	dm := idm.dataManager
	cc := idm.constraintCollection

	for _, dl := range cc.decorationLimitations {
		d := dm.getDecorationById(dl.decorationId)
		s := idm.getSkillById(d.skillId)
		s.availableDecorationCount = min(s.availableDecorationCount, dl.maxCount)
	}
}

func (idm *intermediateDataManager) evaluateCandidateArmorList() {
	dm := idm.dataManager
	for componentId, armorIdList := range dm.armorIdListsByComponent {
		dim := &idm.componentDimensions[componentId]
		for _, armorId := range armorIdList {
			a := idm.getArmorById(armorId)
			if a.doesEnhanceRequiredSkill {
				idm.statistic.armorIdListHaveRequiredSkill = append(idm.statistic.armorIdListHaveRequiredSkill, a.id)
				// do nothing here
			} else if a.hasAnyDecorationSlot() {
				eqMap := dim.slotCombinationIdToEquivalentArmorIdListMap
				eqList := eqMap[a.slotCombination.id]
				if len(eqList) <= 0 {
					panic(errors.New(fmt.Sprintf("unexpected empty eq list of %v", a.name)))
				}

				if eqList[0] == armorId {
					// do nothing here
				} else {
					idm.statistic.armorIdListPrunedDueToEquivalent = append(idm.statistic.armorIdListPrunedDueToEquivalent, a.id)
					continue
				}
			} else {
				idm.statistic.armorIdListPrunedDueToNoHelp = append(idm.statistic.armorIdListPrunedDueToNoHelp, a.id)
				continue
			}

			if a.matchingScore < dim.necessaryMatchingScore {
				idm.statistic.armorIdListPrunedDueToMatchingScore = append(idm.statistic.armorIdListPrunedDueToMatchingScore, a.id)
				continue
			}

			dim.candidateArmorIdList = append(dim.candidateArmorIdList, armorId)
		}

		sort.Slice(
			dim.candidateArmorIdList,
			func(i, j int) bool {
				lhs := idm.getArmorById(dim.candidateArmorIdList[i])
				rhs := idm.getArmorById(dim.candidateArmorIdList[j])
				if lhs.doesEnhanceRequiredSkill && rhs.doesEnhanceRequiredSkill {
					return lhs.matchingScore > rhs.matchingScore
				} else if lhs.doesEnhanceRequiredSkill {
					return true
				} else if rhs.doesEnhanceRequiredSkill {
					return false
				} else {
					return lhs.matchingScore > rhs.matchingScore
				}
			})
	}
}

func (idm *intermediateDataManager) printCandidateArmors() {
	for i, _ := range idm.componentDimensions {
		dim := &idm.componentDimensions[i]
		fmt.Printf("%v: %v\n", armorComponentIdToComponentName(i), dim.necessaryMatchingScore)
		for _, id := range dim.candidateArmorIdList {
			armor := idm.getArmorById(id)
			fmt.Printf("  %v [%v]=%v %v %v\n",
				armor.doesEnhanceRequiredSkill,
				armor.slotCombination.id,
				armor.slotCombination.layout,
				armor.matchingScore,
				armor.name)
		}
	}
}
