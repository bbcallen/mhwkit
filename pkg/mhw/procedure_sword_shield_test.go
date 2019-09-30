package mhw

import (
    "testing"
)

// 奶片
func testExecuteProcedureSwordShield1(t *testing.T) {
    cc := newConstraintCollection(loadDataManager())
    cc.addExtraSlot(3, 1)
    cc.addExtraSlot(1, 1)

    addCommonDecorationLimitations(cc)

    cc.addRequiredSkillByName("廣域化", 5) // <= 5
    cc.addRequiredSkillByName("快吃", 3)    // <= 3
    cc.addRequiredSkillByName("體力增強", 3)  // <= 3

    cc.addRequiredSkillByName("最愛菇類", 1)   // <= 3
    cc.addRequiredSkillByName("耳塞", 5) // <= 5
    cc.addRequiredSkillByName("滿足感", 1)    // <= 1
    cc.addRequiredSkillByName("耐震", 3)    // <= 3

    cc.addRequiredSkillByName("屬性解放／裝填擴充", 3)    // <= 3
    // cc.addRequiredSkillByName("騎乘名人", 1)    // <= 1
    cc.addRequiredSkillByName("精靈加護", 3)    // <= 3

    // cc.addRequiredSkillByName("跳躍鐵人", 1)    // <= 1

    // cc.addRequiredSkillByName("突破耐力上限", 1)    // <= 1
    // cc.addRequiredSkillByName("心眼／彈道強化", 1)    // <= 1
 
    proc := prepareProcedure(*cc)
    proc.execute()
}
