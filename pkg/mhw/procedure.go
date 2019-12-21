package mhw

import (
    "errors"
    "fmt"
)

type armorCombinationEvaluator struct {
    idm                                      *intermediateDataManager
    armorIds                                 [armorComponentCount]int
    enhancedSkillLevels                      []int
    armorSetBonusIdToActivatingArmorCountMap []int
    slotSizeDistribution                     [4]int // lv1, lv2, lv3, lv4
}

func newArmorCombinationEvaluatorFrom(idm *intermediateDataManager) *armorCombinationEvaluator {
    acEval := &armorCombinationEvaluator{}
    acEval.idm = idm
    acEval.idm = idm
    acEval.enhancedSkillLevels = make([]int, len(idm.skills))
    acEval.armorSetBonusIdToActivatingArmorCountMap = make([]int, len(idm.dataManager.armorSetBonuses))
    return acEval
}

func (acEval *armorCombinationEvaluator) evaluateSelectingArmorId(armorId int, selected bool) {
    idm := acEval.idm
    dm := idm.dataManager

    var sign int
    if selected {
        sign = 1
    } else {
        sign = -1
    }

    armor := idm.getArmorById(armorId)
    for _, es := range armor.skillEnhancements {
        if es.skillId > 0 {
            /*
            fmt.Printf("  skill %v: %v.score += %v * %v :: %v",
                dm.getSkillById(es.skillId).name,
                armor.armor.name,
                es.enhancedLevel,
                sign,
                acEval.enhancedSkillLevels[es.skillId])
                */
            acEval.enhancedSkillLevels[es.skillId] += es.enhancedLevel * sign
            /*
            fmt.Printf(" => %v\n",
                acEval.enhancedSkillLevels[es.skillId])
                */
        }
    }
    for i, count := range armor.slotCombination.sizeDistribution {
        acEval.slotSizeDistribution[i] += count * sign
    }
    if armor.setBonusId != 0 {
        armorSetBonus := dm.getArmorSetBonusById(armor.setBonusId)
        activatingArmorCount := acEval.armorSetBonusIdToActivatingArmorCountMap[armor.setBonusId]
        effectedSkillId := 0
        if selected {
            // fmt.Printf("activatingArmorCount++: %v\n", activatingArmorCount)
            activatingArmorCount++
            effectedSkillId = armorSetBonus.activatedSkillIdsByComponentCount[activatingArmorCount]
        } else {
            effectedSkillId = armorSetBonus.activatedSkillIdsByComponentCount[activatingArmorCount]
            // fmt.Printf("activatingArmorCount--: %v\n", activatingArmorCount)
            activatingArmorCount--
        }
        acEval.armorSetBonusIdToActivatingArmorCountMap[armor.setBonusId] = activatingArmorCount
        if effectedSkillId != 0 {
            // fmt.Printf("activate skill: %v[%v]\n", effectedSkillId, dm.getSkillById(effectedSkillId))
            acEval.enhancedSkillLevels[effectedSkillId] += sign
        }
    }

    idm.statistic.armorChangingCounter++
}


type solutionEvaluator struct {
    idm             *intermediateDataManager
    acEval          *armorCombinationEvaluator

    matchingScore          int
    armorIdsToBeEvaluated  [armorComponentCount]int
    doesEnhanceRequiredSkillFlags [armorComponentCount]bool

    availableSlotCountGroupBySize [4]int
    skillEnhancementsByDecoration []skillEnhancement
}

func newSolutionEvaluatorFrom(idm *intermediateDataManager) *solutionEvaluator {
    slnEval := &solutionEvaluator{}
    slnEval.idm = idm
    slnEval.acEval = newArmorCombinationEvaluatorFrom(idm)
    return slnEval
}

func (slnEval *solutionEvaluator) replaceWithArmorIdAtComponent(componentId int, armorId int) {
    idm := slnEval.idm
    newArmor := idm.getArmorById(armorId)
    previousArmorId := slnEval.armorIdsToBeEvaluated[componentId]
    if previousArmorId != armorId {
        previousArmor := idm.getArmorById(previousArmorId)
        slnEval.matchingScore -= previousArmor.matchingScore
        slnEval.matchingScore += newArmor.matchingScore
        slnEval.armorIdsToBeEvaluated[componentId] = armorId
        slnEval.doesEnhanceRequiredSkillFlags[componentId] = newArmor.doesEnhanceRequiredSkill
    }
}

func (slnEval *solutionEvaluator) hasEnoughMatchingScore() bool {
    idm := slnEval.idm
    return slnEval.matchingScore >= idm.totalRequisiteMatchingScore
}

func (slnEval *solutionEvaluator) evaluateArmorCombination() {
    // fmt.Printf("  精英・火龍 火屬性攻擊強化\n")
    idm := slnEval.idm
    acEval := slnEval.acEval
    for componentId, previousArmorId := range acEval.armorIds {
        newArmorId := slnEval.armorIdsToBeEvaluated[componentId]
        if newArmorId != previousArmorId {
            acEval.evaluateSelectingArmorId(previousArmorId, false)
            acEval.evaluateSelectingArmorId(newArmorId, true)
            acEval.armorIds[componentId] = newArmorId
        }
    }
    idm.statistic.armorCombinationEvaluationCounter++
}

func (slnEval *solutionEvaluator) evaluateDecorations() bool {
    idm := slnEval.idm
    acEval := slnEval.acEval

    slnEval.skillEnhancementsByDecoration = slnEval.skillEnhancementsByDecoration[:0]
    copy(slnEval.availableSlotCountGroupBySize[:], acEval.slotSizeDistribution[:])

    for i, count := range idm.extraSlots() {
        slnEval.availableSlotCountGroupBySize[i] += count
        if (slnEval.availableSlotCountGroupBySize[i] < 0) {
            return false
        } 
    }

nextSkill:
    for _, se := range idm.requisiteSkillEnhancements() {
        requisiteLevel := se.enhancedLevel - acEval.enhancedSkillLevels[se.skillId]
        if requisiteLevel <= 0 {
            // already meet the requirement
            continue nextSkill
        }

        skill := idm.getSkillById(se.skillId)
        if skill.availableDecorationCount < requisiteLevel {
            // no enough decorations
            return false
        }

        var skillEnhancementByDecoration skillEnhancement
        skillEnhancementByDecoration.skillId = skill.id
        for j, availableSlotCount := range slnEval.availableSlotCountGroupBySize[skill.decorationSize-1:] {
            usedDecorationCount := min(requisiteLevel, availableSlotCount)
            slnEval.availableSlotCountGroupBySize[skill.decorationSize-1+j] -= usedDecorationCount
            requisiteLevel -= usedDecorationCount
            skillEnhancementByDecoration.enhancedLevel += usedDecorationCount
            if requisiteLevel <= 0 {
                break
            }
        }
        if requisiteLevel > 0 {
            // no enough slots
            return false
        }

        slnEval.skillEnhancementsByDecoration = append(slnEval.skillEnhancementsByDecoration, skillEnhancementByDecoration)
    }

    return true
}

func (slnEval *solutionEvaluator) printSolution() {
    idm := slnEval.idm
    dm := idm.dataManager
    acEval := slnEval.acEval

    eqCount := 1
    for _, armorId := range acEval.armorIds {
        eqList := idm.getEquivalentArmorIdListByArmorId(armorId)
        count := len(eqList)
        if count > 1 {
            eqCount *= count
        }
    }

    fmt.Printf("Setup[%v]:\n", idm.statistic.basicSolutionCounter)
    totalScore := 0
    for i, armorId := range acEval.armorIds[1:] {
        componentId := i + 1
        armor := idm.getArmorById(armorId)
        totalScore += armor.matchingScore
        equivalentArmorIdList := idm.getEquivalentArmorIdListByArmorId(armorId)
        equivalentArmorIdCount := len(equivalentArmorIdList)
        if equivalentArmorIdCount > 1 {
            fmt.Printf("    %v: %v (%v equivalent): ",
                armorComponentIdToComponentName(i),
                armor.name,
                equivalentArmorIdCount-1)
            for _, equivalentArmorId := range equivalentArmorIdList[1:min(4, equivalentArmorIdCount)] {
                equivalentArmor := idm.getArmorById(equivalentArmorId)
                fmt.Printf("%v ", equivalentArmor.name)
            }
            fmt.Println("...")
        } else if armorId == 0 {
            fmt.Printf("    %v: <任意>\n",
                armorComponentIdToComponentName(componentId))
        } else {
            fmt.Printf("    %v[%v]: %v %v = %v\n",
                armorComponentIdToComponentName(componentId),
                armor.id,
                armor.name,
                armor.slotCombination.layout,
                armor.matchingScore)
        }
    }
    fmt.Printf("    score: %v\n", totalScore)
    // TODO: simplify
    totalSkillEnhancements := make([]int, len(idm.skills))
    for skillId, level := range acEval.enhancedSkillLevels {
        if level > 0 {
            totalSkillEnhancements[skillId] = level
            // fmt.Printf("    skill by armor: %v %v\n", dm.getSkillById(skillId).name, level)
        }
    }
    fmt.Printf("    ----\n")
    for _, se := range slnEval.skillEnhancementsByDecoration {
        if se.enhancedLevel > 0 {
            // fmt.Printf("    skill by deco: %v %v\n", dm.getSkillById(se.skillId).name, se.enhancedLevel)
            decorationId := idm.getSkillById(se.skillId).decorationId
            d := dm.getDecorationById(decorationId)
            totalSkillEnhancements[d.skillId] += se.enhancedLevel
            fmt.Printf("    %v %v\n", d.name, se.enhancedLevel)
        }
    }
    for i, count := range slnEval.availableSlotCountGroupBySize {
        if count > 0 {
            fmt.Printf("    鑲嵌槽[%v] %v\n", i+1, count)
        }
    }
    fmt.Printf("    ----\n")
    for skillId, level := range totalSkillEnhancements {
        if level <= 0 {
            continue
        }
        var prefix string
        skill := idm.getSkillById(skillId)
        if skill.isRequired() {
            prefix = "*"
        } else {
            prefix = " "
        }
        if level > skill.requiredLevel {
            idm.extraEnhancedSkillIds[skillId] = level
        }
        fmt.Printf("    %v %v %v/%v\n", prefix, skill.name, level, skill.maxLevel)
    }
}

func (slnEval *solutionEvaluator) printStatistic() {
    idm := slnEval.idm
    dm := idm.dataManager

    fmt.Printf("%v total armors\n", len(idm.armors))
    fmt.Printf("    have required skill: %v armor\n", len(idm.statistic.armorIdListHaveRequiredSkill))
    fmt.Printf("%v armor pruned:\n",
        len(idm.statistic.armorIdListPrunedDueToNoHelp)+
            len(idm.statistic.armorIdListPrunedDueToEquivalent)+
            len(idm.statistic.armorIdListPrunedDueToMatchingScore))
    fmt.Printf("    no help:    %v\n", len(idm.statistic.armorIdListPrunedDueToNoHelp))
    fmt.Printf("    equivalent: %v\n", len(idm.statistic.armorIdListPrunedDueToEquivalent))
    fmt.Printf("    low-score:  %v\n", len(idm.statistic.armorIdListPrunedDueToMatchingScore))
    armorCount := 0
    totalCombinations := 1
    for i, _ := range idm.componentDimensions[1:] {
        componentId := i + 1
        count := len(idm.getArmorListByComponentId(componentId))
        totalCombinations *= count
        armorCount += count
    }
    fmt.Printf("%v armors left after pruned\n", armorCount)
    for i, _ := range idm.componentDimensions[1:] {
        componentId := i + 1
        fmt.Printf("    %v %v/%v\n",
            armorComponentIdToComponentName(componentId),
            len(idm.getArmorListByComponentId(componentId)),
            len(dm.armorIdListsByComponent[componentId]))
    }
    fmt.Printf("%v armor combinatons expected\n", totalCombinations)
    fmt.Printf("%v armor combinatons pruned\n", idm.statistic.armorCombinationPrunedCounterByMatchingScore)
    fmt.Printf("armor combinaton trees backtracing %v\n", idm.statistic.armorCombinationBacktracingCounterByDepth)
    fmt.Printf("%v armor combinatons evaluated\n", idm.statistic.armorCombinationEvaluationCounter)
    fmt.Printf("%v armor select/deselect changed\n", idm.statistic.armorChangingCounter)
    fmt.Printf("%v basic solutions\n", idm.statistic.basicSolutionCounter)
    fmt.Printf("%v extra enhanced skills:\n", len(idm.extraEnhancedSkillIds))
    for skillId, level := range idm.extraEnhancedSkillIds {
        skill := idm.getSkillById(skillId)
        fmt.Printf("      %v %v/%v\n", skill.name, level, skill.maxLevel)
    }
}




type armorIdEnumerator struct {
    armorIdList []int
    index       int
}

func newArmorIdEnumerator() *armorIdEnumerator {
    aidEtor := &armorIdEnumerator{}
    return aidEtor
}

func beginEnumerateArmorIdList(armorIdList []int) *armorIdEnumerator {
    aidEtor := newArmorIdEnumerator()
    aidEtor.armorIdList = armorIdList
    aidEtor.index = 0
    return aidEtor
}

func (aidEtor *armorIdEnumerator) moveToEnd() {
    aidEtor.index = len(aidEtor.armorIdList)
}

func (aidEtor *armorIdEnumerator) hasValue() bool {
    return aidEtor.index < len(aidEtor.armorIdList)
}

func (aidEtor *armorIdEnumerator) findNext() bool {
    if !aidEtor.hasValue() {
        return false
    }

    aidEtor.index++
    return aidEtor.hasValue()
}

func (aidEtor *armorIdEnumerator) armorId() int {
    return aidEtor.armorIdList[aidEtor.index]
}

func (aidEtor *armorIdEnumerator) beginAgain() bool {
    aidEtor.index = 0
    return aidEtor.hasValue()
}

func (aidEtor *armorIdEnumerator) unenumerated() int {
    length := len(aidEtor.armorIdList)
    if length <= 0 {
        return 0
    }
    return length - aidEtor.index
}


type procedure struct {
    dataManager             *dataManager
    intermediateDataManager *intermediateDataManager
    constraintCollection    constraintCollection
}

func newProcedure() *procedure {
    proc := &procedure{}
    return proc
}

func prepareProcedure(cc constraintCollection) *procedure {
    proc := newProcedure()
    proc.constraintCollection = cc
    proc.dataManager = cc.dataManager
    idm := loadIntermediateDataManager(proc.dataManager)
    proc.intermediateDataManager = idm
    idm.evaluateConstraints(cc)
    return proc
}

func (proc *procedure) execute() {
    dm := proc.dataManager
    cc := proc.constraintCollection
    idm := proc.intermediateDataManager
    if idm == nil {
        idm := loadIntermediateDataManager(dm)
        proc.intermediateDataManager = idm
        idm.evaluateConstraints(cc)
    }

    // init evaluator and enumerators
    slnEval := newSolutionEvaluatorFrom(idm)
    aidEtors := [armorComponentCount]armorIdEnumerator{}
    for componentId, _ := range aidEtors {
        aidEtors[componentId] = *beginEnumerateArmorIdList(idm.getArmorListByComponentId(componentId))
        aidEtor := &aidEtors[componentId]
        if !aidEtor.hasValue() {
            panic(errors.New(fmt.Sprintf("empty candidate armor list for component[%v]: %v", componentId, aidEtor)))
        }

        slnEval.replaceWithArmorIdAtComponent(componentId, aidEtor.armorId())
        // fmt.Printf("aidEtor [%v]:%v\n", componentId, aidEtor)
    }

    // main enumerator loop
    fmt.Printf("score required: %v \n", idm.totalRequisiteMatchingScore)
nextCombination:
    for {
        needBacktracing := false
        if slnEval.hasEnoughMatchingScore() {
            slnEval.evaluateArmorCombination()
            if slnEval.evaluateDecorations() {
                idm.statistic.basicSolutionCounter++
                slnEval.printSolution()
            }
        } else {
            needBacktracing = true
            idm.statistic.armorCombinationPrunedCounterByMatchingScore++
        }
        for i, _ := range aidEtors[1:] {
            componentId := i + 1
            aidEtor := &aidEtors[componentId]
            if needBacktracing {
                armor := idm.getArmorById(aidEtor.armorId())
                if armor.doesEnhanceRequiredSkill {
                    needBacktracing = false
                } else {
                    idm.statistic.armorCombinationBacktracingCounterByDepth[componentId] += aidEtor.unenumerated()
                    aidEtor.beginAgain()
                    slnEval.replaceWithArmorIdAtComponent(componentId, aidEtor.armorId())
                    continue
                }
            }
            if aidEtor.findNext() {
                slnEval.replaceWithArmorIdAtComponent(componentId, aidEtor.armorId())
                continue nextCombination
            } else {
                aidEtor.beginAgain()
                slnEval.replaceWithArmorIdAtComponent(componentId, aidEtor.armorId())
            }
        }
        break
    }

    slnEval.printStatistic()
    fmt.Printf("\n")
}
