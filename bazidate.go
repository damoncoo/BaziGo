package bazi

import (
	"fmt"
)

// NewBaziDate 从新历转成八字历
func NewBaziDate(pSolarDate *TSolarDate) *TBaziDate {
	p := &TBaziDate{}
	p.init(pSolarDate)
	return p
}

// TBaziDate 八字历法
// 八字历法的年  和 新历的 和 农历的都不一样. 八字历法是按照立春为1年. 然后每个节气为月
type TBaziDate struct {
	NYear  int `json:"nYear"`  // 年. 立春
	NMonth int `json:"nMonth"` // 月.
	NDay   int `json:"nDay"`   // 天
	NHour  int `json:"nHour"`  // xiaohsi

	PJieQi       *TJieQi     `json:"pJieQi"`       // 节气名称
	PPreviousJie *TJieQiDate `json:"pPreviousJie"` // 上一个节(气)
	PNextJie     *TJieQiDate `json:"pNextJie"`     // 下一个节(气)
}

func (self *TBaziDate) init(pSolarDate *TSolarDate) *TBaziDate {
	self.NYear = GetLiChunYear(pSolarDate)                      // 拿到八字年, 根据立春来的
	self.PPreviousJie, self.PNextJie = GetJieQiDate(pSolarDate) // 拿到前后两个的日期
	// 节气
	nJieQi := self.PPreviousJie.JieQi
	self.PJieQi = &nJieQi
	// 月
	self.NMonth = self.PJieQi.Month()
	return self
}

func (self *TBaziDate) String() string {
	return fmt.Sprintf("八字历: %4d 年 %02d 月 \n上一个:%v\n下一个:%v",
		self.NYear, self.NMonth, self.PPreviousJie, self.PNextJie)
}

// Year  年. 立春
func (self *TBaziDate) Year() int {
	return self.NYear
}

// Month  月.
func (self *TBaziDate) Month() int {
	return self.NMonth
}

// Day  天
func (self *TBaziDate) Day() int {
	return self.NDay
}

// Hour 小时
func (self *TBaziDate) Hour() int {
	return self.NHour
}

// PreviousJie 上一个节气
func (self *TBaziDate) PreviousJie() *TJieQiDate {
	return self.PPreviousJie
}

// NextJie 下一个节气
func (self *TBaziDate) NextJie() *TJieQiDate {
	return self.PNextJie
}
