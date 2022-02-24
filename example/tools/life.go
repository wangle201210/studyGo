package main

import (
	"fmt"
	"github.com/wangle201210/studyGo/tools/life"
)

func housingFund() {
	// 一个人的公积金
	data := []int64{
		320, 320, 320, 320, 320, 320, 320, 320, 320, 320, 320, 320, 320, 320, 320, 320, 320, 320, 320, 320, 320, 320, 320, 320, 320, // hj
		300,                                                                                      //yf
		1560, 1560, 1560, 1560, 1560, 1560, 1560, 1560, 1560, 1560, 1560, 1560, 1560, 1560, 1560, // yl 2020-12 - 2022-2
		1560, 1560, 1560, 1560, 1560, 1560,
	}
	// 第二份公积金
	//data2 := []int64{
	//	1574, 1574, 1574, 1574, 1574, 1574, 1574,
	//	1574, 1574, 1574, 1574, 1574, 1574,
	//	//1574, 1574, 1574, 1574, 1574, 1574,
	//	//1574,
	//}
	fund := life.NewHosingFund(
		data,
		//life.WithData2(data2),
		life.WithLoan(150*10000),         // 贷款120w
		life.WithRateBusiness(5.7),       // 利率5.7
		life.WithMaxHosingFund(40*10000), // 最高70w
	)
	res := fund.InterestMonth()
	realNeed := int64(res) - data[len(data)-1]
	fmt.Printf("每月需要还：%.2f\n减去公积金后还需要 %d", res, realNeed)
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

//func wechatDo() {
//	wechat.Do()
//}
