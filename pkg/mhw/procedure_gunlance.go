package mhw

import (
    "testing"
)

// 冥赤龙铳枪
func testExecuteProcedureGunlance_4_4_AnyCharm(t *testing.T) {
    cc := newConstraintCollection(loadDataManager())
    addCommonDecorationLimitations(cc)

    cc.addExtraSlot(4, 2)

    cc.addRequiredSkillByName("任意護石", 1)

    cc.addExtraSlot(4, -1)
    cc.addRequiredSkillByName("減輕膽怯", -1)
    cc.addRequiredSkillByName("體力增強", -1)
    cc.addExtraSlot(4, -1)
    cc.addRequiredSkillByName("集中", -1)
    cc.addRequiredSkillByName("體力增強", -1)
    cc.addExtraSlot(4, -1)
    cc.addRequiredSkillByName("砲術", -2)
    cc.addExtraSlot(4, -1)
    cc.addRequiredSkillByName("防禦性能", -2)
    cc.addExtraSlot(4, -1)
    cc.addRequiredSkillByName("砲術", -2)
    // cc.addExtraSlot(4, -1)
    // cc.addRequiredSkillByName("防禦性能", -2)

    cc.addRequiredSkillByName("體力增強", 3)  // <= 3
    cc.addRequiredSkillByName("砲術", 3)    // <= 3
    cc.addRequiredSkillByName("投射器裝填數UP", 3)    // <= 3
    cc.addRequiredSkillByName("砲彈裝填數UP", 1)    // <= 1
    cc.addRequiredSkillByName("防禦性能", 5)    // <= 5
    cc.addRequiredSkillByName("防禦強化", 1)    // <= 1

    cc.addRequiredSkillByName("砲術・極意", 1)    // <= 1
    cc.addRequiredSkillByName("砲術", 2)    // <= 3>5
    cc.addRequiredSkillByName("投射器裝填・極意", 1)    // <= 1
    cc.addRequiredSkillByName("投射器裝填數UP", 2)    // <= 3>5

    cc.addRequiredSkillByName("減輕膽怯", 1)    // <= 3
    cc.addRequiredSkillByName("地質學", 1)   // <= 5

    cc.addRequiredSkillByName("集中", 3)   // <= 3
    cc.addRequiredSkillByName("耐震", 3)    // <= 3

    cc.addRequiredSkillByName("耳塞", 1)   // <= 5
    // cc.addRequiredSkillByName("迴避距離UP", 1)   // <= 3

    cc.addRequiredSkillByName("風壓耐性", 1)   // <= 3

    // cc.addRequiredSkillByName("利刃／彈藥節約", 1) // <= 1
    // cc.addRequiredSkillByName("砥石使用高速化", 3) // <= 3
    // cc.addRequiredSkillByName("弱點特效", 3) // <= 3
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

