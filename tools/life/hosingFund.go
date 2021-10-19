package life

// 最大值为40w
const maxHosingFund = 40 * 10000

// HouseOld 老算法 data 为每个月缴存的金额
func HouseOld(data []int64) (sum int64) {
	l := len(data)
	for i := 0; i < l; i++ {
		sum += data[i]
	}
	sum *= 20
	if sum > maxHosingFund {
		return maxHosingFund
	}
	return
}

// HouseNew 新算法 data 为每个月缴存的金额(每月取整了的，可能会有总体下来可能会有几块钱的误差)
func HouseNew(data []int64) (sum int64) {
	l := len(data)
	for i := 0; i < l; i++ {
		sum += data[i] * int64(l-i) * 9 / 10
	}
	if sum > maxHosingFund {
		return maxHosingFund
	}
	return
}
