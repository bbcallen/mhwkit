package mhw

import (
	"errors"
	"fmt"
)

func (dm *dataManager) loadDecorations() {
	// https://www.mhchinese.wiki/adornments
	const rawData = `
0 0 无 无

1 5 耐毒珠【1】 毒耐性
1 5 耐麻珠【1】 麻痺耐性
1 5 耐眠珠【1】 睡眠耐性
1 6 耐絕珠【1】 昏厥耐性
1 5 耐爆珠【1】 爆破異常狀態的耐性
1 5 耐裂珠【1】 裂傷耐性
1 5 耐防珠【1】 防御力DOWN耐性
3 6 防音珠【3】 耳塞
2 6 防風珠【2】 風壓耐性
2 6 耐震珠【2】 耐震

1 7 攻擊珠【1】 攻擊
1 6 防禦珠【1】 防禦
1 6 體力珠【1】 體力增強
1 6 早復珠【1】 回復速度
1 5 耐火珠【1】 火耐性
1 5 耐水珠【1】 水耐性
1 5 耐冰珠【1】 冰耐性
1 5 耐雷珠【1】 雷耐性
1 5 耐龍珠【1】 龍耐性
1 6 耐屬珠【1】 屬性異常狀態耐性

1 6 火炎珠【1】 火屬性攻擊強化
1 6 流水珠【1】 水屬性攻擊強化
1 6 冰結珠【1】 冰屬性攻擊強化
1 6 雷光珠【1】 雷屬性攻擊強化
1 6 破龍珠【1】 龍屬性攻擊強化
1 6 毒珠【1】 毒屬性強化
1 7 麻痺珠【1】 麻痺屬性強化
1 7 睡眠珠【1】 睡眠屬性強化
1 7 爆破珠【1】 爆破屬性強化
3 6 毒瓶珠【3】 毒瓶追加

3 7 痺瓶珠【3】 麻痺瓶追加
3 7 眠瓶珠【3】 睡眠瓶追加
3 7 爆瓶珠【3】 爆破瓶追加
3 8 解放珠【3】 屬性解放／裝填擴充
1 7 達人珠【1】 看破
2 8 超心珠【2】 超會心
2 6 痛擊珠【2】 弱點特效
2 8 縮短珠【2】 集中
3 8 匠珠【3】 匠
2 8 拔刀珠【2】 拔刀術【技】

2 6 重擊珠【2】 破壞王
2 6 KO珠【2】 KO術
1 6 奪氣珠【1】 奪取耐力
2 5 飛燕珠【2】 飛燕
2 7 全開珠【2】 力量解放
2 7 挑戰珠【2】 挑戰者
2 7 無傷珠【2】 無傷
2 7 底力珠【2】 火場怪力
1 5 逆境珠【1】 不屈
2 7 逆上珠【2】 怨恨

1 6 鼓笛珠【1】 吹笛名人
2 7 増彈珠【2】 砲彈裝填數UP
1 7 特射珠【1】 特殊射擊強化
1 8 砲術珠【1】 砲術
1 5 砲手珠【1】 砲擊手
2 8 強走珠【2】 跑者
2 6 體術珠【2】 體術
1 5 無食珠【1】 飢餓耐性
2 6 迴避珠【2】 迴避性能
2 6 跳躍珠【2】 迴避距離UP

1 7 鐵壁珠【1】 防禦性能
1 6 速納珠【1】 收刀術
1 6 友愛珠【1】 廣域化
1 6 持續珠【1】 道具使用強化
1 5 節食珠【1】 滿足感
1 6 快吃珠【1】 快吃
1 6 研磨珠【1】 砥石使用高速化
1 6 爆師珠【1】 炸彈客
1 8 茸好珠【1】 最愛菇類
1 5 加護珠【1】 精靈加護

1 5 指示珠【1】 指示隨從
1 5 植學珠【1】 植生學
1 5 地學珠【1】 地質學
1 6 投石珠【1】 投射器裝填數UP
2 7 渾身珠【2】 精神抖擻
1 5 潛伏珠【1】 潛伏
3 6 耐衝珠【3】 減輕膽怯
1 5 環境珠【1】 環境利用知識
1 5 沼渡珠【1】 適應水場・深雪
1 5 標本珠【1】 昆蟲標本達人

1 5 耐瘴珠【1】 瘴氣耐性
1 5 嗅覚珠【1】 追蹤達人
1 5 威嚇珠【1】 威嚇
2 6 滑走珠【2】 滑走強化
1 6 治癒珠【1】 體力回復量UP
3 7 強彈珠【3】 通常彈・通常箭強化
3 8 散彈珠【3】 散彈・剛射強化
2 7 昂揚珠【2】 強化持續
1 5 絕路珠【1】 死裡逃生
3 7 龍封珠【3】 龍封力強化

1 5 整備珠【1】 整備
2 8 強弓珠【2】 解放弓的蓄力階段
2 8 心眼珠【2】 心眼／彈道強化
2 7 強壁珠【2】 防禦強化
2 7 剛刃珠【2】 剛刃研磨
2 6 無擊珠【2】 無屬性強化
4 9 耐冰珠II【4】 冰耐性 2
4 12 達人珠II【4】 看破 2

##
2 7 早氣珠【2】 耐力急速回復
3 8 貫通珠【3】 貫通彈・龍之箭強化
#
`
	if len(dm.decorations) > 0 {
		panic(errors.New(fmt.Sprintf("duplicated decorations loading")))
	}

	scanner := newDataScannerFromString(rawData)
	for scanner.scanRow() {
		d := newDecoration()
		for scanner.scanField() {
			switch scanner.fieldIndex {
			case 0:
				d.size = scanner.intValue()
			case 1:
				d.rarity = scanner.intValue()
			case 2:
				d.name = scanner.textValue()
			case 3:
				d.skillId = dm.getSkillIdByName(scanner.textValue())
			case 4:
				d.enhancedLevel = scanner.intValue()
			default:
				fmt.Println("unexpected field near: ", scanner.line)
				panic(errors.New("unexpected field"))
			}
		}
		dm.registerDecoration(*d)
	}
}
