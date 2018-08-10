package mhw

import (
    "testing"
)

// 大剑喷气日常
func testExecuteProcedureGreatSword1(t *testing.T) {
    // 龍熱機關式【鋼翼】 改
    cc := newConstraintCollection(loadDataManager())
    cc.addExtraSlot(1, 2)

    addCommonDecorationLimitations(cc)

    cc.addRequiredSkillByName("無屬性強化", 1) // <= 1
    cc.addRequiredSkillByName("攻擊", 5)    // <= 7
    cc.addRequiredSkillByName("體力增強", 3)  // <= 3

    cc.addRequiredSkillByName("集中", 3)   // <= 3
    cc.addRequiredSkillByName("匠", 3)    // <= 5
    cc.addRequiredSkillByName("弱點特效", 3) // <= 3
    cc.addRequiredSkillByName("看破", 3)   // <= 7

    cc.addRequiredSkillByName("超會心", 1) // <= 3
    cc.addRequiredSkillByName("收刀術", 1) // <= 3
    cc.addRequiredSkillByName("KO術", 1) // <= 3
    cc.addRequiredSkillByName("破壞王", 1) // <= 3

    proc := prepareProcedure(*cc)
    proc.execute()
}

// 大剑喷气耐火
func testExecuteProcedureGreatSword2(t *testing.T) {
    // 龍熱機關式【鋼翼】 改
    cc := newConstraintCollection(loadDataManager())
    cc.addExtraSlot(1, 2)

    addCommonDecorationLimitations(cc)

    cc.addRequiredSkillByName("無屬性強化", 1) // <= 1
    cc.addRequiredSkillByName("攻擊", 4)    // <= 7
    cc.addRequiredSkillByName("體力增強", 3)  // <= 3
    cc.addRequiredSkillByName("火耐性", 3)  // <= 3

    cc.addRequiredSkillByName("集中", 3)   // <= 3
    cc.addRequiredSkillByName("匠", 2)    // <= 5
    cc.addRequiredSkillByName("弱點特效", 3) // <= 3
    cc.addRequiredSkillByName("看破", 2)   // <= 7

    cc.addRequiredSkillByName("超會心", 1) // <= 3
    cc.addRequiredSkillByName("收刀術", 1) // <= 3
    cc.addRequiredSkillByName("KO術", 1) // <= 3
    cc.addRequiredSkillByName("破壞王", 1) // <= 3

    proc := prepareProcedure(*cc)
    proc.execute()
}

// 大剑灭鬼
func testExecuteProcedureGreatSword3(t *testing.T) {
    // 灭鬼凶器
    cc := newConstraintCollection(loadDataManager())
    cc.addExtraSlot(1, 1)

    addCommonDecorationLimitations(cc)

    cc.addRequiredSkillByName("攻擊", 4)    // <= 7
    cc.addRequiredSkillByName("體力增強", 3)  // <= 3

    cc.addRequiredSkillByName("集中", 3)   // <= 3
    cc.addRequiredSkillByName("弱點特效", 3) // <= 3
    cc.addRequiredSkillByName("剛刃研磨", 1)   // <= 1
    cc.addRequiredSkillByName("看破", 1)   // <= 7

    cc.addRequiredSkillByName("破壞王", 3) // <= 3
    cc.addRequiredSkillByName("KO術", 3) // <= 3

    cc.addRequiredSkillByName("收刀術", 1) // <= 3

    proc := prepareProcedure(*cc)
    proc.execute()
}

// 大剑灭鬼耐瘴
func testExecuteProcedureGreatSword4(t *testing.T) {
    // 灭鬼凶器
    cc := newConstraintCollection(loadDataManager())
    cc.addExtraSlot(1, 1)

    addCommonDecorationLimitations(cc)

    cc.addRequiredSkillByName("攻擊", 4)    // <= 7
    cc.addRequiredSkillByName("體力增強", 3)  // <= 3
    cc.addRequiredSkillByName("瘴氣耐性", 3)  // <= 3

    cc.addRequiredSkillByName("集中", 3)   // <= 3
    cc.addRequiredSkillByName("弱點特效", 3) // <= 3
    cc.addRequiredSkillByName("剛刃研磨", 1)   // <= 1
    cc.addRequiredSkillByName("看破", 4)   // <= 7

    cc.addRequiredSkillByName("精靈加護", 3) // <= 3
    cc.addRequiredSkillByName("超會心", 1) // <= 3
    // cc.addRequiredSkillByName("收刀術", 1) // <= 3
    // cc.addRequiredSkillByName("破壞王", 3) // <= 3
    // cc.addRequiredSkillByName("KO術", 2) // <= 3

    // cc.addRequiredSkillByName("收刀術", 1) // <= 3
    cc.addRequiredSkillByName("飛燕【屬性】", 1) // <= 1
    

    proc := prepareProcedure(*cc)
    proc.execute()
}

