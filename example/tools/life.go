package main

import (
	"fmt"
	"github.com/wangle201210/studyGo/tools/life"
)

func housingFund() {
	// 一个人的公积金
	data := []int64{
		320, 320, 320, 320, 320, 320, 320, 320, 320, 320, 320, 320, 320, 320, 320, 320, 320, 320, 320, 320, 320, 320, 320, 320, 320, // hj
		300,                                                                                            //yf
		1560, 1560, 1560, 1560, 1560, 1560, 1560, 1560, 1560, 1560, 1560, 1560, 1560, 1560, 1560, 1560, // yl 2020-12 - 2022-3
		1560, 1560, 1560, 1560, 1560,
	}
	// 第二份公积金
	data2 := []int64{}
	total := int64(100 * 10000)
	fund := life.NewHosingFund(
		data,
		life.WithData2(data2),
		life.WithLoan(total),             // 总贷款150w
		life.WithRateBusiness(4.45),      // 商贷利率
		life.WithMaxHosingFund(00*10000), // 公积金最高额度
		life.WithRateMaxBusiness(6.25),   // 曾经的最高利率，用来计算当前便宜了多少
	)
	//houseNew := fund.HouseNew()
	//println(houseNew)
	//每月需要还：1740.83
	//减去公积金后还需要 180
	//总共需要还 626697.10
	//其中利息有 226697.10
	//利息比例为 0.36173
	//
	//now, old := fund.InterestMaxReduceMonth()
	//fmt.Printf("贷款总额%d"+
	//	"\n现在每月需要%.2f"+
	//	"\n最高时每月需要%.2f"+
	//	"\n每月少还%.2f"+
	//	"\n总共少还%.2f",
	//	total, now, old, old-now, (old-now)*12*30)
	//return
	res := fund.InterestMonth()
	realNeed := int64(res) - data[len(data)-1]
	if len(data2) > 0 {
		realNeed -= data2[len(data2)-1]
	}
	needAll := res * 30 * 12
	fmt.Printf("贷款总额： %d"+
		"\n每月需要还：%.2f"+
		"\n减去公积金后还需要 %d"+
		"\n总共需要还 %.2f"+
		"\n其中利息有 %.2f"+
		"\n利息比例为 %.5f",
		total, res, realNeed, needAll, needAll-float64(total), 1-float64(total)/needAll)

	fmt.Printf("贷款总额： %d"+
		"\n每月需要还本金：%.2f"+
		"\n总共需要还 %.2f"+
		"\n其中利息有 %.2f"+
		"\n利息比例为 %.5f",
		total, total/12/30, 3000, needAll, needAll-float64(total), 1-float64(total)/needAll)
	//month := fund.CapitalMonth()
	//for i, i2 := range month {
	//	fmt.Printf("第 %3d 月还 %.2f \n", i+1, i2)
	//}
	//houseNew := fund.HouseNew()
	//houseOld := fund.HouseOld()
	//moreUse, str := fund.MoreUse()
	//fmt.Printf("以前能贷公积金(%d)\n"+
	//	"现在能贷公积金(%d)\n"+
	//	"比以前少贷了(%d)\n"+
	//	"额外利息增加了(%d)\n"+
	//	"%s",
	//	houseOld, houseNew, houseOld-houseNew, moreUse, str)
}
