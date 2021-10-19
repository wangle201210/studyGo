package life

import (
	"fmt"
	"math"
)

type HosingFund struct {
	// 每个月的公积金
	data []int64
	// 公积金最大值为40w
	maxHosingFund int64
	// 贷款总月数
	month int64
	// 公积金利率
	rateFond float64
	// 商贷利率
	rateBusiness float64
	// 打印的行数
	printNum int64
	// 还款方式，默认等额本息
	repayment int64
}

func NewHosingFund(data []int64, opt ...HosingFundOption) *HosingFund {
	hf := getDefault()
	hf.data = data
	for _, option := range opt {
		option(hf)
	}
	return hf
}

type HosingFundOption func(*HosingFund)

func getDefault() *HosingFund {
	return &HosingFund{
		maxHosingFund: 40 * 10000,
		month:         30 * 12,
		rateFond:      3.25 / 100 / 12,
		rateBusiness:  6.125 / 100 / 12,
	}
}

func WithMaxHosingFund(i int64) HosingFundOption {
	return func(f *HosingFund) {
		f.maxHosingFund = i
	}
}

func WithMonth(i int64) HosingFundOption {
	return func(f *HosingFund) {
		f.month = i
	}
}

func WithRateFund(i float64) HosingFundOption {
	return func(f *HosingFund) {
		f.rateFond = i / 12
	}
}

func WithRateBusiness(i float64) HosingFundOption {
	return func(f *HosingFund) {
		f.rateBusiness = i / 12
	}
}

func WithPrintNum(i int64) HosingFundOption {
	return func(f *HosingFund) {
		f.printNum = i
	}
}

func WithRepayment(i int64) HosingFundOption {
	// 0 本息, 1 本金, 乱传还是算本息
	if i > 1 {
		i = 0
	}
	return func(f *HosingFund) {
		f.repayment = i
	}
}

// HouseOld 老算法 data 为每个月缴存的金额
func (h *HosingFund) HouseOld() (sum int64) {
	l := len(h.data)
	for i := 0; i < l; i++ {
		sum += h.data[i]
	}
	sum *= 20
	if sum > h.maxHosingFund {
		return h.maxHosingFund
	}
	return
}

// HouseNew 新算法 data 为每个月缴存的金额(取整了的，可能会有总体下来可能会有几毛钱的误差)
func (h *HosingFund) HouseNew() (sum int64) {
	l := len(h.data)
	for i := 0; i < l; i++ {
		sum += h.data[i] * int64(l-i)
	}
	sum = sum * 9 / 10
	if sum > h.maxHosingFund {
		return h.maxHosingFund
	}
	return
}

// 等额本金
func (h *HosingFund) MoreUse() (more int64, str string) {
	// 需要从公积金转为商贷的金额
	houseMore := h.HouseOld() - h.HouseNew()
	// 这些钱产生的额外利息
	if h.repayment == 0 {
		// 本息
		more, str = h.interest(float64(houseMore))
	} else {
		// 本金
		more, str = h.capital(float64(houseMore))
	}
	return
}

// 等额本金
func (h *HosingFund) capital(houseMore float64) (more int64, str string) {
	mon := int(h.month)
	printNum := int(h.printNum)
	// 从基金转为商贷的 每个月金额 (按照整数算的，所以可能会有小数点的误差)
	m := houseMore / float64(h.month)
	var mf float64
	for i := 0; i < mon; i++ {
		m := m * float64(mon-i) * (h.rateBusiness - h.rateFond)
		mf += m
		if i < printNum || mon-i <= printNum {
			str += fmt.Sprintf("在第(%d)个月产生的额外利息为(%d)\n此时一共需要多交的钱为(%d)\n", i+1, int64(m), int64(more))
		}
		if printNum != 0 && i == printNum {
			str += fmt.Sprintf("... ... ...\n... ... ...\n... ... ...\n")
		}
	}
	more = int64(mf)
	return
}

// 等额本息计算公式 m * (R*(1+R)^N)/((1+R)^N-1)
// 月还款本息=贷款总额×月利率×（1+月利率）的还款期数次方÷[(1+月利率)的还款期数次方-1]
// 推导过程 http://www.baiozhuntuixing.com/p/103.html
func (h *HosingFund) interest(m float64) (more int64, str string) {
	N := float64(h.month)
	powRF := math.Pow(1+h.rateFond, N)
	powRB := math.Pow(1+h.rateBusiness, N)
	// 公积金每月 等额本息 需要还的钱
	rateFond := m * (h.rateFond * powRF) / (powRF - 1)
	// 商贷每月 等额本息 需要还的钱
	rateBusiness := m * (h.rateBusiness * powRB) / (powRB - 1)
	add := rateBusiness - rateFond
	more = int64(add * N)
	if h.printNum != 0 {
		str += fmt.Sprintf("每月多还(%d)", int64(add))
	}
	return
}
