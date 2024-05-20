package life

import (
	"fmt"
	"math"
)

type HosingFund struct {
	// 每个月的公积金
	data []int64
	// 第二个人，每个月的公积金
	data2 []int64
	// 公积金最大值为40w
	maxHosingFund int64
	// 贷款总月数
	month int64
	// 公积金利率
	rateFond float64
	// 商贷利率
	rateBusiness float64
	// 商贷最高时的利率，用来对比现在可以少还多少
	rateMaxBusiness float64
	// 打印的行数(方便查看)
	printNum int64
	// 还款方式，默认等额本息
	repayment int64
	// 贷款总额
	loan int64
}

// NewHosingFund 返回一个默认对象
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
		rateFond:      2.85 / 100 / 12,
		rateBusiness:  5.625 / 100 / 12,
	}
}

// WithMaxHosingFund 最高公积金贷款额度（也可以理解为打算公积金贷款多少）
func WithMaxHosingFund(i int64) HosingFundOption {
	return func(f *HosingFund) {
		f.maxHosingFund = i
	}
}

// WithData 第一个人的公积金缴纳情况
func WithData(i []int64) HosingFundOption {
	return func(f *HosingFund) {
		f.data = i
	}
}

// WithData2 第二人的公积金缴纳情况
func WithData2(i []int64) HosingFundOption {
	return func(f *HosingFund) {
		f.data2 = i
	}
}

// WithMonth 贷款多少个月
func WithMonth(i int64) HosingFundOption {
	return func(f *HosingFund) {
		f.month = i
	}
}

// WithRateFund 公积金利率
func WithRateFund(i float64) HosingFundOption {
	return func(f *HosingFund) {
		f.rateFond = i / 100 / 12
	}
}

// WithRateBusiness 商贷利率
func WithRateBusiness(i float64) HosingFundOption {
	return func(f *HosingFund) {
		f.rateBusiness = i / 100 / 12
	}
}

// WithRateMaxBusiness 需要比较的比较高的商贷利率
func WithRateMaxBusiness(i float64) HosingFundOption {
	return func(f *HosingFund) {
		f.rateMaxBusiness = i / 100 / 12
	}
}

func WithPrintNum(i int64) HosingFundOption {
	return func(f *HosingFund) {
		f.printNum = i
	}
}

// WithRepayment 还款方式
func WithRepayment(i int64) HosingFundOption {
	// 0 本息, 1 本金, 乱传还是算本息
	if i > 1 {
		i = 0
	}
	return func(f *HosingFund) {
		f.repayment = i
	}
}

// WithLoan 打算贷款的总金额
func WithLoan(i int64) HosingFundOption {
	return func(f *HosingFund) {
		f.loan = i
	}
}

// HouseOld 公积金贷款额度老算法
func (h *HosingFund) HouseOld() (sum int64) {
	sum = houseOld(h.data) + houseOld(h.data2)
	if sum > h.maxHosingFund {
		return h.maxHosingFund
	}
	return
}

// houseOld 老算法
func houseOld(data []int64) (sum int64) {
	l := len(data)
	for i := 0; i < l; i++ {
		sum += data[i]
	}
	sum *= 20
	return
}

// HouseNew 新算法 data 为每个月缴存的金额(取整了的，可能会有总体下来可能会有几毛钱的误差)
func (h *HosingFund) HouseNew() (sum int64) {
	sum = houseNew(h.data) + houseNew(h.data2)
	if sum > h.maxHosingFund {
		return h.maxHosingFund
	}
	return
}

// 新算法
func houseNew(data []int64) (sum int64) {
	l := len(data)
	for i := 0; i < l; i++ {
		sum += data[i] * int64(l-i)
	}
	sum = sum * 9 / 10
	return
}

// MoreUse 公积金算法变了后比以前多花的钱（满贷后就不会产生多花的了）
func (h *HosingFund) MoreUse() (more int64, str string) {
	// 需要从公积金转为商贷的金额
	houseMore := h.HouseOld() - h.HouseNew()
	// 这些钱产生的额外利息
	if h.repayment == 0 {
		// 本息
		more, str = h.interestMore(float64(houseMore))
	} else {
		// 本金
		more, str = h.capitalMore(float64(houseMore))
	}
	return
}

// 等额本金 (每个月还相同的本金，越还越少)
func (h *HosingFund) capitalMore(houseMore float64) (more int64, str string) {
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

func (h *HosingFund) interestMore(m float64) (more int64, str string) {
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

// InterestMonth 等额本息 每月还款额度
// 等额本息（每月还款金额一致）计算公式 m * (R*(1+R)^N)/((1+R)^N-1)
// 月还款本息=贷款总额×月利率×（1+月利率）的还款期数次方÷[(1+月利率)的还款期数次方-1]
// 推导过程 http://www.baiozhuntuixing.com/p/103.html
func (h *HosingFund) InterestMonth() (res float64) {
	m := float64(h.loan)
	N := float64(h.month)
	moneyFond := float64(h.HouseNew())
	moneyBusiness := m - moneyFond
	powRB := math.Pow(1+h.rateBusiness, N)
	powRF := math.Pow(1+h.rateFond, N)
	// 商贷每月 等额本息 需要还的钱
	businessMon := moneyBusiness * (h.rateBusiness * powRB) / (powRB - 1)
	// 公积金每月 等额本息 需要还的钱
	fondMon := moneyFond * (h.rateFond * powRF) / (powRF - 1)
	res = fondMon + businessMon
	return
}

// CapitalMonth 等额本金 每月还款额度 (每个月还相同的本金，越还越少)
func (h *HosingFund) CapitalMonth() (res []float64) {
	m := float64(h.loan)
	moneyFond := float64(h.HouseNew())
	moneyBusiness := m - moneyFond
	// 每个月还的本金(公积金)
	everyFond := moneyFond / float64(h.month)
	// 每个月还的本金(商贷)
	everyBusiness := moneyBusiness / float64(h.month)
	mon := int(h.month)
	// 从基金转为商贷的 每个月金额 (按照整数算的，所以可能会有小数点的误差)
	for i := 0; i < mon; i++ {
		mf := everyFond * float64(mon-i) * h.rateFond
		mb := everyBusiness * float64(mon-i) * h.rateBusiness
		mall := mf + mb + everyFond + everyBusiness
		res = append(res, mall)
	}
	return
}

func (h *HosingFund) InterestMaxReduceMonth() (now float64, old float64) {
	now = h.InterestMonth()
	h.rateBusiness = h.rateMaxBusiness
	old = h.InterestMonth()
	return
}
