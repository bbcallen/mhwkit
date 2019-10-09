package mhw

import (
	"errors"
	"fmt"
)

func (dm *dataManager) loadSkills() {
	// https://www.mhchinese.wiki/skills
	const rawData = `
0 无
3 體力增強
3 體力回復量UP
5 體術
3 KO術

5 力量解放

6 火屬性攻擊強化
5 水屬性攻擊強化
1 不屈
3 火耐性
3 水耐性
5 火場怪力
1 心眼／彈道強化

1 加速再生

5 冰屬性攻擊強化
3 收刀術
3 回復速度
3 冰耐性
5 耳塞
1 死裡逃生
5 匠
3 地質學

7 攻擊
7 防禦
5 防禦性能
3 快吃
3 投射器裝填數UP
1 吹笛名人
3 防御力DOWN耐性
1 利刃／彈丸節約
1 防禦強化

3 昏厥耐性
3 泥耐性
3 昆蟲標本達人
1 拔刀術【力】
3 拔刀術【技】
1 幸運

3 毒屬性強化
7 看破
3 耐力急速回復
3 毒耐性
5 風壓耐性
1 毒瓶追加
3 威嚇
5 指示隨從
1 研究者
1 突破耐力上限
3 耐震
1 毒傷害強化
1 飛身躍入
1 飛燕
3 炸彈客
5 挑戰者
1 風壓完全無効
5 怨恨

3 砲術
5 迴避性能
3 迴避距離UP
3 飢餓耐性
2 特殊射擊強化
1 振奮
3 弱點特效
1 閃光強化
1 追蹤達人
3 砥石使用高速化
1 剝取鐵人
2 砲擊手
1 砲彈裝填數UP
1 捕獲名人
1 烤肉名人
3 破壞王
1 剛刃研磨
1 剝取名人

3 麻痺屬性強化
3 麻痺耐性
1 採集達人
1 異臭名人
1 探索者的幸運
1 釣魚名人
1 貫通彈・龍之箭強化
3 強化持續
1 通常彈・通常箭強化
1 強運
1 麻痺瓶追加

3 集中
3 跑者
3 最愛菇類
4 植生學
1 鈍器能手
3 減輕膽怯
3 裂傷耐性
3 超會心
1 無屬性強化
1 超回復力
3 無傷
1 散彈・剛射強化

5 雷屬性攻擊強化
3 道具使用強化
3 雷耐性
1 搬運達人
1 跳躍鐵人
1 會心攻擊【屬性】
1 會心攻擊【特殊】
1 滑走強化
1 蜂蜜獵人
1 解放弓的蓄力階段
1 達人藝

3 睡眠屬性強化
3 奪取耐力
3 精靈加護
3 睡眠耐性
3 環境利用知識
1 察覺
1 滿足感
1 睡眠瓶追加
3 精神抖擻

3 潛伏
3 適應水場・深雪
5 廣域化
1 適應瘴氣環境
1 熱傷害無效
1 毅力

5 龍屬性攻擊強化
3 龍耐性
1 導蟲反應距離UP
3 整備
1 龍封力強化
3 瘴氣耐性

1 騎乘名人

1 攀岩者
1 蹲下移動速度UP
3 爆破屬性強化
3 爆破異常狀態的耐性
1 爆破瓶追加

3 屬性解放／裝填擴充
3 屬性異常狀態耐性
1 飛燕【屬性】

3 攻擊守勢
1 會心擊【特殊】
1 炸彈客・極意
1 KO術・極意
1 真・利刃／彈丸節約
1 渾身・極意
1 挑戰者・極意
1 滿足感・極意
3 轉禍為福
1 會心擊【屬性】
1 寒氣鍊成
1 属性加速
1 真・属性加速
1 大地缝缠
1 真・大地缝缠
`
	if len(dm.skills) > 0 {
		panic(errors.New(fmt.Sprintf("duplicated skills loading")))
	}

	scanner := newDataScannerFromString(rawData)
	for scanner.scanRow() {
		// fmt.Printf("scanRow: %v\n", scanner.line)
		s := newSkill()
		for scanner.scanField() {
			// fmt.Printf("scanField: %v\n", scanner.fieldIndex)

			switch scanner.fieldIndex {
			case 0:
				s.maxLevel = scanner.intValue()
			case 1:
				s.name = scanner.textValue()
			default:
				fmt.Println("unexpected field near: ", scanner.line)
				panic(errors.New("unexpected field"))
			}
		}
		// fmt.Printf("skill: %v\n", s)
		dm.registerSkill(*s)
	}
}
