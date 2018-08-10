package mhw

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadIntermediateDataManager(t *testing.T) {
	dm := loadDataManager()
	idm := loadIntermediateDataManager(dm)
	if testing.Verbose() {
		fmt.Println("intermediateDataManager:", idm)
	}
}

func TestEvaluateCandidateArmorList(t *testing.T) {
	dm := loadDataManager()
	idm := loadIntermediateDataManager(dm)
	cc := newConstraintCollection(dm)
	cc.addRequiredSkillByName("體力增強", 3)
	cc.addRequiredSkillByName("攻擊", 2)
	idm.evaluateConstraints(*cc)
	assert.NotEmpty(t, idm.componentDimensions[chestArmor].candidateArmorIdList)

	for _, dim := range idm.componentDimensions {
		for _, armorId := range dim.candidateArmorIdList {
			a := idm.getArmorById(armorId)
			switch {
			case a.id == 0:
			case a.doesEnhanceSkillId(dm.getSkillIdByName("體力增強")):
			case a.doesEnhanceSkillId(dm.getSkillIdByName("攻擊")):
			case a.hasAnyDecorationSlot():
			default:
				assert.FailNow(t, fmt.Sprintf("unexpected candidate armor %v", a))
			}
		}
	}

	if testing.Verbose() {
		idm.printCandidateArmors()
		fmt.Printf("pruned %v armors\n", len(idm.statistic.armorIdListPrunedDueToNoHelp))
		/*
			for _, armorId := range idm.statistic.armorIdListPrunedDueToNoHelp {
				a := idm.getArmorById(armorId)
				fmt.Printf("    %v\n", a.name)
			}
		*/
	}
}
