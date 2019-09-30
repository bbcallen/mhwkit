package mhw

import (
    "testing"
)

// 轻弩
// 皇后甲壳・炎妃 毅力
// 冰结 风压无效
func testExecuteProcedureLightBowgun1(t *testing.T) {
    cc := newConstraintCollection(loadDataManager())
    addCommonDecorationLimitations(cc)
        
    cc.addExtraSlot(1, 1)
    cc.addRequiredSkillByName("屬性解放／裝填擴充", 3)    // <= 3
    cc.addRequiredSkillByName("體力增強", 3)  // <= 3
    cc.addRequiredSkillByName("熱傷害無效", 1)  // <= 1
    cc.addRequiredSkillByName("會心攻擊【屬性】", 1)  // <= 1
    cc.addRequiredSkillByName("冰屬性攻擊強化", 3)  // <= 5

    cc.addRequiredSkillByName("心眼／彈道強化", 1) // <= 1
    cc.addRequiredSkillByName("火耐性", 3)  // <= 3

    // cc.addRequiredSkillByName("精神抖擻", 1)  // <= 3

    cc.addRequiredSkillByName("昏厥耐性", 3) // <= 3
    cc.addRequiredSkillByName("風壓耐性", 5) // <= 5

    // cc.addRequiredSkillByName("廣域化", 2) // <= 5

    proc := prepareProcedure(*cc)
    proc.execute()
}

// 轻弩
// 皇后甲壳・炎妃 毅力
// 冰结 风压无效
func testExecuteProcedureLightBowgun2(t *testing.T) {
    cc := newConstraintCollection(loadDataManager())
    addCommonDecorationLimitations(cc)
        
    cc.addExtraSlot(1, 1)
    cc.addRequiredSkillByName("屬性解放／裝填擴充", 3)    // <= 3
    cc.addRequiredSkillByName("體力增強", 3)  // <= 3
    cc.addRequiredSkillByName("熱傷害無效", 1)  // <= 1
    cc.addRequiredSkillByName("會心攻擊【屬性】", 1)  // <= 1
    cc.addRequiredSkillByName("冰屬性攻擊強化", 3)  // <= 5

    cc.addRequiredSkillByName("心眼／彈道強化", 1) // <= 1
    cc.addRequiredSkillByName("火耐性", 3)  // <= 3
    cc.addRequiredSkillByName("昏厥耐性", 3) // <= 3

    cc.addRequiredSkillByName("精神抖擻", 2)  // <= 3

    cc.addRequiredSkillByName("挑戰者", 2)  // <= 5
    cc.addRequiredSkillByName("回復速度", 2)  // <= 3
    // cc.addRequiredSkillByName("迴避性能", 2)  // <= 5

    // cc.addRequiredSkillByName("看破", 1)  // <= 7
    // cc.addRequiredSkillByName("超會心", 1)  // <= 3

    // cc.addRequiredSkillByName("廣域化", 2) // <= 5

    proc := prepareProcedure(*cc)
    proc.execute()
}

// 轻弩
// 皇后甲壳・冥灯 利刃／彈藥節約
// 斩裂
func testExecuteProcedureLightBowgun3(t *testing.T) {
    cc := newConstraintCollection(loadDataManager())
    addCommonDecorationLimitations(cc)
        
    cc.addExtraSlot(3, 1)
    cc.addExtraSlot(1, 1)
    cc.addRequiredSkillByName("屬性解放／裝填擴充", 3)    // <= 3
    cc.addRequiredSkillByName("體力增強", 3)  // <= 3
    cc.addRequiredSkillByName("熱傷害無效", 1)  // <= 1
    cc.addRequiredSkillByName("火耐性", 3)  // <= 3

    cc.addRequiredSkillByName("昏厥耐性", 3) // <= 3
    cc.addRequiredSkillByName("風壓耐性", 5) // <= 5
    cc.addRequiredSkillByName("睡眠屬性強化", 3) // <= 3
    cc.addRequiredSkillByName("廣域化", 2) // <= 5

    // cc.addRequiredSkillByName("回復速度", 2) // <= 3

    // cc.addRequiredSkillByName("快吃", 1) // <= 3    
    // cc.addRequiredSkillByName("弱點特效", 3) // <= 3
    // cc.addRequiredSkillByName("攻擊", 4)    // <= 7
    // cc.addRequiredSkillByName("利刃／彈藥節約", 1)  // <= 1
    // cc.addRequiredSkillByName("精靈加護", 3) // <= 3
    // cc.addRequiredSkillByName("無傷", 1) // <= 3
    // cc.addRequiredSkillByName("減輕膽怯", 2) // <= 3
    proc := prepareProcedure(*cc)
    proc.execute()
}


// 轻弩
// 皇后甲壳・冥灯 利刃／彈藥節約
// 雷速射 斩裂 渣渣辉
func testExecuteProcedureLightBowgun4(t *testing.T) {
    cc := newConstraintCollection(loadDataManager())
    addCommonDecorationLimitations(cc)
        
    cc.addExtraSlot(3, 1)
    cc.addExtraSlot(1, 1)
    cc.addRequiredSkillByName("屬性解放／裝填擴充", 3)    // <= 3
    cc.addRequiredSkillByName("體力增強", 3)  // <= 3
    cc.addRequiredSkillByName("會心攻擊【屬性】", 1)  // <= 1
    cc.addRequiredSkillByName("雷屬性攻擊強化", 3)  // <= 3

    // cc.addRequiredSkillByName("火耐性", 3)  // <= 3
    cc.addRequiredSkillByName("破壞王", 3)  // <= 3
    cc.addRequiredSkillByName("精神抖擻", 3)  // <= 3

    cc.addRequiredSkillByName("看破", 5)  // <= 7
    cc.addRequiredSkillByName("超會心", 1)  // <= 3

    cc.addRequiredSkillByName("炸彈客", 2)  // <= 3
    // cc.addRequiredSkillByName("攻擊", 1)  // <= 7

    // cc.addRequiredSkillByName("昏厥耐性", 3) // <= 3
    // cc.addRequiredSkillByName("風壓耐性", 5) // <= 5
    // cc.addRequiredSkillByName("睡眠屬性強化", 3) // <= 3
    // cc.addRequiredSkillByName("廣域化", 2) // <= 5

    // cc.addRequiredSkillByName("回復速度", 2) // <= 3

    // cc.addRequiredSkillByName("快吃", 1) // <= 3    
    // cc.addRequiredSkillByName("弱點特效", 3) // <= 3
    // cc.addRequiredSkillByName("攻擊", 4)    // <= 7
    // cc.addRequiredSkillByName("利刃／彈藥節約", 1)  // <= 1
    // cc.addRequiredSkillByName("精靈加護", 3) // <= 3
    // cc.addRequiredSkillByName("無傷", 1) // <= 3
    // cc.addRequiredSkillByName("減輕膽怯", 2) // <= 3
    proc := prepareProcedure(*cc)
    proc.execute()
}
