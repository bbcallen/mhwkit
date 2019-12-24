package mhw

import (
    "testing"
)

// 轻弩水速射vs冥赤龙
func testExecuteProcedureLightBowgun_4_WaterAmmo_VsSafiJiiva(t *testing.T) {
    cc := newConstraintCollection(loadDataManager())
    addCommonDecorationLimitations(cc)
        
    cc.addExtraSlot(4, 1)

    // cc.addExtraSlot(4, -1) // 达人II
    // cc.addRequiredSkillByName("看破", -2)
    cc.addExtraSlot(4, -1)
    cc.addRequiredSkillByName("體力增強", -1)
    cc.addRequiredSkillByName("破壞王", -1)
    cc.addExtraSlot(4, -1)
    cc.addRequiredSkillByName("水屬性攻擊強化", -2)
    cc.addExtraSlot(4, -1)
    cc.addRequiredSkillByName("屬性解放／裝填擴充", -1) // <= 3
    cc.addRequiredSkillByName("精靈加護", -1)  // <= 3
    cc.addExtraSlot(4, -1)
    cc.addRequiredSkillByName("破壞王", -1) // <= 3
    cc.addRequiredSkillByName("整備", -1)  // <= 3
    // cc.addExtraSlot(4, -1)
    // cc.addRequiredSkillByName("屬性異常狀態耐性", -1)
    // cc.addRequiredSkillByName("攻擊", -1)
    // cc.addExtraSlot(4, -1)
    // cc.addRequiredSkillByName("挑戰者", -1) // <= 3
    // cc.addRequiredSkillByName("體力增強", -1)  // <= 3


    cc.addRequiredSkillByName("真・龍脈覺醒", 1)
    cc.addRequiredSkillByName("龍脈覺醒", 1)
    // cc.addRequiredSkillByName("迴避性能", 4) // <= 5
    cc.addRequiredSkillByName("屬性異常狀態耐性", 3)
    cc.addRequiredSkillByName("超會心", 3)
    // cc.addRequiredSkillByName("精神抖擻", 1) // <= 3

    cc.addRequiredSkillByName("利刃／彈丸節約", 1)
    cc.addRequiredSkillByName("屬性解放／裝填擴充", 3)
    cc.addRequiredSkillByName("體力增強", 3)
    cc.addRequiredSkillByName("破壞王", 3)
    cc.addRequiredSkillByName("水屬性攻擊強化", 6)

    cc.addRequiredSkillByName("怨恨", 1)

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


func testExecuteProcedureLightBowgun_4_StickyAmmo_VsSafiJiiva(t *testing.T) {
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

