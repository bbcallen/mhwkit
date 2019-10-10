package mhw

import (
	"testing"
)

func TestExecuteProcedure(t *testing.T) {
	testExecuteProcedureGreatSword_4_VsLunastra(t)
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
