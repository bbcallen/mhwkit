package mhw

import (
    "testing"
)


func testExecuteProcedureGreatSword_4(t *testing.T) {
    cc := newConstraintCollection(loadDataManager())

    // 巨牛大斧II
    cc.addExtraSlot(4, 1)
    // 衣装
    // cc.addExtraSlot(4, 1)
    // cc.addExtraSlot(2, 1)
    // 衣装
    // cc.addExtraSlot(3, 1)
    // cc.addExtraSlot(3, 1)
    cc.addExtraSlot(4, -1) // 达人II
    cc.addRequiredSkillByName("看破", -2)

    addCommonDecorationLimitations(cc)

    cc.addRequiredSkillByName("無屬性強化", 1) // <= 1
    cc.addRequiredSkillByName("攻擊", 4)    // <= 7
    cc.addRequiredSkillByName("體力增強", 3)  // <= 3

    cc.addRequiredSkillByName("集中", 3)   // <= 3
    // cc.addRequiredSkillByName("匠", 3)    // <= 5
    cc.addRequiredSkillByName("弱點特效", 3) // <= 3
    cc.addRequiredSkillByName("看破", 6)   // <= 7
    cc.addRequiredSkillByName("超會心", 3) // <= 3

    // extra
    // cc.addRequiredSkillByName("心眼／彈道強化", 1)
    cc.addRequiredSkillByName("昏厥耐性", 2)
    cc.addRequiredSkillByName("跑者", 1)
    // cc.addRequiredSkillByName("跳躍鐵人", 1)
    // cc.addRequiredSkillByName("火屬性攻擊強化", 1)
    // cc.addRequiredSkillByName("道具使用強化", 1)
    // cc.addRequiredSkillByName("廣域化", 1)
    // cc.addRequiredSkillByName("耳塞", 1)
    cc.addRequiredSkillByName("力量解放", 1)
    cc.addRequiredSkillByName("滑走強化", 1)

    proc := prepareProcedure(*cc)
    proc.execute()
}

func testExecuteProcedureGreatSword_4_ResistIce(t *testing.T) {
    cc := newConstraintCollection(loadDataManager())

    // 巨牛大斧II
    cc.addExtraSlot(4, 1)
    // 衣装
    // cc.addExtraSlot(4, 1)
    // cc.addExtraSlot(2, 1)
    // 衣装
    // cc.addExtraSlot(3, 1)
    // cc.addExtraSlot(3, 1)
    cc.addExtraSlot(4, -1) // 达人II
    cc.addRequiredSkillByName("看破", -2)
    cc.addExtraSlot(4, -1) // 冰耐性II
    cc.addRequiredSkillByName("冰耐性", -2)

    addCommonDecorationLimitations(cc)

    cc.addRequiredSkillByName("無屬性強化", 1) // <= 1
    cc.addRequiredSkillByName("攻擊", 4)    // <= 7
    cc.addRequiredSkillByName("體力增強", 3)  // <= 3

    cc.addRequiredSkillByName("冰耐性", 3)  // <= 3
    
    cc.addRequiredSkillByName("集中", 3)   // <= 3
    // cc.addRequiredSkillByName("匠", 3)    // <= 5
    cc.addRequiredSkillByName("弱點特效", 3) // <= 3
    cc.addRequiredSkillByName("看破", 6)   // <= 7
    cc.addRequiredSkillByName("超會心", 2) // <= 3

    // extra
    // cc.addRequiredSkillByName("心眼／彈道強化", 1)
    // cc.addRequiredSkillByName("昏厥耐性", 2)
    // cc.addRequiredSkillByName("跑者", 1)
    // cc.addRequiredSkillByName("跳躍鐵人", 1)
    // cc.addRequiredSkillByName("火屬性攻擊強化", 1)
    // cc.addRequiredSkillByName("道具使用強化", 1)
    // cc.addRequiredSkillByName("廣域化", 1)
    // cc.addRequiredSkillByName("耳塞", 1)
    cc.addRequiredSkillByName("力量解放", 1)
    cc.addRequiredSkillByName("滑走強化", 1)

    proc := prepareProcedure(*cc)
    proc.execute()
}

func testExecuteProcedureGreatSword_1_1_ResistIce(t *testing.T) {
    cc := newConstraintCollection(loadDataManager())

    // 死神镰刀II
    cc.addExtraSlot(1, 2)
    // 衣装
    // cc.addExtraSlot(4, 1)
    // cc.addExtraSlot(2, 1)
    // 衣装
    // cc.addExtraSlot(3, 1)
    // cc.addExtraSlot(3, 1)
    cc.addExtraSlot(4, -1) // 达人II
    cc.addRequiredSkillByName("看破", -2)
    cc.addExtraSlot(4, -1) // 耐冰珠II
    cc.addRequiredSkillByName("冰耐性", -2)

    addCommonDecorationLimitations(cc)

    cc.addRequiredSkillByName("無屬性強化", 1) // <= 1
    cc.addRequiredSkillByName("攻擊", 4)    // <= 7
    cc.addRequiredSkillByName("體力增強", 3)  // <= 3

    cc.addRequiredSkillByName("冰耐性", 3)  // <= 3

    cc.addRequiredSkillByName("集中", 3)   // <= 3
    // cc.addRequiredSkillByName("匠", 3)    // <= 5
    cc.addRequiredSkillByName("弱點特效", 3) // <= 3
    cc.addRequiredSkillByName("看破", 7)   // <= 7
    cc.addRequiredSkillByName("超會心", 2) // <= 3

    // extra
    // cc.addRequiredSkillByName("道具使用強化", 1)
    // cc.addRequiredSkillByName("廣域化", 1)
    // cc.addRequiredSkillByName("耳塞", 1)
    // cc.addRequiredSkillByName("力量解放", 2)
    cc.addRequiredSkillByName("心眼／彈道強化", 1)
    // cc.addRequiredSkillByName("昏厥耐性", 2)
    // cc.addRequiredSkillByName("跑者", 1)
    // cc.addRequiredSkillByName("滑走強化", 1)
    cc.addRequiredSkillByName("跳躍鐵人", 1)
    cc.addRequiredSkillByName("火屬性攻擊強化", 5)

    proc := prepareProcedure(*cc)
    proc.execute()
}


func testExecuteProcedureGreatSword_1_1(t *testing.T) {
    cc := newConstraintCollection(loadDataManager())

    // 死神镰刀II
    cc.addExtraSlot(1, 2)
    // 衣装
    // cc.addExtraSlot(4, 1)
    // cc.addExtraSlot(2, 1)
    // 衣装
    // cc.addExtraSlot(3, 1)
    // cc.addExtraSlot(3, 1)
    cc.addExtraSlot(4, -1) // 达人II
    cc.addRequiredSkillByName("看破", -2)

    addCommonDecorationLimitations(cc)

    cc.addRequiredSkillByName("無屬性強化", 1) // <= 1
    cc.addRequiredSkillByName("攻擊", 4)    // <= 7
    cc.addRequiredSkillByName("體力增強", 3)  // <= 3

    cc.addRequiredSkillByName("集中", 3)   // <= 3
    // cc.addRequiredSkillByName("匠", 3)    // <= 5
    cc.addRequiredSkillByName("弱點特效", 3) // <= 3
    cc.addRequiredSkillByName("看破", 7)   // <= 7
    cc.addRequiredSkillByName("超會心", 3) // <= 3

    // extra
    // cc.addRequiredSkillByName("心眼／彈道強化", 1)
    cc.addRequiredSkillByName("昏厥耐性", 2)
    cc.addRequiredSkillByName("跑者", 1)
    // cc.addRequiredSkillByName("跳躍鐵人", 1)
    cc.addRequiredSkillByName("火屬性攻擊強化", 1)
    // cc.addRequiredSkillByName("道具使用強化", 1)
    // cc.addRequiredSkillByName("廣域化", 1)
    // cc.addRequiredSkillByName("耳塞", 1)
    // cc.addRequiredSkillByName("力量解放", 2)
    cc.addRequiredSkillByName("滑走強化", 1)

    proc := prepareProcedure(*cc)
    proc.execute()
}


func testExecuteProcedureGreatSword_4_2_ResistIce(t *testing.T) {
    cc := newConstraintCollection(loadDataManager())

    // 火碎剑I
    cc.addExtraSlot(4, 1)
    cc.addExtraSlot(2, 1)
    // 衣装
    // cc.addExtraSlot(4, 1)
    // cc.addExtraSlot(2, 1)
    // 衣装
    // cc.addExtraSlot(3, 1)
    // cc.addExtraSlot(3, 1)
    // cc.addExtraSlot(4, -1) // 达人II
    // cc.addRequiredSkillByName("看破", -2)
    cc.addExtraSlot(4, -1) // 耐冰珠II
    cc.addRequiredSkillByName("冰耐性", -2)

    addCommonDecorationLimitations(cc)

    cc.addRequiredSkillByName("攻擊", 4)    // <= 7
    cc.addRequiredSkillByName("體力增強", 3)  // <= 3
    cc.addRequiredSkillByName("拔刀術【技】", 3)  // <= 3
    // cc.addRequiredSkillByName("收刀術", 1)  // <= 3
    // cc.addRequiredSkillByName("拔刀術【力】", 1)  // <= 1

    cc.addRequiredSkillByName("冰耐性", 3)  // <= 3
    cc.addRequiredSkillByName("火屬性攻擊強化", 5) // <= 6
    // cc.addRequiredSkillByName("屬性異常狀態耐性", 3)  // <= 3
    

    cc.addRequiredSkillByName("集中", 3)   // <= 3
    // cc.addRequiredSkillByName("匠", 3)    // <= 5
    cc.addRequiredSkillByName("弱點特效", 3) // <= 3
    // cc.addRequiredSkillByName("看破", 2)   // <= 7
    cc.addRequiredSkillByName("超會心", 3) // <= 3

    // extra
    // cc.addRequiredSkillByName("道具使用強化", 1)
    // cc.addRequiredSkillByName("廣域化", 1)
    // cc.addRequiredSkillByName("耳塞", 1)
    // cc.addRequiredSkillByName("力量解放", 2)
    cc.addRequiredSkillByName("心眼／彈道強化", 1)
    cc.addRequiredSkillByName("攻擊守勢", 1)
    // cc.addRequiredSkillByName("昏厥耐性", 2)
    // cc.addRequiredSkillByName("跑者", 1)
    // cc.addRequiredSkillByName("滑走強化", 1)
    // cc.addRequiredSkillByName("跳躍鐵人", 1)

    proc := prepareProcedure(*cc)
    proc.execute()
}


func testExecuteProcedureGreatSword_4_1_ResistSleep(t *testing.T) {
    cc := newConstraintCollection(loadDataManager())

    // 巨牛大斧I
    cc.addExtraSlot(4, 1)
    // 衣装
    // cc.addExtraSlot(4, 1)
    // cc.addExtraSlot(2, 1)
    // 衣装
    // cc.addExtraSlot(3, 1)
    // cc.addExtraSlot(3, 1)
    cc.addExtraSlot(4, -1) // 达人II
    cc.addRequiredSkillByName("看破", -2)
    cc.addExtraSlot(4, -1) // 耐眠珠II
    cc.addRequiredSkillByName("睡眠耐性", -2)

    addCommonDecorationLimitations(cc)

    cc.addRequiredSkillByName("無屬性強化", 1) // <= 1
    cc.addRequiredSkillByName("攻擊", 4)    // <= 7
    cc.addRequiredSkillByName("體力增強", 3)  // <= 3

    cc.addRequiredSkillByName("睡眠耐性", 3)   // <= 3

    cc.addRequiredSkillByName("集中", 2)   // <= 3
    cc.addRequiredSkillByName("匠", 3)    // <= 5
    cc.addRequiredSkillByName("弱點特效", 3) // <= 3
    cc.addRequiredSkillByName("看破", 5)   // <= 7
    // cc.addRequiredSkillByName("超會心", 1) // <= 3

    // extra
    cc.addRequiredSkillByName("快吃", 2)
    cc.addRequiredSkillByName("昏厥耐性", 2)
    // cc.addRequiredSkillByName("道具使用強化", 1)
    // cc.addRequiredSkillByName("廣域化", 1)
    cc.addRequiredSkillByName("耳塞", 1)
    cc.addRequiredSkillByName("跑者", 1)

    // cc.addRequiredSkillByName("力量解放", 1)
    // cc.addRequiredSkillByName("昏厥耐性", 2)
    // cc.addRequiredSkillByName("跑者", 1)
    // cc.addRequiredSkillByName("滑走強化", 1)
    // cc.addRequiredSkillByName("收刀術", 1) // <= 3
    // cc.addRequiredSkillByName("KO術", 1) // <= 3
    // cc.addRequiredSkillByName("破壞王", 1) // <= 3
    // cc.addRequiredSkillByName("炸彈客", 2) // <= 3
    // cc.addRequiredSkillByName("廣域化", 2)    // <= 5

    proc := prepareProcedure(*cc)
    proc.execute()
}


func testExecuteProcedureGreatSword_4_1_ResistWater(t *testing.T) {
    cc := newConstraintCollection(loadDataManager())

    // 巨牛大斧I
    cc.addExtraSlot(4, 1)
    // 衣装
    // cc.addExtraSlot(4, 1)
    // cc.addExtraSlot(2, 1)
    // 衣装
    // cc.addExtraSlot(3, 1)
    // cc.addExtraSlot(3, 1)
    cc.addExtraSlot(4, -1) // 达人II
    cc.addRequiredSkillByName("看破", -2)

    addCommonDecorationLimitations(cc)

    cc.addRequiredSkillByName("無屬性強化", 1) // <= 1
    cc.addRequiredSkillByName("攻擊", 4)    // <= 7
    cc.addRequiredSkillByName("體力增強", 3)  // <= 3

    cc.addRequiredSkillByName("水耐性", 3)   // <= 3

    cc.addRequiredSkillByName("集中", 2)   // <= 3
    cc.addRequiredSkillByName("匠", 3)    // <= 5
    cc.addRequiredSkillByName("弱點特效", 3) // <= 3
    cc.addRequiredSkillByName("看破", 4)   // <= 7
    // cc.addRequiredSkillByName("超會心", 1) // <= 3

    // extra
    // cc.addRequiredSkillByName("快吃", 2)
    // cc.addRequiredSkillByName("昏厥耐性", 2)
    cc.addRequiredSkillByName("道具使用強化", 1)
    cc.addRequiredSkillByName("廣域化", 1)
    cc.addRequiredSkillByName("耳塞", 1)
    // cc.addRequiredSkillByName("跑者", 1)
    // cc.addRequiredSkillByName("力量解放", 1)
    // cc.addRequiredSkillByName("滑走強化", 1)
    // cc.addRequiredSkillByName("收刀術", 1) // <= 3
    // cc.addRequiredSkillByName("KO術", 1) // <= 3
    // cc.addRequiredSkillByName("破壞王", 1) // <= 3
    // cc.addRequiredSkillByName("炸彈客", 2) // <= 3
    // cc.addRequiredSkillByName("廣域化", 2)    // <= 5

    proc := prepareProcedure(*cc)
    proc.execute()
}

func testExecuteProcedureGreatSword_1_2(t *testing.T) {
    cc := newConstraintCollection(loadDataManager())

    // 龍熱機關式【鋼翼】 改
    cc.addExtraSlot(1, 2)
    // 衣装
    // cc.addExtraSlot(4, 1)
    // cc.addExtraSlot(2, 1)
    // 衣装
    // cc.addExtraSlot(3, 1)
    // cc.addExtraSlot(3, 1)
    cc.addExtraSlot(4, -1) // 达人II

    addCommonDecorationLimitations(cc)

    cc.addRequiredSkillByName("無屬性強化", 1) // <= 1
    cc.addRequiredSkillByName("攻擊", 4)    // <= 7
    cc.addRequiredSkillByName("體力增強", 3)  // <= 3

    cc.addRequiredSkillByName("集中", 2)   // <= 3
    cc.addRequiredSkillByName("匠", 3)    // <= 5
    cc.addRequiredSkillByName("弱點特效", 3) // <= 3
    cc.addRequiredSkillByName("看破", 5)   // <= 7
    // cc.addRequiredSkillByName("超會心", 1) // <= 3

    // extra
    cc.addRequiredSkillByName("力量解放", 1) // <= 精英・泥魚龍頭盔β
    cc.addRequiredSkillByName("昏厥耐性", 2)
    cc.addRequiredSkillByName("跑者", 1)
    cc.addRequiredSkillByName("滑走強化", 1)
    cc.addRequiredSkillByName("耳塞", 1)
    // cc.addRequiredSkillByName("收刀術", 1) // <= 3
    // cc.addRequiredSkillByName("KO術", 1) // <= 3
    // cc.addRequiredSkillByName("破壞王", 1) // <= 3
    // cc.addRequiredSkillByName("炸彈客", 2) // <= 3
    // cc.addRequiredSkillByName("廣域化", 2)    // <= 5

    proc := prepareProcedure(*cc)
    proc.execute()
}

// 大剑喷气耐眠
func testExecuteProcedureGreatSword_1_2_ResistSleep(t *testing.T) {
    cc := newConstraintCollection(loadDataManager())

    // 龍熱機關式【鋼翼】 改
    cc.addExtraSlot(1, 2)
    // 衣装
    // cc.addExtraSlot(4, 1)
    // cc.addExtraSlot(2, 1)
    // 衣装
    // cc.addExtraSlot(3, 1)
    // cc.addExtraSlot(3, 1)
    cc.addExtraSlot(4, -1) // 达人珠II
    cc.addExtraSlot(4, -1) // 耐眠珠II

    addCommonDecorationLimitations(cc)

    cc.addRequiredSkillByName("無屬性強化", 1) // <= 1
    cc.addRequiredSkillByName("攻擊", 4)    // <= 7
    cc.addRequiredSkillByName("體力增強", 3)  // <= 3

    cc.addRequiredSkillByName("睡眠耐性", 1)   // <= 3

    cc.addRequiredSkillByName("集中", 3)   // <= 3
    cc.addRequiredSkillByName("匠", 3)    // <= 5
    cc.addRequiredSkillByName("弱點特效", 3) // <= 3
    cc.addRequiredSkillByName("看破", 3)   // <= 7
    // cc.addRequiredSkillByName("超會心", 1) // <= 3

    // extra
    // cc.addRequiredSkillByName("廣域化", 1)
    // cc.addRequiredSkillByName("昏厥耐性", 2)

    proc := prepareProcedure(*cc)
    proc.execute()
}

// 大剑喷气耐毒
func testExecuteProcedureGreatSword_1_2_ResistPoison(t *testing.T) {
    cc := newConstraintCollection(loadDataManager())

    // 龍熱機關式【鋼翼】 改
    cc.addExtraSlot(1, 2)
    // 衣装
    // cc.addExtraSlot(4, 1)
    // cc.addExtraSlot(2, 1)
    // 衣装
    // cc.addExtraSlot(3, 1)
    // cc.addExtraSlot(3, 1)
    cc.addExtraSlot(4, -1) // 达人珠II

    addCommonDecorationLimitations(cc)

    cc.addRequiredSkillByName("無屬性強化", 1) // <= 1
    cc.addRequiredSkillByName("攻擊", 4)    // <= 7
    cc.addRequiredSkillByName("體力增強", 3)  // <= 3

    cc.addRequiredSkillByName("毒耐性", 3)   // <= 3

    cc.addRequiredSkillByName("集中", 3)   // <= 3
    cc.addRequiredSkillByName("匠", 3)    // <= 5
    cc.addRequiredSkillByName("弱點特效", 3) // <= 3
    cc.addRequiredSkillByName("看破", 3)   // <= 7
    // cc.addRequiredSkillByName("超會心", 1) // <= 3

    // extra
    // cc.addRequiredSkillByName("廣域化", 1)
    // cc.addRequiredSkillByName("昏厥耐性", 2)
    cc.addRequiredSkillByName("力量解放", 2)
    cc.addRequiredSkillByName("昏厥耐性", 2)
    cc.addRequiredSkillByName("跑者", 1)
    cc.addRequiredSkillByName("滑走強化", 1)

    proc := prepareProcedure(*cc)
    proc.execute()
}

// 大剑喷气拔刀
func testExecuteProcedureGreatSword5_(t *testing.T) {
    // 龍熱機關式【鋼翼】 改
    cc := newConstraintCollection(loadDataManager())
    cc.addExtraSlot(1, 2)

    addCommonDecorationLimitations(cc)

    cc.addRequiredSkillByName("無屬性強化", 1) // <= 1
    // cc.addRequiredSkillByName("攻擊", 4)    // <= 7
    cc.addRequiredSkillByName("體力增強", 3)  // <= 3

    cc.addRequiredSkillByName("拔刀術【技】", 3)   // <= 3
    cc.addRequiredSkillByName("集中", 3)   // <= 3
    cc.addRequiredSkillByName("匠", 3)    // <= 5
    cc.addRequiredSkillByName("弱點特效", 3) // <= 3

    // cc.addRequiredSkillByName("看破", 5)   // <= 7
    cc.addRequiredSkillByName("超會心", 1) // <= 3

    cc.addRequiredSkillByName("精靈加護", 3) // <= 3
    cc.addRequiredSkillByName("收刀術", 3) // <= 3

    cc.addRequiredSkillByName("廣域化", 2)    // <= 5
    // cc.addRequiredSkillByName("攻擊", 1)    // <= 7
    // cc.addRequiredSkillByName("KO術", 1) // <= 3
    // cc.addRequiredSkillByName("破壞王", 1) // <= 3

    proc := prepareProcedure(*cc)
    proc.execute()
}
