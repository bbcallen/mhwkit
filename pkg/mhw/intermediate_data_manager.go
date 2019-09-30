package mhw

import (
	"errors"
	"fmt"
	"sort"
)

var (
	sizeToMatchingScorePerSkillLevel []int = []int{
		10000000,       // no decoration for this skill, could be bonus skill
		10000010,       // 1
		10001000,       // 2
		10100000,       // 3
		11000000,       // 4
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
	matchingScore 			 int
}

func newIntermediateSkill() *intermediateSkill {
	skill := &intermediateSkill{}
	return skill
}

func (skill *intermediateSkill) isRequired() bool {
	return skill.requiredLevel > 0
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

	statistic 					intermediateStatistic

	extraEnhancedSkillIds  		map[int]int
}

func newIntermediateDataManager(dm *dataManager) *intermediateDataManager {
	idm := &intermediateDataManager{}
	idm.dataManager = dm
	idm.extraEnhancedSkillIds = make(map[int]int)
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

func (idm *intermediateDataManager) extraSlots() [4]int {
	var extraSlots [4]int
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
			s.matchingScore = sizeToMatchingScorePerSkillLevel[s.decorationSize]
		} else {
			s.matchingScore = sizeToMatchingScorePerSkillLevel[0]
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
		idm.totalRequisiteMatchingScore += se.enhancedLevel * s.matchingScore
	}

	for i, count := range cc.extraSlotCountGroupBySize {
		idm.totalRequisiteMatchingScore -= count * sizeToMatchingScorePerSkillLevel[i+1]
	}

	for i, _ := range idm.armors {
		a := &idm.armors[i]
		for _, se := range a.skillEnhancements {
			if se.skillId > 0 {
				s := idm.getSkillById(se.skillId)
				if s.isRequired() {
					a.doesEnhanceRequiredSkill = true
					a.matchingScore += min(se.enhancedLevel, s.requiredLevel) * s.matchingScore
					/*
					fmt.Printf("  skill %v: %v.score += min(%v, %v) * %v\n %v\n",
						s.skill.name,
						a.armor.name,
						se.enhancedLevel,
						s.requiredLevel,
						s.matchingScore,
						a.matchingScore)
						*/
				}
			}
		}
		if a.setBonusId > 0 {
			armorSetBonus := dm.getArmorSetBonusById(a.setBonusId)
			for componentCount, skillId := range armorSetBonus.activatedSkillIdsByComponentCount {
				if skillId > 0 {
					s := idm.getSkillById(skillId)
					if s.isRequired() {
						a.doesEnhanceRequiredSkill = true
						a.matchingScore += s.matchingScore / componentCount + 1
						/*
						fmt.Printf("  bonus %v %v: %v.score += %v / %v => %v\n",
							armorSetBonus.name,
							s.skill.name,
							a.armor.name,
							s.matchingScore,
							componentCount,
							a.matchingScore)
							*/
					}
				}
			}
		}

		for i, count := range a.slotCombination.sizeDistribution {
			a.matchingScore += count * sizeToMatchingScorePerSkillLevel[i+1]
			/*
			fmt.Printf("  slot:  %v.score += %v * table[%v+1](%v) => %v\n",
				a.armor.name,
				count,
				i,
				sizeToMatchingScorePerSkillLevel[i+1],
				a.matchingScore)
				*/
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
