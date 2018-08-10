package mhw

import (
	"testing"
)

func addCommonDecorationLimitations(cc *constraintCollection) {
	cc.addDecorationLimitationBySkillName("無屬性強化", 3)
	cc.addDecorationLimitationBySkillName("攻擊", 2)
	cc.addDecorationLimitationBySkillName("體力增強", 7)

	cc.addDecorationLimitationBySkillName("集中", 0)
	cc.addDecorationLimitationBySkillName("匠", 1)
	cc.addDecorationLimitationBySkillName("弱點特效", 4)
	cc.addDecorationLimitationBySkillName("看破", 1)
	cc.addDecorationLimitationBySkillName("耳塞", 6)

	cc.addDecorationLimitationBySkillName("收刀術", 1)
	cc.addDecorationLimitationBySkillName("最愛菇類", 1)
	cc.addDecorationLimitationBySkillName("KO術", 3)
	cc.addDecorationLimitationBySkillName("破壞王", 3)
	
	
	cc.addDecorationLimitationByName("昂揚珠【2】", 1)
	cc.addDecorationLimitationByName("麻痺珠【1】", 1)
	cc.addDecorationLimitationByName("超心珠【2】", 1)
	cc.addDecorationLimitationByName("解放珠【3】", 0)
	cc.addDecorationLimitationByName("早氣珠【2】", 0)
}

// 铠罗麻
func testExecuteProcedureInsectGlaive1(t *testing.T) {
	cc := newConstraintCollection(loadDataManager())
	addCommonDecorationLimitations(cc)
		
	cc.addExtraSlot(3, 1)
    cc.addRequiredSkillByName("屬性解放／裝填擴充", 3)    // <= 3
	cc.addRequiredSkillByName("體力增強", 3)  // <= 3

	cc.addRequiredSkillByName("攻擊", 5)    // <= 7
	cc.addRequiredSkillByName("麻痺屬性強化", 3) // <= 3
	cc.addRequiredSkillByName("飛燕【屬性】", 1) // <= 1
	// cc.addRequiredSkillByName("達人藝", 1) // <= 1
	// cc.addRequiredSkillByName("會心攻擊【屬性】", 1) // <= 1
	
	// cc.addRequiredSkillByName("跳躍鐵人", 1) // <= 1
	// cc.addRequiredSkillByName("騎乘名人", 1) // <= 1
	cc.addRequiredSkillByName("看破", 5) // <= 7
	cc.addRequiredSkillByName("超會心", 1) // <= 3

	cc.addRequiredSkillByName("飛燕", 1) 	// <= 1	
	cc.addRequiredSkillByName("強化持續", 2) // <= 3
	cc.addRequiredSkillByName("匠", 1) // <= 1
	// cc.addRequiredSkillByName("挑戰者", 1) // <= 5
	// cc.addRequiredSkillByName("耐力急速回復", 1) // <= 3
	 


	proc := prepareProcedure(*cc)
	proc.execute()
}

func TestExecuteProcedure(t *testing.T) {
	testExecuteProcedureGreatSword4(t)
}

func BenchmarkExecuteProcedure(b *testing.B) {
	// 龍熱機關式【鋼翼】 改
	cc := newConstraintCollection(loadDataManager())
	cc.addExtraSlot(1, 2)
	// cc.addExtraSlot(1, -3) // 保留耐性

	addCommonDecorationLimitations(cc)

	cc.addRequiredSkillByName("無屬性強化", 1) // <= 1
	cc.addRequiredSkillByName("攻擊", 4)    // <= 7
	cc.addRequiredSkillByName("體力增強", 3)  // <= 3

	cc.addRequiredSkillByName("集中", 3)   // <= 3
	cc.addRequiredSkillByName("匠", 2)    // <= 5
	cc.addRequiredSkillByName("弱點特效", 3) // <= 3
	cc.addRequiredSkillByName("看破", 3)   // <= 7
	// cc.addRequiredSkillByName("耳塞", 5) // <= 5

	cc.addRequiredSkillByName("收刀術", 1) // <= 3
	cc.addRequiredSkillByName("KO術", 1) // <= 3
	cc.addRequiredSkillByName("破壞王", 1) // <= 3

	// cc.addRequiredSkillByName("龍屬性攻擊強化", 1) // <= 5

	// cc.addRequiredSkillByName("火耐性", 3) // <= 3

	b.ResetTimer()
	proc := prepareProcedure(*cc)
	proc.execute()
}

func TestDummy(t *testing.T) {
	// do nothing
}
