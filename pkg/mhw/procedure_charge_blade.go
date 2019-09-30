package mhw

import (
    "testing"
)

// 铠罗兵装角
func testExecuteProcedureChargeBlade1(t *testing.T) {
    cc := newConstraintCollection(loadDataManager())
    addCommonDecorationLimitations(cc)

    cc.addRequiredSkillByName("體力增強", 3)  // <= 3

    cc.addRequiredSkillByName("無屬性強化", 1) // <= 1
    cc.addRequiredSkillByName("集中", 3)   // <= 3
    cc.addRequiredSkillByName("砲術", 3)    // <= 3
    cc.addRequiredSkillByName("砲彈裝填數UP", 1)    // <= 1

    cc.addRequiredSkillByName("弱點特效", 3) // <= 3
    cc.addRequiredSkillByName("匠", 3)    // <= 5
    cc.addRequiredSkillByName("攻擊", 4)    // <= 7
    cc.addRequiredSkillByName("看破", 4)   // <= 7

    cc.addRequiredSkillByName("廣域化", 2) // <= 5
 
    proc := prepareProcedure(*cc)
    proc.execute()
}
