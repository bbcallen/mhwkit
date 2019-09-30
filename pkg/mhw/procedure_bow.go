package mhw

import (
    "testing"
)

// 高耳拔刀龙矢: 角王弓疾风
// 物理：铠罗火
// 火：蛮颚龙射手，铠罗火
// 水：凯罗弓水，勇气猎弓
// 雷：凯罗弓雷，飞雷弓羽羽矢
// 冰：风飘弓/铠罗风飘
// 龙：龙骨弓

// 火: 蛮颚龙射手
// lv4 火 390 -> 510
func testExecuteProcedureBow1(t *testing.T) {
    cc := newConstraintCollection(loadDataManager())
    addCommonDecorationLimitations(cc)

    // cc.addExtraSlot(1, 2)

    cc.addRequiredSkillByName("體力增強", 3)  // <= 3
    cc.addRequiredSkillByName("體術", 5) // <= 5
    // cc.addRequiredSkillByName("最愛菇類", 1)   // <= 3

    // cc.addRequiredSkillByName("屬性解放／裝填擴充", 3) // <= 3
    cc.addRequiredSkillByName("火屬性攻擊強化", 4) // <= 5

    cc.addRequiredSkillByName("解放弓的蓄力階段", 1) // <= 1
    cc.addRequiredSkillByName("散彈・剛射強化", 1) // <= 1
    cc.addRequiredSkillByName("通常彈・通常箭強化", 1) // <= 1

    cc.addRequiredSkillByName("弱點特效", 3) // <= 3
    cc.addRequiredSkillByName("突破耐力上限", 1) // <= 1
    
    // cc.addRequiredSkillByName("看破", 3) // <= 7
    cc.addRequiredSkillByName("會心攻擊【屬性】", 1) // <= 1

    // cc.addRequiredSkillByName("超會心", 3) // <= 3

    // cc.addRequiredSkillByName("迴避性能", 2) // <= 2

    proc := prepareProcedure(*cc)
    proc.execute()
}

// 水: 铠罗水
// lv4 水 390 -> 490 (510)
func testExecuteProcedureBow2(t *testing.T) {
    cc := newConstraintCollection(loadDataManager())
    addCommonDecorationLimitations(cc)

    cc.addExtraSlot(3, 1)

    cc.addRequiredSkillByName("體力增強", 3)  // <= 3
    cc.addRequiredSkillByName("體術", 5) // <= 5
    // cc.addRequiredSkillByName("最愛菇類", 1)   // <= 3

    cc.addRequiredSkillByName("屬性解放／裝填擴充", 3) // <= 3
    cc.addRequiredSkillByName("水屬性攻擊強化", 3) // <= 5

    cc.addRequiredSkillByName("解放弓的蓄力階段", 1) // <= 1
    cc.addRequiredSkillByName("散彈・剛射強化", 1) // <= 1
    // cc.addRequiredSkillByName("通常彈・通常箭強化", 1) // <= 1

    cc.addRequiredSkillByName("弱點特效", 3) // <= 3
    // cc.addRequiredSkillByName("突破耐力上限", 1) // <= 1
    
    // cc.addRequiredSkillByName("看破", 3) // <= 7
    cc.addRequiredSkillByName("會心攻擊【屬性】", 1) // <= 1

    // cc.addRequiredSkillByName("超會心", 3) // <= 3

    // cc.addRequiredSkillByName("迴避性能", 2) // <= 2

    proc := prepareProcedure(*cc)
    proc.execute()
}


// 雷: 铠罗雷
// lv4 雷 360 -> 460
func testExecuteProcedureBow3(t *testing.T) {
    cc := newConstraintCollection(loadDataManager())
    addCommonDecorationLimitations(cc)

    cc.addExtraSlot(3, 1)

    cc.addRequiredSkillByName("體力增強", 3)  // <= 3
    cc.addRequiredSkillByName("體術", 5) // <= 5
    // cc.addRequiredSkillByName("最愛菇類", 1)   // <= 3

    cc.addRequiredSkillByName("屬性解放／裝填擴充", 3) // <= 3
    cc.addRequiredSkillByName("雷屬性攻擊強化", 3) // <= 5

    cc.addRequiredSkillByName("解放弓的蓄力階段", 1) // <= 1
    cc.addRequiredSkillByName("散彈・剛射強化", 1) // <= 1
    // cc.addRequiredSkillByName("通常彈・通常箭強化", 1) // <= 1

    cc.addRequiredSkillByName("弱點特效", 3) // <= 3
    // cc.addRequiredSkillByName("突破耐力上限", 1) // <= 1
    
    // cc.addRequiredSkillByName("看破", 3) // <= 7
    cc.addRequiredSkillByName("會心攻擊【屬性】", 1) // <= 1

    // cc.addRequiredSkillByName("超會心", 3) // <= 3

    // cc.addRequiredSkillByName("迴避性能", 2) // <= 2

    proc := prepareProcedure(*cc)
    proc.execute()
}

// 冰: 风飘弓
// lv4 冰 390 ->
func testExecuteProcedureBow4(t *testing.T) {
    cc := newConstraintCollection(loadDataManager())
    addCommonDecorationLimitations(cc)

    // cc.addExtraSlot(1, 2)

    cc.addRequiredSkillByName("體力增強", 3)  // <= 3
    cc.addRequiredSkillByName("體術", 5) // <= 5
    // cc.addRequiredSkillByName("最愛菇類", 1)   // <= 3

    // cc.addRequiredSkillByName("屬性解放／裝填擴充", 3) // <= 3
    cc.addRequiredSkillByName("冰屬性攻擊強化", 3) // <= 5

    cc.addRequiredSkillByName("解放弓的蓄力階段", 1) // <= 1
    cc.addRequiredSkillByName("散彈・剛射強化", 1) // <= 1
    cc.addRequiredSkillByName("通常彈・通常箭強化", 1) // <= 1

    cc.addRequiredSkillByName("弱點特效", 3) // <= 3
    cc.addRequiredSkillByName("突破耐力上限", 0) // <= 1
    
    // cc.addRequiredSkillByName("看破", 3) // <= 7
    cc.addRequiredSkillByName("會心攻擊【屬性】", 1) // <= 1

    // cc.addRequiredSkillByName("超會心", 3) // <= 3

    // cc.addRequiredSkillByName("迴避性能", 2) // <= 2

    proc := prepareProcedure(*cc)
    proc.execute()
}

// 龙: 龙骨弓III
// lv4 龙 420 ->
func testExecuteProcedureBow5(t *testing.T) {
    cc := newConstraintCollection(loadDataManager())
    addCommonDecorationLimitations(cc)

    cc.addExtraSlot(1, 1)
    cc.addExtraSlot(1, 1)

    cc.addRequiredSkillByName("體力增強", 3)  // <= 3
    cc.addRequiredSkillByName("體術", 5) // <= 5
    // cc.addRequiredSkillByName("最愛菇類", 1)   // <= 3

    // cc.addRequiredSkillByName("屬性解放／裝填擴充", 3) // <= 3
    cc.addRequiredSkillByName("龍屬性攻擊強化", 4) // <= 5

    cc.addRequiredSkillByName("解放弓的蓄力階段", 1) // <= 1
    cc.addRequiredSkillByName("散彈・剛射強化", 1) // <= 1
    cc.addRequiredSkillByName("通常彈・通常箭強化", 1) // <= 1

    cc.addRequiredSkillByName("弱點特效", 3) // <= 3
    cc.addRequiredSkillByName("突破耐力上限", 1) // <= 1
    
    cc.addRequiredSkillByName("看破", 2) // <= 7
    cc.addRequiredSkillByName("會心攻擊【屬性】", 1) // <= 1

    // cc.addRequiredSkillByName("超會心", 3) // <= 3
    // cc.addRequiredSkillByName("迴避性能", 2) // <= 2

    cc.addRequiredSkillByName("廣域化", 2) // <= 2
    cc.addRequiredSkillByName("整備", 2) // <= 2

    proc := prepareProcedure(*cc)
    proc.execute()
}

// 物理: 凯罗火
func testExecuteProcedureBow6(t *testing.T) {
    cc := newConstraintCollection(loadDataManager())
    addCommonDecorationLimitations(cc)

    cc.addExtraSlot(3, 1)

    cc.addRequiredSkillByName("體力增強", 3)  // <= 3
    cc.addRequiredSkillByName("體術", 5) // <= 5
    // cc.addRequiredSkillByName("最愛菇類", 1)   // <= 3

    cc.addRequiredSkillByName("無屬性強化", 1) // <= 1

    cc.addRequiredSkillByName("解放弓的蓄力階段", 1) // <= 1
    cc.addRequiredSkillByName("散彈・剛射強化", 1) // <= 1
    cc.addRequiredSkillByName("通常彈・通常箭強化", 1) // <= 1

    cc.addRequiredSkillByName("弱點特效", 3) // <= 3
    cc.addRequiredSkillByName("突破耐力上限", 1) // <= 1
    
    cc.addRequiredSkillByName("看破", 7) // <= 7
    // cc.addRequiredSkillByName("攻擊", 4)    // <= 7
    // cc.addRequiredSkillByName("超會心", 3) // <= 3
    // cc.addRequiredSkillByName("迴避性能", 2) // <= 2

    // cc.addRequiredSkillByName("廣域化", 2) // <= 2
    // cc.addRequiredSkillByName("整備", 2) // <= 2

    proc := prepareProcedure(*cc)
    proc.execute()
}

