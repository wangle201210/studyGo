package main

import (
	"fmt"
	"github.com/wangle201210/studyGo/tools/life"
)

func housingFund() {
	data := []int64{
		320, 320, 320, 320, 320, 320, 320, 320, 320, 320, 320, 320, 320, 320, 320, 320, 320, 320, 320, 320, 320, 320, 320, 320, 320,
		300,
		1560, 1560, 1560, 1560, 1560, 1560, 1560, 1560, 1560, 1560, 1560, // 11个月
		//1560, 1560, 1560, //1560, 1560,
	}

	fund := life.NewHosingFund(data, life.WithPrintNum(1))
	houseNew := fund.HouseNew()
	houseOld := fund.HouseOld()
	moreUse, str := fund.MoreUse()
	fmt.Printf("以前能贷公积金(%d)\n"+
		"现在能贷公积金(%d)\n"+
		"比以前少贷了(%d)\n"+
		"额外利息增加了(%d)\n"+
		"%s",
		houseOld, houseNew, houseOld-houseNew, moreUse, str)
}

func main() {
	housingFund()
}
