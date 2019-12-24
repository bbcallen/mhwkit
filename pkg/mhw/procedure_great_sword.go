package mhw

import (
    "testing"
)

// 冥赤龙大剑 日常
func testExecuteProcedureGreatSword_4(t *testing.T) {
    cc := newConstraintCollection(loadDataManager())

    cc.addExtraSlot(4, 1)

    cc.addExtraSlot(4, -1) // 达人II
    cc.addRequiredSkillByName("看破", -2)
    cc.addExtraSlot(4, -1)
    cc.addRequiredSkillByName("挑戰者", -1)
    cc.addRequiredSkillByName("精靈加護", -1)
    cc.addExtraSlot(4, -1)
    cc.addRequiredSkillByName("挑戰者", -1)
    cc.addRequiredSkillByName("體力增強", -1)

    addCommonDecorationLimitations(cc)

    cc.addRequiredSkillByName("攻擊", 4)    // <= 7
    cc.addRequiredSkillByName("體力增強", 3)  // <= 3

    cc.addRequiredSkillByName("集中", 3)   // <= 3
    cc.addRequiredSkillByName("弱點特效", 3) // <= 3
    cc.addRequiredSkillByName("看破", 7)   // <= 7
    cc.addRequiredSkillByName("超會心", 3) // <= 3

    // cc.addRequiredSkillByName("匠", 3)

    // extra
    // cc.addRequiredSkillByName("心眼／彈道強化", 1)
    cc.addRequiredSkillByName("昏厥耐性", 3)
    cc.addRequiredSkillByName("精靈加護", 3)
    cc.addRequiredSkillByName("挑戰者", 2)
    cc.addRequiredSkillByName("體力回復量UP", 1)
    // cc.addRequiredSkillByName("轉禍為福", 1)
    // cc.addRequiredSkillByName("怨恨", 2)
    
    
    // cc.addRequiredSkillByName("減輕膽怯", 1)
    // cc.addRequiredSkillByName("跑者", 1)

    proc := prepareProcedure(*cc)
    proc.execute()
}

// 冥赤龙大剑 耐Any
func testExecuteProcedureGreatSword_4_AnyCharm(t *testing.T) {
    cc := newConstraintCollection(loadDataManager())

    cc.addExtraSlot(4, 1)
    cc.addRequiredSkillByName("任意護石", 1)

    cc.addExtraSlot(4, -1) // 达人II
    cc.addRequiredSkillByName("看破", -2)
    cc.addExtraSlot(4, -1)
    cc.addRequiredSkillByName("挑戰者", -1)
    cc.addRequiredSkillByName("精靈加護", -1)
    cc.addExtraSlot(4, -1)
    cc.addRequiredSkillByName("挑戰者", -1)
    cc.addRequiredSkillByName("體力增強", -1)
    cc.addExtraSlot(4, -1)
    cc.addRequiredSkillByName("集中", -1)
    cc.addRequiredSkillByName("體力增強", -1)
    // cc.addExtraSlot(4, -1)
    // cc.addRequiredSkillByName("減輕膽怯", -1)
    // cc.addRequiredSkillByName("體力增強", -1)

    addCommonDecorationLimitations(cc)

    cc.addRequiredSkillByName("攻擊", 4)    // <= 7
    cc.addRequiredSkillByName("體力增強", 3)  // <= 3

    cc.addRequiredSkillByName("集中", 3)   // <= 3
    cc.addRequiredSkillByName("弱點特效", 3) // <= 3
    cc.addRequiredSkillByName("看破", 7)   // <= 7
    cc.addRequiredSkillByName("超會心", 3) // <= 3

    // cc.addRequiredSkillByName("匠", 3)

    // extra
    // cc.addRequiredSkillByName("心眼／彈道強化", 1)
    cc.addRequiredSkillByName("地質學", 1)
    cc.addRequiredSkillByName("昏厥耐性", 3)
    cc.addRequiredSkillByName("挑戰者", 2)
    cc.addRequiredSkillByName("耳塞", 1)
    // cc.addRequiredSkillByName("體力回復量UP", 1)
    // cc.addRequiredSkillByName("轉禍為福", 1)
    
    
    
    // cc.addRequiredSkillByName("減輕膽怯", 1)
    // cc.addRequiredSkillByName("跑者", 1)

    proc := prepareProcedure(*cc)
    proc.execute()
}

// 灭尽大剑 耐Any
func testExecuteProcedureGreatSword_1_1_AnyCharm(t *testing.T) {
    cc := newConstraintCollection(loadDataManager())

    cc.addExtraSlot(1, 1)
    cc.addRequiredSkillByName("任意護石", 1)

    cc.addExtraSlot(4, -1) // 达人II
    cc.addRequiredSkillByName("看破", -2)
    cc.addExtraSlot(4, -1) // 耐绝II
    cc.addRequiredSkillByName("昏厥耐性", -2)

    addCommonDecorationLimitations(cc)

    cc.addRequiredSkillByName("攻擊", 4)    // <= 7
    cc.addRequiredSkillByName("體力增強", 3)  // <= 3

    cc.addRequiredSkillByName("集中", 3)   // <= 3
    cc.addRequiredSkillByName("弱點特效", 3) // <= 3
    cc.addRequiredSkillByName("看破", 7)   // <= 7
    cc.addRequiredSkillByName("超會心", 3) // <= 3

    cc.addRequiredSkillByName("地質學", 1) // <= 3

    // extra
    cc.addRequiredSkillByName("昏厥耐性", 2)
    cc.addRequiredSkillByName("精靈加護", 1)

    // cc.addRequiredSkillByName("心眼／彈道強化", 1)
    // cc.addRequiredSkillByName("體力回復量UP", 1)
    // cc.addRequiredSkillByName("怨恨", 2)    
    // cc.addRequiredSkillByName("減輕膽怯", 1)

    proc := prepareProcedure(*cc)
    proc.execute()
}

// 采集
func testExecuteProcedureGreatSword_4_Gathering(t *testing.T) {
    cc := newConstraintCollection(loadDataManager())

    // 巨牛大斧II
    cc.addExtraSlot(4, 1)
    // 衣装
    // cc.addExtraSlot(4, 1)
    // cc.addExtraSlot(2, 1)
    // 衣装
    // cc.addExtraSlot(3, 1)
    // cc.addExtraSlot(3, 1)
    // cc.addExtraSlot(4, -1) // 达人II
    // cc.addRequiredSkillByName("看破", -2)

    addCommonDecorationLimitations(cc)

    cc.addExtraSlot(4, -1) // 潛伏珠III
    cc.addRequiredSkillByName("潛伏", -3) // <= 3
    cc.addExtraSlot(4, -1) // 威嚇珠III
    cc.addRequiredSkillByName("威嚇", -3) // <= 3
    cc.addExtraSlot(4, -1) // 地學珠II
    cc.addRequiredSkillByName("地質學", -2) // <= 3
    cc.addExtraSlot(4, -1) // 無食珠II
    cc.addRequiredSkillByName("飢餓耐性", -2) // <= 3

    cc.addRequiredSkillByName("地質學", 3) // <= 3
    cc.addRequiredSkillByName("威嚇", 3) // <= 3
    cc.addRequiredSkillByName("潛伏", 3) // <= 3
    cc.addRequiredSkillByName("耳塞", 5) // <= 5
    cc.addRequiredSkillByName("跑者", 3) // <= 3
    cc.addRequiredSkillByName("察覺", 1) // <= 1
    cc.addRequiredSkillByName("耐力急速回復", 3) // <= 3
    cc.addRequiredSkillByName("飢餓耐性", 3) // <= 3
    cc.addRequiredSkillByName("整備", 3) // <= 3

    // extra
    cc.addRequiredSkillByName("適應瘴氣環境", 1) // <= 1
    cc.addRequiredSkillByName("熱傷害無效", 1) // <= 1
    cc.addRequiredSkillByName("寒冷耐性", 1) // <= 1
    cc.addRequiredSkillByName("適應水場・深雪", 1) // <= 1
    cc.addRequiredSkillByName("突破耐力上限", 1) // <= 1
    cc.addRequiredSkillByName("體術", 1) // <= 5

    // cc.addRequiredSkillByName("攀岩者", 1) // <= 1    
    // cc.addRequiredSkillByName("採集達人", 1) // <= 1
    // cc.addRequiredSkillByName("滿足感", 1) // <= 1
    // cc.addRequiredSkillByName("最愛菇類", 1) // <= 1
    // cc.addRequiredSkillByName("屬性異常狀態耐性", 1) // <= 3
    // cc.addRequiredSkillByName("探索者的幸運", 1) // <= 1
    // cc.addRequiredSkillByName("導蟲反應距離UP", 1) // <= 1

    proc := prepareProcedure(*cc)
    proc.execute()
}

// 冰翼灵羽 对炎王龙 耐火 耐爆
func testExecuteProcedureGreatSword_4_VsTeostra(t *testing.T) {
    cc := newConstraintCollection(loadDataManager())

    cc.addExtraSlot(4, 1)

    cc.addExtraSlot(4, -1) // 达人II
    cc.addRequiredSkillByName("看破", -2)
    cc.addExtraSlot(4, -1)
    cc.addRequiredSkillByName("火耐性", -2)
    cc.addExtraSlot(4, -1)
    cc.addRequiredSkillByName("爆破異常狀態的耐性", -2)
    // cc.addExtraSlot(4, -1)
    // cc.addRequiredSkillByName("體力增強", -1)
    // cc.addRequiredSkillByName("精靈加護", -1)

    addCommonDecorationLimitations(cc)

    cc.addRequiredSkillByName("火耐性", 3)
    cc.addRequiredSkillByName("爆破異常狀態的耐性", 3)

    // cc.addRequiredSkillByName("攻擊", 4)    // <= 7
    cc.addRequiredSkillByName("體力增強", 3)  // <= 3

    cc.addRequiredSkillByName("集中", 3)   // <= 3
    cc.addRequiredSkillByName("弱點特效", 3) // <= 3
    cc.addRequiredSkillByName("看破", 7)   // <= 7
    cc.addRequiredSkillByName("超會心", 3) // <= 3

    cc.addRequiredSkillByName("匠", 1)

    // extra
    // cc.addRequiredSkillByName("適應水場・深雪", 1)
    cc.addRequiredSkillByName("心眼／彈道強化", 1)
    cc.addRequiredSkillByName("昏厥耐性", 2)
    cc.addRequiredSkillByName("精靈加護", 2)
    // cc.addRequiredSkillByName("減輕膽怯", 1)
    // cc.addRequiredSkillByName("轉禍為福", 1)
    // cc.addRequiredSkillByName("體力回復量UP", 1)
    // cc.addRequiredSkillByName("跑者", 1)
    // cc.addRequiredSkillByName("跳躍鐵人", 1)
    // cc.addRequiredSkillByName("火屬性攻擊強化", 1)
    // cc.addRequiredSkillByName("道具使用強化", 1)
    // cc.addRequiredSkillByName("廣域化", 1)
    // cc.addRequiredSkillByName("耳塞", 1)
    // cc.addRequiredSkillByName("力量解放", 1)
    // cc.addRequiredSkillByName("滑走強化", 1)

    proc := prepareProcedure(*cc)
    proc.execute()
}

// 冰翼灵羽 对炎妃龙 耐火 热伤害无效
func testExecuteProcedureGreatSword_4_VsLunastra(t *testing.T) {
    cc := newConstraintCollection(loadDataManager())

    cc.addExtraSlot(4, 1)

    cc.addExtraSlot(4, -1) // 达人II
    cc.addRequiredSkillByName("看破", -2)
    cc.addExtraSlot(4, -1)
    cc.addRequiredSkillByName("火耐性", -2)
    cc.addExtraSlot(4, -1)
    cc.addRequiredSkillByName("集中", -1)
    cc.addRequiredSkillByName("體力增強", -1)
    // cc.addExtraSlot(4, -1)
    // cc.addRequiredSkillByName("減輕膽怯", -1)
    // cc.addRequiredSkillByName("體力增強", -1)
    // cc.addExtraSlot(4, -1)
    // cc.addRequiredSkillByName("挑戰者", -1)
    // cc.addRequiredSkillByName("精靈加護", -1)
    // cc.addExtraSlot(4, -1)
    // cc.addRequiredSkillByName("挑戰者", -1)
    // cc.addRequiredSkillByName("體力增強", -1)

    addCommonDecorationLimitations(cc)

    cc.addRequiredSkillByName("熱傷害無效", 1)
    cc.addRequiredSkillByName("火耐性", 3)

    cc.addRequiredSkillByName("攻擊", 4)    // <= 7
    cc.addRequiredSkillByName("體力增強", 3)  // <= 3

    cc.addRequiredSkillByName("集中", 3)   // <= 3
    cc.addRequiredSkillByName("弱點特效", 3) // <= 3
    cc.addRequiredSkillByName("看破", 7)   // <= 7
    cc.addRequiredSkillByName("超會心", 3) // <= 3

    // cc.addRequiredSkillByName("匠", 3)

    // extra
    // cc.addRequiredSkillByName("心眼／彈道強化", 1)
    cc.addRequiredSkillByName("地質學", 1)
    cc.addRequiredSkillByName("昏厥耐性", 3)
    cc.addRequiredSkillByName("精靈加護", 1)
    cc.addRequiredSkillByName("體力回復量UP", 1)
    
    // cc.addRequiredSkillByName("耳塞", 1)
    // cc.addRequiredSkillByName("挑戰者", 2)
    // cc.addRequiredSkillByName("耳塞", 1)
    // cc.addRequiredSkillByName("轉禍為福", 1)
    
    
    
    // cc.addRequiredSkillByName("減輕膽怯", 1)
    // cc.addRequiredSkillByName("跑者", 1)

    proc := prepareProcedure(*cc)
    proc.execute()
}

// 冰翼灵羽 对碎龙 冰攻 耐爆
func testExecuteProcedureGreatSword_4_VsBrachydios(t *testing.T) {
    cc := newConstraintCollection(loadDataManager())

    cc.addExtraSlot(4, 1)
    //
    // cc.addExtraSlot(4, -1) // 达人II
    // cc.addRequiredSkillByName("看破", -2)
    cc.addExtraSlot(4, -1) // 达人II
    cc.addRequiredSkillByName("爆破異常狀態的耐性", -2)

    addCommonDecorationLimitations(cc)

    // cc.addRequiredSkillByName("攻擊", 4)    // <= 7
    cc.addRequiredSkillByName("體力增強", 3)  // <= 3

    cc.addRequiredSkillByName("寒氣鍊成", 1)  // <= 1
    cc.addRequiredSkillByName("會心擊【屬性】", 1)  // <= 1
    
    cc.addRequiredSkillByName("集中", 3)   // <= 3
    cc.addRequiredSkillByName("收刀術", 3)   // <= 3
    cc.addRequiredSkillByName("拔刀術【技】", 3)   // <= 3

    cc.addRequiredSkillByName("超會心", 3) // <= 3

    cc.addRequiredSkillByName("爆破異常狀態的耐性", 3)
    cc.addRequiredSkillByName("冰屬性攻擊強化", 3)

    // extra
    // cc.addRequiredSkillByName("適應水場・深雪", 1)
    // cc.addRequiredSkillByName("心眼／彈道強化", 1)
    // cc.addRequiredSkillByName("昏厥耐性", 2)
    cc.addRequiredSkillByName("精靈加護", 2)
    cc.addRequiredSkillByName("減輕膽怯", 1)
    // cc.addRequiredSkillByName("轉禍為福", 1)
    cc.addRequiredSkillByName("體力回復量UP", 1)
    // cc.addRequiredSkillByName("跑者", 1)
    // cc.addRequiredSkillByName("跳躍鐵人", 1)
    // cc.addRequiredSkillByName("火屬性攻擊強化", 1)
    // cc.addRequiredSkillByName("道具使用強化", 1)
    // cc.addRequiredSkillByName("廣域化", 1)
    // cc.addRequiredSkillByName("耳塞", 1)
    // cc.addRequiredSkillByName("力量解放", 1)
    // cc.addRequiredSkillByName("滑走強化", 1)

    proc := prepareProcedure(*cc)
    proc.execute()
}

// 泻湖巨像II 对斩龙 水攻
func testExecuteProcedureGreatSword_4_VsGlavenus(t *testing.T) {
    cc := newConstraintCollection(loadDataManager())

    cc.addExtraSlot(2, 1)
    cc.addExtraSlot(1, 1)
    //
    // cc.addExtraSlot(4, -1) // 达人II
    // cc.addRequiredSkillByName("看破", -2)

    addCommonDecorationLimitations(cc)

    cc.addRequiredSkillByName("攻擊", 4)    // <= 7
    cc.addRequiredSkillByName("體力增強", 3)  // <= 3

    cc.addRequiredSkillByName("寒氣鍊成", 1)  // <= 1
    cc.addRequiredSkillByName("會心擊【屬性】", 1)  // <= 1
    
    cc.addRequiredSkillByName("集中", 3)   // <= 3
    cc.addRequiredSkillByName("收刀術", 3)   // <= 3
    cc.addRequiredSkillByName("拔刀術【技】", 3)   // <= 3

    cc.addRequiredSkillByName("超會心", 3) // <= 3

    // cc.addRequiredSkillByName("爆破異常狀態的耐性", 3)

    // extra
    cc.addRequiredSkillByName("匠", 1)    // <= 5
    // cc.addRequiredSkillByName("適應水場・深雪", 1)
    // cc.addRequiredSkillByName("心眼／彈道強化", 1)
    cc.addRequiredSkillByName("昏厥耐性", 2)
    cc.addRequiredSkillByName("精靈加護", 2)
    cc.addRequiredSkillByName("減輕膽怯", 1)
    // cc.addRequiredSkillByName("轉禍為福", 1)
    // cc.addRequiredSkillByName("體力回復量UP", 1)
    // cc.addRequiredSkillByName("跑者", 1)
    // cc.addRequiredSkillByName("跳躍鐵人", 1)
    // cc.addRequiredSkillByName("火屬性攻擊強化", 1)
    // cc.addRequiredSkillByName("道具使用強化", 1)
    // cc.addRequiredSkillByName("廣域化", 1)
    // cc.addRequiredSkillByName("耳塞", 1)
    // cc.addRequiredSkillByName("力量解放", 1)
    // cc.addRequiredSkillByName("滑走強化", 1)

    proc := prepareProcedure(*cc)
    proc.execute()
}

// 斩龙之炎II 对冰咒龙 火攻 冰耐
func testExecuteProcedureGreatSword_4_VsVelkhana(t *testing.T) {
    cc := newConstraintCollection(loadDataManager())

    cc.addExtraSlot(1, 2)
    //
    // cc.addExtraSlot(4, -1) // 达人II
    // cc.addRequiredSkillByName("看破", -2)
    cc.addExtraSlot(4, -1) // 冰耐性II
    cc.addRequiredSkillByName("冰耐性", -2)

    addCommonDecorationLimitations(cc)

    // cc.addRequiredSkillByName("攻擊", 4)    // <= 7
    cc.addRequiredSkillByName("體力增強", 3)  // <= 3

    cc.addRequiredSkillByName("寒氣鍊成", 1)  // <= 1
    cc.addRequiredSkillByName("會心擊【屬性】", 1)  // <= 1
    
    cc.addRequiredSkillByName("集中", 3)   // <= 3
    cc.addRequiredSkillByName("收刀術", 3)   // <= 3
    cc.addRequiredSkillByName("拔刀術【技】", 3)   // <= 3

    cc.addRequiredSkillByName("超會心", 3) // <= 3

    cc.addRequiredSkillByName("冰耐性", 3)
    cc.addRequiredSkillByName("火屬性攻擊強化", 3)

    // extra
    // cc.addRequiredSkillByName("匠", 1)    // <= 5
    // cc.addRequiredSkillByName("適應水場・深雪", 1)
    // cc.addRequiredSkillByName("心眼／彈道強化", 1)
    // cc.addRequiredSkillByName("昏厥耐性", 2)
    cc.addRequiredSkillByName("精靈加護", 2)
    cc.addRequiredSkillByName("減輕膽怯", 1)
    // cc.addRequiredSkillByName("轉禍為福", 1)
    // cc.addRequiredSkillByName("體力回復量UP", 1)
    cc.addRequiredSkillByName("跑者", 1)
    // cc.addRequiredSkillByName("跳躍鐵人", 1)
    // cc.addRequiredSkillByName("火屬性攻擊強化", 1)
    // cc.addRequiredSkillByName("道具使用強化", 1)
    // cc.addRequiredSkillByName("廣域化", 1)
    // cc.addRequiredSkillByName("耳塞", 1)
    // cc.addRequiredSkillByName("力量解放", 1)
    // cc.addRequiredSkillByName("滑走強化", 1)

    proc := prepareProcedure(*cc)
    proc.execute()
}

// 冰翼灵羽 对金狮子
func testExecuteProcedureGreatSword_1_2_VsRajang(t *testing.T) {
    cc := newConstraintCollection(loadDataManager())

    cc.addExtraSlot(4, 1)

    cc.addExtraSlot(4, -1)
    cc.addRequiredSkillByName("集中", -1)
    cc.addRequiredSkillByName("整備", -1)
    cc.addExtraSlot(4, -1)
    cc.addRequiredSkillByName("拔刀術【技】", -1)
    cc.addRequiredSkillByName("精靈加護", -1)
    cc.addExtraSlot(4, -1)
    cc.addRequiredSkillByName("減輕膽怯", -1)
    cc.addRequiredSkillByName("體力增強", -1)
    cc.addExtraSlot(4, -1)
    cc.addRequiredSkillByName("耐震", -1)
    cc.addRequiredSkillByName("體力回復量UP", -1)

    // cc.addExtraSlot(4, -1)
    // cc.addRequiredSkillByName("挑戰者", -1)
    // cc.addRequiredSkillByName("精靈加護", -1)
    // cc.addExtraSlot(4, -1)
    // cc.addRequiredSkillByName("挑戰者", -1)
    // cc.addRequiredSkillByName("體力增強", -1)
    // cc.addExtraSlot(4, -1)
    // cc.addRequiredSkillByName("減輕膽怯", -1)
    // cc.addRequiredSkillByName("精靈加護", -1)
    // cc.addExtraSlot(4, -1)
    // cc.addRequiredSkillByName("昏厥耐性", -2)
    // cc.addExtraSlot(4, -1)
    // cc.addRequiredSkillByName("不屈", -1)
    // cc.addRequiredSkillByName("匠", -1)

    addCommonDecorationLimitations(cc)

    // cc.addRequiredSkillByName("無屬性強化", 1) // <= 1
    // cc.addRequiredSkillByName("攻擊", 4)    // <= 7
    cc.addRequiredSkillByName("體力增強", 3)  // <= 3

    cc.addRequiredSkillByName("寒氣鍊成", 1)  // <= 1
    cc.addRequiredSkillByName("會心擊【屬性】", 1)  // <= 1
    
    cc.addRequiredSkillByName("集中", 3)   // <= 3
    cc.addRequiredSkillByName("收刀術", 3)   // <= 3
    cc.addRequiredSkillByName("拔刀術【技】", 3)   // <= 3

    // cc.addRequiredSkillByName("匠", 3)    // <= 5
    // cc.addRequiredSkillByName("弱點特效", 3) // <= 3
    // cc.addRequiredSkillByName("看破", 7)   // <= 7
    cc.addRequiredSkillByName("超會心", 3) // <= 3

    cc.addRequiredSkillByName("耐震", 3) // <= 3
    // cc.addRequiredSkillByName("昏厥耐性", 3)
    // cc.addRequiredSkillByName("心眼／彈道強化", 1)

    // size-4
    cc.addRequiredSkillByName("減輕膽怯", 2)
    cc.addRequiredSkillByName("精靈加護", 3)
    cc.addRequiredSkillByName("整備", 1)
    cc.addRequiredSkillByName("體力回復量UP", 1)

    // extra
    cc.addRequiredSkillByName("跑者", 1)
    cc.addRequiredSkillByName("轉禍為福", 2)
    // cc.addRequiredSkillByName("跳躍鐵人", 1)
    // cc.addRequiredSkillByName("火屬性攻擊強化", 1)
    // cc.addRequiredSkillByName("道具使用強化", 1)
    // cc.addRequiredSkillByName("廣域化", 1)
    // cc.addRequiredSkillByName("耳塞", 1)
    // cc.addRequiredSkillByName("力量解放", 1)
    // cc.addRequiredSkillByName("滑走強化", 1)

    proc := prepareProcedure(*cc)
    proc.execute()
}

// 对天地煌啼龙
func testExecuteProcedureGreatSword_4_VsSharaIshvalda(t *testing.T) {
    cc := newConstraintCollection(loadDataManager())

    // 冰翼灵羽
    cc.addExtraSlot(4, 1)
    //
    cc.addExtraSlot(4, -1) // 达人II
    cc.addRequiredSkillByName("看破", -2)

    addCommonDecorationLimitations(cc)

    cc.addRequiredSkillByName("攻擊", 4)    // <= 7
    cc.addRequiredSkillByName("體力增強", 3)  // <= 3

    cc.addRequiredSkillByName("集中", 3)   // <= 3
    cc.addRequiredSkillByName("匠", 0)    // <= 5
    cc.addRequiredSkillByName("弱點特效", 3) // <= 3
    cc.addRequiredSkillByName("看破", 7)   // <= 7
    cc.addRequiredSkillByName("超會心", 3) // <= 3

    cc.addRequiredSkillByName("適應水場・深雪", 1) // <= 1

    // extra
    // cc.addRequiredSkillByName("適應水場・深雪", 1)
    // cc.addRequiredSkillByName("心眼／彈道強化", 1)
    cc.addRequiredSkillByName("昏厥耐性", 2)
    cc.addRequiredSkillByName("精靈加護", 1)
    cc.addRequiredSkillByName("體力回復量UP", 1)
    cc.addRequiredSkillByName("跑者", 1)
    // cc.addRequiredSkillByName("跳躍鐵人", 1)
    // cc.addRequiredSkillByName("火屬性攻擊強化", 1)
    // cc.addRequiredSkillByName("道具使用強化", 1)
    // cc.addRequiredSkillByName("廣域化", 1)
    // cc.addRequiredSkillByName("耳塞", 1)
    // cc.addRequiredSkillByName("力量解放", 1)
    // cc.addRequiredSkillByName("滑走強化", 1)

    proc := prepareProcedure(*cc)
    proc.execute()
}

// 对冥赤龙
func testExecuteProcedureGreatSword_4_VsSafiJiiva(t *testing.T) {
    cc := newConstraintCollection(loadDataManager())

    // 冥赤龙大剑
    cc.addExtraSlot(4, 1)

    cc.addExtraSlot(4, -1) // 达人II
    cc.addRequiredSkillByName("看破", -2)
    cc.addExtraSlot(4, -1)
    cc.addRequiredSkillByName("火耐性", -2)
    cc.addExtraSlot(4, -1) // 耐绝II
    cc.addRequiredSkillByName("昏厥耐性", -2)
    cc.addExtraSlot(4, -1) // 重击+体力
    cc.addRequiredSkillByName("體力增強", -1)  // <= 3
    cc.addRequiredSkillByName("破壞王", -1) // <= 3

    addCommonDecorationLimitations(cc)

    cc.addRequiredSkillByName("攻擊", 3)    // <= 7
    cc.addRequiredSkillByName("體力增強", 3)  // <= 3

    cc.addRequiredSkillByName("集中", 3)   // <= 3
    cc.addRequiredSkillByName("弱點特效", 3) // <= 3
    cc.addRequiredSkillByName("看破", 7)   // <= 7
    cc.addRequiredSkillByName("超會心", 3) // <= 3

    cc.addRequiredSkillByName("破壞王", 3) // <= 3
    cc.addRequiredSkillByName("火耐性", 3)
    // extra
    cc.addRequiredSkillByName("昏厥耐性", 3)
    cc.addRequiredSkillByName("爆破屬性強化", 2)
    // cc.addRequiredSkillByName("迴避性能", 1)

    // cc.addRequiredSkillByName("心眼／彈道強化", 1)
    // cc.addRequiredSkillByName("體力回復量UP", 1)
    // cc.addRequiredSkillByName("怨恨", 2)    
    // cc.addRequiredSkillByName("減輕膽怯", 1)

    proc := prepareProcedure(*cc)
    proc.execute()
}
