package mhw

import (
    "testing"
)

// 帝王麻，输出
func testExecuteProcedureInsectGlaive1(t *testing.T) {
    cc := newConstraintCollection(loadDataManager())
    addCommonDecorationLimitations(cc)

    cc.addRequiredSkillByName("體力增強", 3)  // <= 3
    cc.addRequiredSkillByName("麻痺屬性強化", 3) // <= 3

    cc.addRequiredSkillByName("攻擊", 4)    // <= 7
    cc.addRequiredSkillByName("看破", 4) // <= 7
    cc.addRequiredSkillByName("超會心", 3) // <= 3

    cc.addRequiredSkillByName("弱點特效", 3) // <= 3
    cc.addRequiredSkillByName("飛燕", 1)  // <= 1

    cc.addRequiredSkillByName("匠", 2)    // <= 5
    cc.addRequiredSkillByName("達人藝", 1) // <= 1
    cc.addRequiredSkillByName("飛燕【屬性】", 1) // <= 1

    cc.addRequiredSkillByName("強化持續", 1) // <= 3

    // cc.addRequiredSkillByName("無傷", 2) // <= 3
    // cc.addRequiredSkillByName("突破耐力上限", 1) // <= 1
    // cc.addRequiredSkillByName("廣域化", 2) // <= 5

    // cc.addRequiredSkillByName("精神抖擻", 2)  // <= 3
    // cc.addRequiredSkillByName("匠", 1) // <= 1
    // cc.addRequiredSkillByName("達人藝", 1) // <= 1
    // cc.addRequiredSkillByName("會心攻擊【屬性】", 1) // <= 1
    // cc.addRequiredSkillByName("跳躍鐵人", 1) // <= 1
    // cc.addRequiredSkillByName("騎乘名人", 1) // <= 1
    // cc.addRequiredSkillByName("挑戰者", 1) // <= 5
    // cc.addRequiredSkillByName("耐力急速回復", 1) // <= 3

    proc := prepareProcedure(*cc)
    proc.execute()
}

// 铠罗麻
func testExecuteProcedureInsectGlaive2(t *testing.T) {
    cc := newConstraintCollection(loadDataManager())
    addCommonDecorationLimitations(cc)

    cc.addRequiredSkillByName("體力增強", 3)  // <= 3
    cc.addRequiredSkillByName("龍屬性攻擊強化", 3) // <= 3

    cc.addRequiredSkillByName("攻擊", 4)    // <= 7
    cc.addRequiredSkillByName("看破", 4) // <= 7
    cc.addRequiredSkillByName("超會心", 3) // <= 3

    cc.addRequiredSkillByName("弱點特效", 3) // <= 3
    cc.addRequiredSkillByName("飛燕", 1)  // <= 1

    cc.addRequiredSkillByName("達人藝", 1) // <= 1
    cc.addRequiredSkillByName("飛燕【屬性】", 1) // <= 1
    cc.addRequiredSkillByName("體術", 3) // <= 5

    // cc.addRequiredSkillByName("強化持續", 1) // <= 3
    cc.addRequiredSkillByName("回復速度", 2) // <= 3
    

    proc := prepareProcedure(*cc)
    proc.execute()
}