package mhw

import (
    "testing"
)

// 太刀 铠罗火
func testExecuteProcedureLongSword1(t *testing.T) {
    cc := newConstraintCollection(loadDataManager())
    cc.addExtraSlot(3, 1)

    addCommonDecorationLimitations(cc)

    cc.addRequiredSkillByName("體力增強", 3)  // <= 3
    cc.addRequiredSkillByName("無屬性強化", 1) // <= 1
    cc.addRequiredSkillByName("攻擊", 4)    // <= 7

    cc.addRequiredSkillByName("弱點特效", 3) // <= 3

    cc.addRequiredSkillByName("超會心", 3) // <= 3
    cc.addRequiredSkillByName("看破", 7)   // <= 7

    cc.addRequiredSkillByName("精神抖擻", 3)   // <= 3
    cc.addRequiredSkillByName("無傷", 2)   // <= 3
    // cc.addRequiredSkillByName("挑戰者", 2)   // <= 3

    // cc.addRequiredSkillByName("回復速度", 2)   // <= 3
    // cc.addRequiredSkillByName("廣域化", 2)   // <= 5

    // cc.addRequiredSkillByName("突破耐力上限", 1)   // <= 1

    proc := prepareProcedure(*cc)
    proc.execute()
}

// 太刀
// 耐火+热伤害无效
func testExecuteProcedureLongSword2(t *testing.T) {
    cc := newConstraintCollection(loadDataManager())
    cc.addExtraSlot(3, 1)

    addCommonDecorationLimitations(cc)

    cc.addRequiredSkillByName("體力增強", 3)  // <= 3
    cc.addRequiredSkillByName("無屬性強化", 1) // <= 1
    cc.addRequiredSkillByName("攻擊", 4)    // <= 7

    cc.addRequiredSkillByName("火耐性", 3)  // <= 3
    cc.addRequiredSkillByName("熱傷害無效", 1)  // <= 1

    cc.addRequiredSkillByName("弱點特效", 3) // <= 3

    cc.addRequiredSkillByName("超會心", 2) // <= 3
    cc.addRequiredSkillByName("看破", 5)   // <= 7

    cc.addRequiredSkillByName("無傷", 1)   // <= 3
    cc.addRequiredSkillByName("挑戰者", 2)   // <= 3
    // cc.addRequiredSkillByName("精神抖擻", 1)   // <= 3

    cc.addRequiredSkillByName("回復速度", 2)   // <= 3
    cc.addRequiredSkillByName("廣域化", 2)   // <= 5

    proc := prepareProcedure(*cc)
    proc.execute()
}


// 太刀
// 耐瘴
func testExecuteProcedureLongSword3(t *testing.T) {
    cc := newConstraintCollection(loadDataManager())
    cc.addExtraSlot(3, 1)

    addCommonDecorationLimitations(cc)

    cc.addRequiredSkillByName("體力增強", 3)  // <= 3
    cc.addRequiredSkillByName("無屬性強化", 1) // <= 1
    cc.addRequiredSkillByName("攻擊", 4)    // <= 7

    cc.addRequiredSkillByName("瘴氣耐性", 3)  // <= 3

    cc.addRequiredSkillByName("弱點特效", 3) // <= 3

    cc.addRequiredSkillByName("超會心", 3) // <= 3
    cc.addRequiredSkillByName("看破", 6)   // <= 7

    // cc.addRequiredSkillByName("無傷", 1)   // <= 3
    // cc.addRequiredSkillByName("挑戰者", 1)   // <= 3
    cc.addRequiredSkillByName("精神抖擻", 1)   // <= 3

    cc.addRequiredSkillByName("回復速度", 2)   // <= 3
    cc.addRequiredSkillByName("廣域化", 2)   // <= 5

    proc := prepareProcedure(*cc)
    proc.execute()
}


// 太刀 天上天下
func testExecuteProcedureLongSword4(t *testing.T) {
    cc := newConstraintCollection(loadDataManager())
    cc.addExtraSlot(3, 1)

    addCommonDecorationLimitations(cc)

    cc.addRequiredSkillByName("體力增強", 3)  // <= 3
    cc.addRequiredSkillByName("無屬性強化", 1) // <= 1
    cc.addRequiredSkillByName("攻擊", 4)    // <= 7

    cc.addRequiredSkillByName("弱點特效", 3) // <= 3

    cc.addRequiredSkillByName("匠", 5)    // <= 5
    cc.addRequiredSkillByName("達人藝", 1)    // <= 1

    cc.addRequiredSkillByName("超會心", 3) // <= 3
    cc.addRequiredSkillByName("看破", 5)   // <= 7

    proc := prepareProcedure(*cc)
    proc.execute()
}

// 太刀炎妃冥灯
func testExecuteProcedureLongSword5(t *testing.T) {
    cc := newConstraintCollection(loadDataManager())
    cc.addExtraSlot(3, 2)

    addCommonDecorationLimitations(cc)

    cc.addRequiredSkillByName("體力增強", 3)  // <= 3
    cc.addRequiredSkillByName("攻擊", 6)    // <= 7

    cc.addRequiredSkillByName("匠", 2)    // <= 5
    cc.addRequiredSkillByName("弱點特效", 3) // <= 3
    cc.addRequiredSkillByName("昏厥耐性", 3) // <= 3
    cc.addRequiredSkillByName("超會心", 3) // <= 3

    cc.addRequiredSkillByName("看破", 7)   // <= 7

    cc.addRequiredSkillByName("廣域化", 2) // <= 5

    proc := prepareProcedure(*cc)
    proc.execute()
}


// 太刀
// 火属性
func testExecuteProcedureLongSword6(t *testing.T) {
    cc := newConstraintCollection(loadDataManager())
    cc.addExtraSlot(3, 1)

    addCommonDecorationLimitations(cc)

    cc.addRequiredSkillByName("體力增強", 3)  // <= 3
    // cc.addRequiredSkillByName("攻擊", 4)    // <= 7
    cc.addRequiredSkillByName("屬性解放／裝填擴充", 3) // <= 3
    cc.addRequiredSkillByName("火屬性攻擊強化", 4) // <= 5

    cc.addRequiredSkillByName("弱點特效", 3) // <= 3

    cc.addRequiredSkillByName("精靈加護", 3)    // <= 3
    cc.addRequiredSkillByName("收刀術", 3) // <= 3

    cc.addRequiredSkillByName("超會心", 2) // <= 3
    cc.addRequiredSkillByName("看破", 3)   // <= 7

    cc.addRequiredSkillByName("廣域化", 2)    // <= 5
    cc.addRequiredSkillByName("匠", 1)    // <= 5
    // cc.addRequiredSkillByName("飛燕【屬性】", 1)    // <= 1    
    cc.addRequiredSkillByName("攻擊", 1)    // <= 7

    // cc.addRequiredSkillByName("無傷", 1)   // <= 3
    // cc.addRequiredSkillByName("挑戰者", 2)   // <= 3
    // cc.addRequiredSkillByName("精神抖擻", 1)   // <= 3

    // cc.addRequiredSkillByName("回復速度", 2)   // <= 3

    proc := prepareProcedure(*cc)
    proc.execute()
}


// 太刀 无体耐雷
func testExecuteProcedureLongSword7(t *testing.T) {
    cc := newConstraintCollection(loadDataManager())
    cc.addExtraSlot(3, 1)

    addCommonDecorationLimitations(cc)

    // cc.addRequiredSkillByName("體力增強", 3)  // <= 3
    cc.addRequiredSkillByName("無屬性強化", 1) // <= 1
    cc.addRequiredSkillByName("攻擊", 4)    // <= 7
    cc.addRequiredSkillByName("雷耐性", 3)    // <= 3

    cc.addRequiredSkillByName("匠", 1)    // <= 5
    cc.addRequiredSkillByName("弱點特效", 3) // <= 3

    cc.addRequiredSkillByName("超會心", 3) // <= 3
    cc.addRequiredSkillByName("看破", 7)   // <= 7

    cc.addRequiredSkillByName("精神抖擻", 2)   // <= 3
    cc.addRequiredSkillByName("無傷", 2)   // <= 3
    // cc.addRequiredSkillByName("挑戰者", 2)   // <= 3

    // cc.addRequiredSkillByName("回復速度", 2)   // <= 3
    // cc.addRequiredSkillByName("廣域化", 2)   // <= 5

    // cc.addRequiredSkillByName("突破耐力上限", 1)   // <= 1

    proc := prepareProcedure(*cc)
    proc.execute()
}