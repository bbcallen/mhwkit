package mhw

import (
	"errors"
	"fmt"
)

func (dm *dataManager) loadCharms() {
	// https://www.mhchinese.wiki/charms
	const rawData = `
0 0 无 无

1 1 任意護石 任意護石

3 3 耐毒護石 毒耐性
3 3 攻擊護石 攻擊
3 3 防禦護石 防禦
3 3 耐痺護石 麻痺耐性
3 3 耐眠護石 睡眠耐性
3 3 耐絕護石 昏厥耐性
3 3 防風護石 風壓耐性
3 3 體力護石 體力增強
3 3 治癒護石 體力回復量UP
3 3 耐火護石 火耐性
3 3 耐水護石 水耐性
3 3 耐雷護石 雷耐性
3 3 火焰護石 火屬性攻擊強化
3 3 流水護石 水屬性攻擊強化
3 3 雷光護石 雷屬性攻擊強化
3 3 毒擊護石 毒屬性強化
3 3 痺擊護石 麻痺屬性強化
3 3 睡擊護石 睡眠屬性強化
3 3 KO護石 KO術
3 3 奪氣護石 奪取耐力
3 3 砲術護石 砲術
3 3 斷食護石 飢餓耐性
3 3 鐵壁護石 防禦性能
3 3 友愛護石 廣域化
3 3 持續護石 道具使用強化
3 3 快吃護石 快吃
3 3 加護護石 精靈加護
3 3 指揮護石 指示隨從
3 3 植學護石 植生學
3 3 地學護石 地質學
3 3 投石護石 投射器裝填數UP
3 3 潛伏護石 潛伏
3 3 環境護石 環境利用知識
3 3 沼渡護石 適應水場・深雪
3 3 昆蟲學護石 昆蟲標本達人
3 3 威嚇護石 威嚇
2 3 砲擊手護石 砲擊手
1 3 小胃口護石 滿足感
1 3 追踨護石 追蹤達人
1 3 不屈護石 不屈
1 3 吹笛護石 吹笛名人
1 3 裝填護石 砲彈裝填數UP

3 4 耐裂護石 裂傷耐性
3 4 早復護石 回復速度
3 4 耐冰護石 冰耐性
3 4 耐屬護石 屬性異常狀態耐性
3 4 冰結護石 冰屬性攻擊強化
3 4 達人護石 看破
3 4 迴避護石 迴避性能
3 4 納刀護石 收刀術
3 4 研磨護石 砥石使用高速化

3 5 耐爆護石 爆破異常狀態的耐性
3 5 耐龍護石 龍耐性
3 5 破龍護石 龍屬性攻擊強化
3 5 爆破護石 爆破屬性強化
2 5 拔刀護石 拔刀術【技】
2 5 特射護石 特殊射擊強化
2 5 強走護石 跑者
3 5 體術護石 體術
2 5 早氣護石 耐力急速回復
2 5 跳躍護石 迴避距離UP
3 5 爆師護石 炸彈客
2 5 嗜菇護石 最愛菇類
2 5 覺醒護石 屬性解放／裝填擴充
3 5 匠之護石 匠
1 5 泥浴護石 泥耐性 飛身躍入
1 5 陷阱彈護石 異臭名人 閃光強化
1 5 狩獵生活護石 烤肉名人 釣魚名人
1 5 籌備護石 搬運達人 蜂蜜獵人
1 5 採集鐵人護石 採集達人 剝取鐵人

3 6 耐防護石 防御力DOWN耐性
3 6 耳塞護石 耳塞
2 6 耐震護石 耐震
2 6 痛擊護石 弱點特效
2 6 集中護石 集中
2 6 重擊護石 破壞王
2 6 暴怒護石 怨恨
2 6 底力護石 火場怪力
3 6 裝備護石 整備

2 7 全開護石 力量解放
2 7 挑戰護石 挑戰者
2 7 無傷護石 無傷
2 7 渾身護石 精神抖擻
3 7 耐衝護石 減輕膽怯
3 7 耐瘴護石 瘴氣耐性
1 7 通常彈護石 通常彈・通常箭強化
1 7 心靜自然涼護石 熱傷害無效 適應瘴氣環境
1 7 騎手護石 騎乘名人 跳躍鐵人
1 7 調查達人護石 導蟲反應距離UP 研究者

1 8 超心護石 超會心
2 8 昂揚護石 強化持續
1 8 貫通彈護石 貫通彈・龍之箭強化
1 8 散彈護石 散彈・剛射強化
1 8 毒瓶護石 毒瓶追加
1 8 麻痺瓶護石 麻痺瓶追加
1 8 睡眠瓶謢石 睡眠瓶追加
1 8 爆破瓶護石 爆破瓶追加
1 8 風水護石 探索者之幸運 察覺
1 8 疾風護石 滑走強化 飛燕
1 8 滅龍護石 龍封力強化 龍屬性攻擊強化
1 8 堅守護石 防禦強化 死裡逃生

#特典
1 8 追風護石 攻擊 精靈加護

#Master
1 11 剛力護石 無屬性強化 鈍器能手
1 11 利刃護石 利刃／彈丸節約
1 11 封印護石 龍封力強化 振奮
1 11 毅力護石 毅力
`
	if len(dm.charms) > 0 {
		panic(errors.New(fmt.Sprintf("duplicated charms loading")))
	}

	scanner := newDataScannerFromString(rawData)
	for scanner.scanRow() {
		c := newCharm()
		for scanner.scanField() {
			field := scanner.fieldIndex
			switch field {
			case 0:
				c.maxLevel = scanner.intValue()
			case 1:
				c.rarity = scanner.intValue()
			case 2:
				c.name = scanner.textValue()
			case 3, 4:
				c.enhancedSkillIds[field-3] = dm.getSkillIdByName(scanner.textValue())
			default:
				fmt.Println("unexpected field near: ", scanner.line)
				panic(errors.New("unexpected field"))
			}
		}
		dm.registerCharm(*c)
	}
}
