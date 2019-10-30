package mhw

func addCommonDecorationLimitations(cc *constraintCollection) {
    // 2
    cc.addDecorationLimitationByName("攻擊珠【1】", 5)   // 攻擊 7
    // 4
    cc.addDecorationLimitationByName("縮短珠【2】", 2)   // 集中 3
    cc.addDecorationLimitationByName("匠珠【3】", 2)    // 匠 5
    cc.addDecorationLimitationByName("拔刀珠【2】", 1)   // 拔刀術【技】 3
    // 5
    cc.addDecorationLimitationByName("逆上珠【2】", 4)   // 怨恨 5
    // 7
    cc.addDecorationLimitationByName("鐵壁珠【1】", 1)   // 防禦性能 5
    cc.addDecorationLimitationByName("茸好珠【1】", 2)   // 最愛菇類 3

    // *
    cc.addDecorationLimitationByName("早氣珠【2】", 0)   // 耐力急速回復 3
    cc.addDecorationLimitationByName("貫通珠【3】", 0)   // 貫通彈・龍之箭強化 1
}

func getArmorBlackList() map[string]bool {
    blackList := make(map[string]bool)

    return blackList
}