package mhw

import (
    "testing"
)

// 长枪 炎妃冥灯
func testExecuteProcedureLance1(t *testing.T) {
    cc := newConstraintCollection(loadDataManager())
    addCommonDecorationLimitations(cc)
        
    cc.addExtraSlot(3, 2)
    cc.addRequiredSkillByName("體力增強", 3)  // <= 3
    cc.addRequiredSkillByName("攻擊", 4)    // <= 7
    cc.addRequiredSkillByName("防禦性能", 3)    // <= 5

    cc.addRequiredSkillByName("匠", 1) // <= 1

    cc.addRequiredSkillByName("弱點特效", 2) // <= 3
    cc.addRequiredSkillByName("看破", 6)   // <= 7
    cc.addRequiredSkillByName("超會心", 3) // <= 3

    cc.addRequiredSkillByName("挑戰者", 0)     // <= 5
    cc.addRequiredSkillByName("達人藝", 1) // <= 3
    cc.addRequiredSkillByName("昏厥耐性", 3) // <= 3

    proc := prepareProcedure(*cc)
    proc.execute()
}

// 炎妃冥灯长枪 耐火
func testExecuteProcedureLance2(t *testing.T) {
    cc := newConstraintCollection(loadDataManager())
    addCommonDecorationLimitations(cc)
        
    cc.addExtraSlot(3, 2)
    cc.addRequiredSkillByName("體力增強", 3)  // <= 3
    cc.addRequiredSkillByName("攻擊", 4)    // <= 7
    cc.addRequiredSkillByName("防禦性能", 3)    // <= 5

    cc.addRequiredSkillByName("匠", 0) // <= 1

    cc.addRequiredSkillByName("弱點特效", 2) // <= 3
    cc.addRequiredSkillByName("看破", 4)   // <= 7
    cc.addRequiredSkillByName("超會心", 3) // <= 3

    cc.addRequiredSkillByName("挑戰者", 0)     // <= 5
    cc.addRequiredSkillByName("達人藝", 1) // <= 3
    cc.addRequiredSkillByName("昏厥耐性", 3) // <= 3
    cc.addRequiredSkillByName("火耐性", 3) // <= 3

    proc := prepareProcedure(*cc)
    proc.execute()
}


// 炎妃冥灯长枪 极贝
func testExecuteProcedureLance3(t *testing.T) {
    cc := newConstraintCollection(loadDataManager())
    addCommonDecorationLimitations(cc)
        
    cc.addExtraSlot(2, 2)
    cc.addExtraSlot(1, 2)

    cc.addRequiredSkillByName("體力增強", 3)  // <= 3
    // cc.addRequiredSkillByName("攻擊", 4)    // <= 7
    cc.addRequiredSkillByName("防禦性能", 3)    // <= 5
    cc.addRequiredSkillByName("防禦強化", 1)    // <= 1
    // cc.addRequiredSkillByName("心眼／彈道強化", 1) // <= 1
    cc.addRequiredSkillByName("匠", 3) // <= 5

    cc.addRequiredSkillByName("昏厥耐性", 3) // <= 3
    cc.addRequiredSkillByName("體力回復量UP", 3) // <= 3
    cc.addRequiredSkillByName("回復速度", 3) // <= 3
    cc.addRequiredSkillByName("精靈加護", 3) // <= 3

    cc.addRequiredSkillByName("耳塞", 5) // <= 5
    // cc.addRequiredSkillByName("突破耐力上限", 1) // <= 1
    // cc.addRequiredSkillByName("廣域化", 2) // <= 5
    // cc.addRequiredSkillByName("快吃", 1)    // <= 3
    // cc.addRequiredSkillByName("龍屬性攻擊強化", 1) // <= 5

    proc := prepareProcedure(*cc)
    proc.execute()
}


// 铠罗惨爪
func testExecuteProcedureLance4(t *testing.T) {
    cc := newConstraintCollection(loadDataManager())
    addCommonDecorationLimitations(cc)
        
    // cc.addExtraSlot(3, 2)
    cc.addRequiredSkillByName("體力增強", 3)  // <= 3
    // cc.addRequiredSkillByName("攻擊", 4)    // <= 7
    cc.addRequiredSkillByName("防禦性能", 3)    // <= 5
    cc.addRequiredSkillByName("防禦強化", 1)    // <= 1
    cc.addRequiredSkillByName("無屬性強化", 1) // <= 1

    cc.addRequiredSkillByName("匠", 2) // <= 1

    cc.addRequiredSkillByName("弱點特效", 3) // <= 3
    cc.addRequiredSkillByName("看破", 4)   // <= 7
    cc.addRequiredSkillByName("超會心", 2) // <= 3

    cc.addRequiredSkillByName("精靈加護", 3)    // <= 3
    cc.addRequiredSkillByName("減輕膽怯", 1)    // <= 3

    // cc.addRequiredSkillByName("挑戰者", 0)     // <= 5
    // cc.addRequiredSkillByName("達人藝", 1) // <= 3
    // cc.addRequiredSkillByName("昏厥耐性", 3) // <= 3

    proc := prepareProcedure(*cc)
    proc.execute()
}


// 火属性
func testExecuteProcedureLance5(t *testing.T) {
    cc := newConstraintCollection(loadDataManager())
    addCommonDecorationLimitations(cc)
        
    cc.addExtraSlot(1, 1)
    cc.addRequiredSkillByName("體力增強", 3)  // <= 3
    // cc.addRequiredSkillByName("攻擊", 4)    // <= 7
    cc.addRequiredSkillByName("防禦性能", 3)    // <= 5
    cc.addRequiredSkillByName("防禦強化", 1)    // <= 1
    cc.addRequiredSkillByName("火屬性攻擊強化", 3) // <= 5

    cc.addRequiredSkillByName("匠", 1) // <= 5

    cc.addRequiredSkillByName("弱點特效", 3) // <= 3
    cc.addRequiredSkillByName("看破", 4)   // <= 7
    cc.addRequiredSkillByName("超會心", 2) // <= 3

    cc.addRequiredSkillByName("精靈加護", 3)    // <= 3
    cc.addRequiredSkillByName("減輕膽怯", 1)    // <= 3

    cc.addRequiredSkillByName("攻擊", 1)    // <= 7
    cc.addRequiredSkillByName("廣域化", 2)    // <= 5
    cc.addRequiredSkillByName("飛燕【屬性】", 1)    // <= 1

    // cc.addRequiredSkillByName("挑戰者", 0)     // <= 5
    // cc.addRequiredSkillByName("達人藝", 1) // <= 3
    // cc.addRequiredSkillByName("昏厥耐性", 3) // <= 3

    proc := prepareProcedure(*cc)
    proc.execute()
}

func testExecuteProcedureLance6(t *testing.T) {
    cc := newConstraintCollection(loadDataManager())
    addCommonDecorationLimitations(cc)
        
    cc.addExtraSlot(3, 2)
    cc.addRequiredSkillByName("體力增強", 3)  // <= 3
    cc.addRequiredSkillByName("攻擊", 4)    // <= 7
    cc.addRequiredSkillByName("防禦性能", 5)    // <= 5
    cc.addRequiredSkillByName("防禦強化", 1)    // <= 1

    cc.addRequiredSkillByName("匠", 1) // <= 1

    cc.addRequiredSkillByName("弱點特效", 3) // <= 3
    cc.addRequiredSkillByName("看破", 5)   // <= 7
    cc.addRequiredSkillByName("超會心", 0) // <= 3

    // cc.addRequiredSkillByName("挑戰者", 0)     // <= 5
    // cc.addRequiredSkillByName("達人藝", 1) // <= 3
    cc.addRequiredSkillByName("昏厥耐性", 3) // <= 3

    cc.addRequiredSkillByName("廣域化", 2) // <= 5
    cc.addRequiredSkillByName("屬性解放／裝填擴充", 1) // <= 3

    proc := prepareProcedure(*cc)
    proc.execute()
}
