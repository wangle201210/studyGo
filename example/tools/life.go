package main

import (
	"fmt"
	"github.com/wangle201210/studyGo/tools/life"
)

func housingFund() {
	// 一个人的公积金
	data := []int64{
		100000000000,
		// 1700,1700,1700,1700,1700,1700,1700,1700,1700,1700,
		// 1700,1700,1700,1700,
	}
	// 4.3
	// 贷款总额： 1000000
	// 每月需要还：4948.71
	// 减去公积金后还需要 3176
	// 总共需要还 1781537.19
	// 其中利息有 781537.19
	// 利息比例为 0.43869
	// 5.28
	// 贷款总额： 1000000
	// 每月需要还：5540.63
	// 减去公积金后还需要 3768
	// 总共需要还 1994627.88
	// 其中利息有 994627.88
	// 利息比例为 0.49865

	// 第二份公积金
	data2 := []int64{
		// 1800, 1800, 1800, 1800, 1800, 1800,
		// 1800, 1800, 1800, 1800, 1800, 1800,
		// 1800, 1800, 1800,
	}
	// 215w 给 65w 首付
	// 每月需要还：3308.46
	// 每月需要还：3416.13

	// 每月需要还：3911.00
	// 每月需要还：3803.33
	total := int64(90 * 10000)
	month := 30 * 12
	fund := life.NewHosingFund(
		data,
		life.WithData2(data2),
		life.WithLoan(total),             // 总贷款150w
		life.WithRateBusiness(4.3),       // 商贷利率
		life.WithMaxHosingFund(90*10000), // 公积金最高额度
		// life.WithRateMaxBusiness(5.28),   // 曾经的最高利率，用来计算当前便宜了多少
		life.WithMonth(int64(month)),
	)
	// println(fund.HouseNew())
	// return
	res := fund.InterestMonth()
	realNeed := int64(res) - data[len(data)-1]
	if len(data2) > 0 {
		realNeed -= data2[len(data2)-1]
	}
	needAll := res * float64(month)
	fmt.Printf("贷款总额： %d"+
		"\n每月需要还：%.2f"+
		"\n减去公积金后还需要 %d"+
		"\n总共需要还 %.2f"+
		"\n其中利息有 %.2f"+
		"\n利息比例为 %.5f",
		total, res, realNeed, needAll, needAll-float64(total), 1-float64(total)/needAll)

	// fmt.Printf("贷款总额： %d"+
	//	"\n每月需要还本金：%.2f"+
	//	"\n总共需要还 %.2f"+
	//	"\n其中利息有 %.2f"+
	//	"\n利息比例为 %.5f",
	//	total, total/12/30, 3000, needAll, needAll-float64(total), 1-float64(total)/needAll)
	// month := fund.CapitalMonth()
	// for i, i2 := range month {
	//	fmt.Printf("第 %3d 月还 %.2f \n", i+1, i2)
	// }
	// houseNew := fund.HouseNew()
	// houseOld := fund.HouseOld()
	// moreUse, str := fund.MoreUse()
	// fmt.Printf("以前能贷公积金(%d)\n"+
	//	"现在能贷公积金(%d)\n"+
	//	"比以前少贷了(%d)\n"+
	//	"额外利息增加了(%d)\n"+
	//	"%s",
	//	houseOld, houseNew, houseOld-houseNew, moreUse, str)
}
