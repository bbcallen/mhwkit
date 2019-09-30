package mhw

import (
    "testing"
)

// 铠罗突击弩・贼 
// 散弹
func testExecuteProcedureHeavyBowgun1(t *testing.T) {
    cc := newConstraintCollection(loadDataManager())
    addCommonDecorationLimitations(cc)
        
    // cc.addExtraSlot(2, 1)
    cc.addRequiredSkillByName("屬性解放／裝填擴充", 3) // <= 3
    cc.addRequiredSkillByName("散彈・剛射強化", 1) // <= 1
    cc.addRequiredSkillByName("體力增強", 3)  // <= 3
    cc.addRequiredSkillByName("無屬性強化", 1) // <= 1

    cc.addRequiredSkillByName("防禦性能", 3)    // <= 5
    cc.addRequiredSkillByName("弱點特效", 3) // <= 3
    cc.addRequiredSkillByName("超會心", 3) // <= 3
    // cc.addRequiredSkillByName("攻擊", 3)    // <= 7

    cc.addRequiredSkillByName("昏厥耐性", 3) // <= 3
    cc.addRequiredSkillByName("迴避距離UP", 2) // <= 3
    
    // cc.addRequiredSkillByName("利刃／彈藥節約", 1)  // <= 1
    // cc.addRequiredSkillByName("攻擊", 5)    // <= 7
    // cc.addRequiredSkillByName("精靈加護", 3) // <= 3
    // cc.addRequiredSkillByName("無傷", 1) // <= 3
    // cc.addRequiredSkillByName("減輕膽怯", 2) // <= 3
    

    proc := prepareProcedure(*cc)
    proc.execute()
}


// 铠罗 援击
// 贯3
func testExecuteProcedureHeavyBowgun2(t *testing.T) {
    cc := newConstraintCollection(loadDataManager())
    addCommonDecorationLimitations(cc)
        
    cc.addExtraSlot(3, 1)

    cc.addRequiredSkillByName("屬性解放／裝填擴充", 3) // <= 3
    cc.addRequiredSkillByName("貫通彈・龍之箭強化", 1) // <= 1
    cc.addRequiredSkillByName("心眼／彈道強化", 1)    // <= 1
    cc.addRequiredSkillByName("體力增強", 3)  // <= 3

    cc.addRequiredSkillByName("防禦性能", 3) // <= 5
    cc.addRequiredSkillByName("防禦強化", 1) // <= 1

    cc.addRequiredSkillByName("弱點特效", 3) // <= 3
    cc.addRequiredSkillByName("看破", 2) // <= 7
    //cc.addRequiredSkillByName("超會心", 1) // <= 3

    cc.addRequiredSkillByName("攻擊", 4)    // <= 7

    // cc.addRequiredSkillByName("昏厥耐性", 3) // <= 3
    // cc.addRequiredSkillByName("迴避距離UP", 2) // <= 3

    cc.addRequiredSkillByName("挑戰者", 1)  // <= 5
    cc.addRequiredSkillByName("廣域化", 2) // <= 5
    
    // cc.addRequiredSkillByName("利刃／彈藥節約", 1)  // <= 1
    // cc.addRequiredSkillByName("精靈加護", 3) // <= 3
    // cc.addRequiredSkillByName("無傷", 1) // <= 3
    // cc.addRequiredSkillByName("熱傷害無效", 2) // <= 3
    

    proc := prepareProcedure(*cc)
    proc.execute()
}

// 风飘弩
// 贯3 王冥灯
func testExecuteProcedureHeavyBowgun3(t *testing.T) {
    cc := newConstraintCollection(loadDataManager())
    addCommonDecorationLimitations(cc)
        
    cc.addExtraSlot(1, 1)

    cc.addRequiredSkillByName("屬性解放／裝填擴充", 3) // <= 3
    cc.addRequiredSkillByName("貫通彈・龍之箭強化", 1) // <= 1
    cc.addRequiredSkillByName("心眼／彈道強化", 1)    // <= 1
    cc.addRequiredSkillByName("體力增強", 3)  // <= 3
    cc.addRequiredSkillByName("熱傷害無效", 1) // <= 1

    cc.addRequiredSkillByName("防禦性能", 3) // <= 5
    cc.addRequiredSkillByName("防禦強化", 1) // <= 1

    cc.addRequiredSkillByName("弱點特效", 3) // <= 3
    //cc.addRequiredSkillByName("看破", 1) // <= 7
    //cc.addRequiredSkillByName("超會心", 1) // <= 3

    // cc.addRequiredSkillByName("攻擊", 3)    // <= 7

    // cc.addRequiredSkillByName("昏厥耐性", 3) // <= 3
    // cc.addRequiredSkillByName("迴避距離UP", 2) // <= 3

    // cc.addRequiredSkillByName("挑戰者", 1)  // <= 5
    // cc.addRequiredSkillByName("廣域化", 2) // <= 5
    
    // cc.addRequiredSkillByName("利刃／彈藥節約", 1)  // <= 1
    // cc.addRequiredSkillByName("精靈加護", 3) // <= 3
    // cc.addRequiredSkillByName("無傷", 1) // <= 3
    // cc.addRequiredSkillByName("熱傷害無效", 2) // <= 3
    

    proc := prepareProcedure(*cc)
    proc.execute()
}


// 黑暗左右:扩散
func testExecuteProcedureHeavyBowgun4(t *testing.T) {
    cc := newConstraintCollection(loadDataManager())
    addCommonDecorationLimitations(cc)

    cc.addRequiredSkillByName("屬性解放／裝填擴充", 3)    // <= 3
    cc.addRequiredSkillByName("利刃／彈藥節約", 1)  // <= 1
    cc.addRequiredSkillByName("體力增強", 3)  // <= 3

    cc.addRequiredSkillByName("攻擊", 4)    // <= 7
    cc.addRequiredSkillByName("無傷", 1) // <= 3

    cc.addRequiredSkillByName("耳塞", 5) // <= 5
    // cc.addRequiredSkillByName("減輕膽怯", 2) // <= 3
    

    proc := prepareProcedure(*cc)
    proc.execute()
}


// 熔山重弩:扩散 极贝
// 扩散弹1 2 3，徹甲榴弹1 2 3，睡眠弹2,（数字代表弹种等级）龙击弹，火药粉2 3
func testExecuteProcedureHeavyBowgun5(t *testing.T) {
    cc := newConstraintCollection(loadDataManager())
    addCommonDecorationLimitations(cc)
        
    cc.addExtraSlot(2, 1)
    cc.addRequiredSkillByName("屬性解放／裝填擴充", 3)    // <= 3
    cc.addRequiredSkillByName("利刃／彈藥節約", 1)  // <= 1
    cc.addRequiredSkillByName("體力增強", 3)  // <= 3

    cc.addRequiredSkillByName("防禦強化", 1)    // <= 1

    cc.addRequiredSkillByName("收刀術", 3) // <= 3
    cc.addRequiredSkillByName("減輕膽怯", 2) // <= 3

    cc.addRequiredSkillByName("無傷", 3) // <= 3
    cc.addRequiredSkillByName("體力回復量UP", 3) // <= 3
    cc.addRequiredSkillByName("廣域化", 2) // <= 5

    // cc.addRequiredSkillByName("攻擊", 4) // <= 7
    // cc.addRequiredSkillByName("無傷", 3) // <= 3

    proc := prepareProcedure(*cc)
    proc.execute()
}


// 铠罗突击弩・贼 
// 散弹，极贝
func testExecuteProcedureHeavyBowgun6(t *testing.T) {
    cc := newConstraintCollection(loadDataManager())
    addCommonDecorationLimitations(cc)
        
    // cc.addExtraSlot(2, 1)
    cc.addRequiredSkillByName("屬性解放／裝填擴充", 3) // <= 3
    cc.addRequiredSkillByName("散彈・剛射強化", 1) // <= 1
    // cc.addRequiredSkillByName("心眼／彈道強化", 1)    // <= 1
    cc.addRequiredSkillByName("體力增強", 3)  // <= 3
    // cc.addRequiredSkillByName("無屬性強化", 1) // <= 1
    cc.addRequiredSkillByName("耳塞", 5) // <= 5

    cc.addRequiredSkillByName("防禦性能", 5)    // <= 5
    cc.addRequiredSkillByName("防禦強化", 1) // <= 1
    // cc.addRequiredSkillByName("弱點特效", 3) // <= 3
    // cc.addRequiredSkillByName("超會心", 3) // <= 3
    // cc.addRequiredSkillByName("攻擊", 3)    // <= 7

    cc.addRequiredSkillByName("昏厥耐性", 3) // <= 3

    // cc.addRequiredSkillByName("毅力", 1) // <= 1
    
    cc.addRequiredSkillByName("收刀術", 1) // <= 3
    // cc.addRequiredSkillByName("體力回復量UP", 2)  // <= 3
    // cc.addRequiredSkillByName("廣域化", 2)  // <= 3

    // cc.addRequiredSkillByName("攻擊", 3)    // <= 7
    // cc.addRequiredSkillByName("回復速度", 3)  // <= 3
    // cc.addRequiredSkillByName("迴避距離UP", 2) // <= 3
    
    // cc.addRequiredSkillByName("攻擊", 5)    // <= 7
    // cc.addRequiredSkillByName("精靈加護", 3) // <= 3
    // cc.addRequiredSkillByName("無傷", 1) // <= 3
    // cc.addRequiredSkillByName("減輕膽怯", 2) // <= 3
    

    proc := prepareProcedure(*cc)
    proc.execute()
}



// 炎妃 冥灯
// 斩裂，贝
func testExecuteProcedureHeavyBowgun7(t *testing.T) {
    cc := newConstraintCollection(loadDataManager())
    addCommonDecorationLimitations(cc)
        
    cc.addExtraSlot(3, 1)
    cc.addExtraSlot(1, 1)

    cc.addRequiredSkillByName("屬性解放／裝填擴充", 3) // <= 3
    // cc.addRequiredSkillByName("散彈・剛射強化", 1) // <= 1
    // cc.addRequiredSkillByName("心眼／彈道強化", 1)    // <= 1
    cc.addRequiredSkillByName("體力增強", 3)  // <= 3
    // cc.addRequiredSkillByName("無屬性強化", 1) // <= 1
    cc.addRequiredSkillByName("耳塞", 5) // <= 5

    cc.addRequiredSkillByName("防禦性能", 5)    // <= 5
    cc.addRequiredSkillByName("防禦強化", 1) // <= 1
    // cc.addRequiredSkillByName("弱點特效", 3) // <= 3
    // cc.addRequiredSkillByName("超會心", 3) // <= 3
    // cc.addRequiredSkillByName("攻擊", 3)    // <= 7

    cc.addRequiredSkillByName("昏厥耐性", 3) // <= 3

    cc.addRequiredSkillByName("毅力", 1) // <= 1
    
    cc.addRequiredSkillByName("收刀術", 1) // <= 3
    cc.addRequiredSkillByName("整備", 2) // <= 3

    // cc.addRequiredSkillByName("體力回復量UP", 2)  // <= 3
    // cc.addRequiredSkillByName("廣域化", 2)  // <= 3

    // cc.addRequiredSkillByName("攻擊", 3)    // <= 7
    // cc.addRequiredSkillByName("回復速度", 3)  // <= 3
    // cc.addRequiredSkillByName("迴避距離UP", 2) // <= 3
    
    // cc.addRequiredSkillByName("攻擊", 5)    // <= 7
    // cc.addRequiredSkillByName("精靈加護", 3) // <= 3
    // cc.addRequiredSkillByName("無傷", 1) // <= 3
    // cc.addRequiredSkillByName("減輕膽怯", 2) // <= 3
    

    proc := prepareProcedure(*cc)
    proc.execute()
}
