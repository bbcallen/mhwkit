package mhw

func addCommonDecorationLimitations(cc *constraintCollection) {
    // 4
    cc.addDecorationLimitationByName("攻擊珠【1】", 6)   // 攻擊 7
    // 7
    cc.addDecorationLimitationByName("匠珠【3】", 2)    // 匠 5
    // 9
    cc.addDecorationLimitationByName("強走珠【2】", 2)   // 跑者 3
    // 9
    cc.addDecorationLimitationByName("早氣珠【2】", 1)   // 耐力急速回復 3
    // 10
    cc.addDecorationLimitationByName("鐵壁珠【1】", 1)   // 防禦性能 5
    // 11
    cc.addDecorationLimitationByName("茸好珠【1】", 2)   // 最愛菇類 3

    // *
    cc.addDecorationLimitationByName("貫通珠【3】", 0)   // 貫通彈・龍之箭強化 1
}

func getArmorBlackList() map[string]bool {
    blackList := make(map[string]bool)
    return blackList
}
