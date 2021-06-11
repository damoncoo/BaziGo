package bazi

import "fmt"

// NewSiZhu 新四柱
func NewSiZhu(pSolarDate *TSolarDate, pBaziDate *TBaziDate) *TSiZhu {
	p := &TSiZhu{
		PYearZhu:   NewZhu(),
		PMonthZhu:  NewZhu(),
		PDayZhu:    NewZhu(),
		PHourZhu:   NewZhu(),
		PSolarDate: pSolarDate,
		PBaziDate:  pBaziDate,
	}
	p.init()
	return p
}

// TSiZhu 四柱
type TSiZhu struct {
	PYearZhu    *TZhu        `json:"pYearZhu"`    // 年柱
	PMonthZhu   *TZhu        `json:"pMonthZhu"`   // 月柱
	PDayZhu     *TZhu        `json:"pDayZhu"`     // 日柱
	PHourZhu    *TZhu        `json:"pHourZhu"`    // 时柱
	PHeHuaChong *THeHuaChong `json:"pHeHuaChong"` // 荷花冲
	PSolarDate  *TSolarDate  `json:"pSolarDate"`  // 新历日期
	PBaziDate   *TBaziDate   `json:"pBaziDate"`   // 八字历日期
}

func (self *TSiZhu) init() *TSiZhu {

	// 通过公历 年月日计算日柱
	nDayGan := self.PDayZhu.genDayGanZhi(self.PSolarDate.GetAllDays()).Gan().Value() // 获取日干(日主)
	// 通过小时 获取时柱
	self.PHourZhu.setDayGan(nDayGan).genHourGanZhi(self.PSolarDate.Hour())
	// 通过八字年来获取年柱
	nYearGan := self.PYearZhu.setDayGan(nDayGan).genYearGanZhi(self.PBaziDate.Year()).Gan().Value()
	// 通过年干支和八字月
	self.PMonthZhu.setDayGan(nDayGan).genMonthGanZhi(self.PBaziDate.Month(), nYearGan)

	return self
}

//  genShiShen 计算十神

func (self *TSiZhu) String() string {
	return fmt.Sprintf("四柱:%v %v %v %v\n命盘解析:\n%v(%v)[%v]\t%v(%v)[%v]\t%v(%v)[%v]\t%v(%v)[%v]\t",
		self.PYearZhu.GanZhi(),
		self.PMonthZhu.GanZhi(),
		self.PDayZhu.GanZhi(),
		self.PHourZhu.GanZhi(),
		self.PYearZhu.Gan(), self.PYearZhu.Gan().ToWuXing(), self.PYearZhu.ShiShen(),
		self.PMonthZhu.Gan(), self.PMonthZhu.Gan().ToWuXing(), self.PMonthZhu.ShiShen(),
		self.PDayZhu.Gan(), self.PDayZhu.Gan().ToWuXing(), "主",
		self.PHourZhu.Gan(), self.PHourZhu.Gan().ToWuXing(), self.PHourZhu.ShiShen(),
	) + fmt.Sprintf("\n%v(%v)   \t%v(%v)   \t%v(%v)   \t%v(%v) \n",
		self.PYearZhu.Zhi(), self.PYearZhu.Zhi().ToWuXing(),
		self.PMonthZhu.Zhi(), self.PMonthZhu.Zhi().ToWuXing(),
		self.PDayZhu.Zhi(), self.PDayZhu.Zhi().ToWuXing(),
		self.PHourZhu.Zhi(), self.PHourZhu.Zhi().ToWuXing(),
	) + fmt.Sprintf("藏干:\n%v   \t%v    \t%v    \t%v\n",
		self.PYearZhu.CangGan(),
		self.PMonthZhu.CangGan(),
		self.PDayZhu.CangGan(),
		self.PHourZhu.CangGan(),
	) + fmt.Sprintf("纳音:\n%v   \t%v    \t%v    \t%v\n",
		self.PYearZhu.GanZhi().ToNaYin(),
		self.PMonthZhu.GanZhi().ToNaYin(),
		self.PDayZhu.GanZhi().ToNaYin(),
		self.PHourZhu.GanZhi().ToNaYin(),
	)
}

// YearZhu 返回年柱
func (self *TSiZhu) YearZhu() *TZhu {
	return self.PYearZhu
}

// MonthZhu 返回月柱
func (self *TSiZhu) MonthZhu() *TZhu {
	return self.PMonthZhu
}

// DayZhu 返回日柱
func (self *TSiZhu) DayZhu() *TZhu {
	return self.PDayZhu
}

// HourZhu 返回时柱
func (self *TSiZhu) HourZhu() *TZhu {
	return self.PHourZhu
}
