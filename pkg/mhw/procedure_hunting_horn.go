package mhw

import (
    "testing"
)

// 泛用：防大，攻大，岩贼龙号角2
// 攻坚：防大，气绝无效，风压无效，衣装时间延长
// 防御：金刚身，耐震，风压完全无效，高级耳栓
// 高周波
//
// 常用：防大，攻大，耐力消耗减少
// 全异常状态无效

// 狩猎笛
func testExecuteProcedureHuntingHorn1(t *testing.T) {
    cc := newConstraintCollection(loadDataManager())
    addCommonDecorationLimitations(cc)
        
    cc.addExtraSlot(1, 1)
    cc.addRequiredSkillByName("體力增強", 3)  // <= 3
    cc.addRequiredSkillByName("廣域化", 5) // <= 5
    cc.addRequiredSkillByName("快吃", 3)    // <= 3
    cc.addRequiredSkillByName("滿足感", 1)    // <= 1

    cc.addRequiredSkillByName("最愛菇類", 1)   // <= 3
    cc.addRequiredSkillByName("收刀術", 3) // <= 3
    cc.addRequiredSkillByName("耳塞", 5) // <= 5
    cc.addRequiredSkillByName("吹笛名人", 1) // <= 1

    proc := prepareProcedure(*cc)
    proc.execute()
}

// 恐暴 日常
func testExecuteProcedureHuntingHorn2(t *testing.T) {
    cc := newConstraintCollection(loadDataManager())
    addCommonDecorationLimitations(cc)

    cc.addRequiredSkillByName("體力增強", 3)  // <= 3
    cc.addRequiredSkillByName("吹笛名人", 1) // <= 1

    cc.addRequiredSkillByName("攻擊", 4)    // <= 7
    cc.addRequiredSkillByName("弱點特效", 3) // <= 3
    cc.addRequiredSkillByName("匠", 3)    // <= 5
    cc.addRequiredSkillByName("看破", 7)   // <= 7
    cc.addRequiredSkillByName("超會心", 3) // <= 3

    proc := prepareProcedure(*cc)
    proc.execute()
}

