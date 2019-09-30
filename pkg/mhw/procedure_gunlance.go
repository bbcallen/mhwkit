package mhw

import (
    "testing"
)

// 铳枪
// 铠罗水 插射流
func testExecuteProcedureGunlance1(t *testing.T) {
    cc := newConstraintCollection(loadDataManager())
    addCommonDecorationLimitations(cc)
        
    cc.addExtraSlot(3, 1)
    cc.addRequiredSkillByName("體力增強", 3)  // <= 3
    cc.addRequiredSkillByName("砲術", 3)    // <= 3
    cc.addRequiredSkillByName("砲彈裝填數UP", 1)    // <= 1
    cc.addRequiredSkillByName("防禦性能", 3)    // <= 5
    cc.addRequiredSkillByName("防禦強化", 1)    // <= 1

    cc.addRequiredSkillByName("無屬性強化", 1) // <= 1
    cc.addRequiredSkillByName("匠", 4) // <= 5
    cc.addRequiredSkillByName("剛刃研磨", 1) // <= 1
    // cc.addRequiredSkillByName("利刃／彈藥節約", 1) // <= 1
    // cc.addRequiredSkillByName("砥石使用高速化", 3) // <= 3

    cc.addRequiredSkillByName("弱點特效", 3) // <= 3

    cc.addRequiredSkillByName("看破", 2)   // <= 7

    // cc.addRequiredSkillByName("超會心", 1) // <= 3
    // cc.addRequiredSkillByName("達人藝", 1) // <= 3

    // cc.addRequiredSkillByName("迴避距離UP", 1)   // <= 3
    // cc.addRequiredSkillByName("攻擊", 1)    // <= 7
    // cc.addRequiredSkillByName("挑戰者", 1)     // <= 5
    // cc.addRequiredSkillByName("廣域化", 2) // <= 5
    // cc.addRequiredSkillByName("達人藝", 1) // <= 3
    // cc.addRequiredSkillByName("昏厥耐性", 3) // <= 3
    proc := prepareProcedure(*cc)
    proc.execute()
}

// 铠罗水 贝爷T
func testExecuteProcedureGunlance2(t *testing.T) {
    cc := newConstraintCollection(loadDataManager())
    addCommonDecorationLimitations(cc)
        
    cc.addExtraSlot(3, 1)
    cc.addRequiredSkillByName("體力增強", 3)  // <= 3
    cc.addRequiredSkillByName("砲術", 3)    // <= 3
    cc.addRequiredSkillByName("砲彈裝填數UP", 1)    // <= 1
    cc.addRequiredSkillByName("防禦性能", 5)    // <= 5
    cc.addRequiredSkillByName("防禦強化", 1)    // <= 1

    cc.addRequiredSkillByName("耳塞", 5) // <= 5
    // cc.addRequiredSkillByName("無屬性強化", 1) // <= 1
    cc.addRequiredSkillByName("匠", 3) // <= 5
    // cc.addRequiredSkillByName("剛刃研磨", 1) // <= 1
    // cc.addRequiredSkillByName("利刃／彈藥節約", 1) // <= 1
    // cc.addRequiredSkillByName("砥石使用高速化", 3) // <= 3

    cc.addRequiredSkillByName("昏厥耐性", 1) // <= 3
    // cc.addRequiredSkillByName("弱點特效", 3) // <= 3

    // cc.addRequiredSkillByName("看破", 2)   // <= 7

    // cc.addRequiredSkillByName("超會心", 1) // <= 3
    // cc.addRequiredSkillByName("達人藝", 1) // <= 3

    // cc.addRequiredSkillByName("迴避距離UP", 1)   // <= 3
    // cc.addRequiredSkillByName("攻擊", 1)    // <= 7
    // cc.addRequiredSkillByName("挑戰者", 1)     // <= 5
    // cc.addRequiredSkillByName("廣域化", 2) // <= 5
    // cc.addRequiredSkillByName("達人藝", 1) // <= 3
    proc := prepareProcedure(*cc)
    proc.execute()
}


// 铳枪
// 铠罗水 耐瘴
func testExecuteProcedureGunlance3(t *testing.T) {
    cc := newConstraintCollection(loadDataManager())
    addCommonDecorationLimitations(cc)
        
    cc.addExtraSlot(3, 1)
    cc.addRequiredSkillByName("體力增強", 3)  // <= 3
    cc.addRequiredSkillByName("砲術", 3)    // <= 5
    cc.addRequiredSkillByName("砲彈裝填數UP", 1)    // <= 1
    cc.addRequiredSkillByName("防禦性能", 3)    // <= 5
    cc.addRequiredSkillByName("防禦強化", 1)    // <= 1

    cc.addRequiredSkillByName("無屬性強化", 1) // <= 1
    cc.addRequiredSkillByName("匠", 4) // <= 5
    cc.addRequiredSkillByName("剛刃研磨", 1) // <= 1
    // cc.addRequiredSkillByName("利刃／彈藥節約", 1) // <= 1
    // cc.addRequiredSkillByName("砥石使用高速化", 3) // <= 3

    cc.addRequiredSkillByName("瘴氣耐性", 3)  // <= 3
    cc.addRequiredSkillByName("弱點特效", 2) // <= 3

    // cc.addRequiredSkillByName("看破", 2)   // <= 7

    // cc.addRequiredSkillByName("超會心", 1) // <= 3
    // cc.addRequiredSkillByName("達人藝", 1) // <= 3

    // cc.addRequiredSkillByName("迴避距離UP", 1)   // <= 3
    // cc.addRequiredSkillByName("攻擊", 1)    // <= 7
    // cc.addRequiredSkillByName("挑戰者", 1)     // <= 5
    // cc.addRequiredSkillByName("廣域化", 2) // <= 5
    // cc.addRequiredSkillByName("達人藝", 1) // <= 3
    // cc.addRequiredSkillByName("昏厥耐性", 3) // <= 3
    proc := prepareProcedure(*cc)
    proc.execute()
}


// 铠罗水 王咩
func testExecuteProcedureGunlance4(t *testing.T) {
    cc := newConstraintCollection(loadDataManager())
    addCommonDecorationLimitations(cc)
        
    cc.addExtraSlot(3, 1)
    cc.addRequiredSkillByName("體力增強", 3)  // <= 3
    cc.addRequiredSkillByName("砲術", 3)    // <= 3
    cc.addRequiredSkillByName("砲彈裝填數UP", 1)    // <= 1
    cc.addRequiredSkillByName("防禦性能", 5)    // <= 5
    cc.addRequiredSkillByName("防禦強化", 1)    // <= 1

    cc.addRequiredSkillByName("耳塞", 5) // <= 5
    // cc.addRequiredSkillByName("無屬性強化", 1) // <= 1
    cc.addRequiredSkillByName("匠", 3) // <= 5
    // cc.addRequiredSkillByName("剛刃研磨", 1) // <= 1
    // cc.addRequiredSkillByName("利刃／彈藥節約", 1) // <= 1
    // cc.addRequiredSkillByName("砥石使用高速化", 3) // <= 3

    cc.addRequiredSkillByName("昏厥耐性", 3) // <= 3
    // cc.addRequiredSkillByName("弱點特效", 3) // <= 3

    // cc.addRequiredSkillByName("精靈加護", 1) // <= 3
    
    // cc.addRequiredSkillByName("看破", 2)   // <= 7

    // cc.addRequiredSkillByName("超會心", 1) // <= 3
    // cc.addRequiredSkillByName("達人藝", 1) // <= 3

    // cc.addRequiredSkillByName("迴避距離UP", 1)   // <= 3
    // cc.addRequiredSkillByName("攻擊", 1)    // <= 7
    // cc.addRequiredSkillByName("挑戰者", 1)     // <= 5
    // cc.addRequiredSkillByName("廣域化", 2) // <= 5
    // cc.addRequiredSkillByName("達人藝", 1) // <= 3
    proc := prepareProcedure(*cc)
    proc.execute()
}

