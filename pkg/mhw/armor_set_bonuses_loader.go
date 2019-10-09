package mhw

import (
	"errors"
	"fmt"
)

func (dm *dataManager) loadArmorSetBonuses() {
	const rawData = `
无 无 0
蠻顎龍之鬥志 振奮 2 突破耐力上限 4
櫻火龍之奧秘 毒傷害強化 3
熔山龍的奧秘 會心攻擊【特殊】 3
風漂龍的恩竉 幸運 2 解放弓的蓄力階段 4
慘爪龍的奧祕 拔刀術【力】 2 剛刃研磨 4
火龍奧祕 會心攻擊【屬性】 2 心眼／彈道強化 4
角龍之奧祕 鈍器能手 2 無屬性強化 4
爆鎚龍的守護 防禦強化 3
爆鱗龍的守護 毅力 3
滅盡龍之飢餓 加速再生 3
炎王龍之武技 達人藝 3
鋼龍之飛翔 風壓完全無効 3
屍套龍之命脈 超回復力 3
麒麟之恩寵 捕獲名人 3
冥燈龍的神秘 利刃／彈丸節約 3
公會的指引 強運 4
調查團的指引 剝取名人 4
炎妃龍的恩寵 突破耐力上限 2 心眼／彈道強化 4
星辰的祝福 強運 5
龍騎士之證 飛燕【屬性】 2 達人藝 4

蠻顎龍的霸气 突破耐力上限 2
雌火龍的真髓 毒傷害強化 2
太古的神秘 會心擊【特殊】 2
調查團的煉金術 炸彈客・極意 3
火龍的真髓 心眼／彈道強化 3
角龍的霸氣 KO術・極意 3
風漂龍的霸氣 解放弓的蓄力階段 3
慘爪龍的真髓 剛刃研磨 3
爆鎚龍的霸氣 防禦強化 2
迅龍的真髓 真・利刃／彈丸節約 2
斬龍的真髓 渾身・極意 3
碎龍的真髓 挑戰者・極意 3
轟龍的真髓 滿足感・極意 3
教官的引導 捕獲名人 2
冰呪龍的神秘 會心擊【屬性】 2 寒氣鍊成 4
屍套龍的靈脈 超回復力 2
幻獸的神秘 強運 3
溟波龍的神秘 属性加速 2 真・属性加速 4
`

	if len(dm.armorSetBonuses) > 0 {
		panic(errors.New(fmt.Sprintf("duplicated armor set bonuses loading")))
	}

	scanner := newDataScannerFromString(rawData)
	for scanner.scanRow() {
		b := newArmorSetBonus()
		skillId := 0
		for scanner.scanField() {
			field := scanner.fieldIndex
			switch field {
			case 0:
				b.name = scanner.textValue()
			case 1, 3:
				skillId = dm.getSkillIdByName(scanner.textValue())
			case 2, 4:
				count := scanner.intValue()
				b.activatedSkillIdsByComponentCount[count] = skillId
			default:
				fmt.Println("unexpected column near: ", scanner.line)
				panic(errors.New("unexpected column"))
			}
		}
		// fmt.Printf("armor set bonus: %v\n", b)
		dm.registerArmorSet(*b)
	}
}
