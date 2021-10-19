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
		//1560, 1560, 1560, 1560, 1560, 1560,
	}
	houseNew := life.HouseNew(data)
	houseOld := life.HouseOld(data)
	fmt.Printf("以前能贷公积金(%d)\n现在能贷公积金(%d)\n比以前少了(%d)", houseOld, houseNew, houseOld-houseNew)
}

func main() {
	housingFund()
}
