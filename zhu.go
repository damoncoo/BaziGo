package bazi

import "fmt"

// TZhu 柱
type TZhu struct {
	PGanZhi  *TGanZhi  `json:"pGanZhi"`  // 干支
	PGan     *TGan     `json:"pGan"`     // 天干
	PZhi     *TZhi     `json:"pZhi"`     // 地支
	PCangGan *TCangGan `json:"pCangGan"` // 藏干
	PShiShen *TShiShen `json:"pShiShen"` // 十神
	NDayGan  int       `json:"nDayGan"`  // 日干值
}

// NewZhu 新建柱子
func NewZhu() *TZhu {
	return &TZhu{}
}

func (self *TZhu) String() string {
	return fmt.Sprintf("%v", self.PGanZhi)
}

// 设置日干值
func (self *TZhu) setDayGan(nDayGan int) *TZhu {
	self.NDayGan = nDayGan
	return self
}

// 生成藏干
func (self *TZhu) genCangGan() {
	// 生成藏干数据
	if self.PZhi != nil {
		self.PCangGan = NewCangGan(self.NDayGan, self.PZhi)
	}
}

// 生成十神
func (self *TZhu) genShiShen() {
	self.PShiShen = NewShiShenFromGan(self.NDayGan, self.PGan)
}

//
func (self *TZhu) genBaseGanZhi(nGanZhi int) *TZhu {
	// 直接设置成品干支
	self.PGanZhi = NewGanZhi(nGanZhi)
	// 拆分干支
	// 获得八字年的干0-9 对应 甲到癸
	// 获得八字年的支0-11 对应 子到亥
	self.PGan, self.PZhi = self.PGanZhi.ExtractGanZhi()

	return self
}

// 生成年干支
func (self *TZhu) genYearGanZhi(nYear int) *TZhu {
	// 通过年获取干支
	// 获得八字年的干支，0-59 对应 甲子到癸亥
	self.PGanZhi = NewGanZhiFromYear(nYear)
	// 拆分干支
	// 获得八字年的干0-9 对应 甲到癸
	// 获得八字年的支0-11 对应 子到亥
	self.PGan, self.PZhi = self.PGanZhi.ExtractGanZhi()

	// 在这里计算藏干
	self.genCangGan()
	self.genShiShen()
	return self
}

//
func (self *TZhu) genMonthGanZhi(nMonth int, nYearGan int) *TZhu {
	// 根据口诀从本年干数计算本年首月的干数
	switch nYearGan {
	case 0, 5:
		// 甲己 丙佐首
		nYearGan = 2
	case 1, 6:
		// 乙庚 戊为头
		nYearGan = 4
	case 2, 7:
		// 丙辛 寻庚起
		nYearGan = 6
	case 3, 8:
		// 丁壬 壬位流
		nYearGan = 8
	case 4, 9:
		// 戊癸 甲好求
		nYearGan = 0
	}

	// 计算本月干数
	nYearGan += ((nMonth - 1) % 10)

	// 拆干
	self.PGan = NewGan(nYearGan % 10)
	self.PZhi = NewZhi((nMonth - 1 + 2) % 12)

	// 组合干支
	self.PGanZhi = CombineGanZhi(self.PGan, self.PZhi)
	// 在这里计算藏干
	self.genCangGan()
	self.genShiShen()
	return self
}

func (self *TZhu) genDayGanZhi(nAllDays int) *TZhu {

	// 通过总天数来获取
	// 获得八字年的干支，0-59 对应 甲子到癸亥
	self.PGanZhi = NewGanZhiFromDay(nAllDays)
	// 拆分干支
	// 获得八字年的干0-9 对应 甲到癸
	// 获得八字年的支0-11 对应 子到亥
	self.PGan, self.PZhi = self.PGanZhi.ExtractGanZhi()

	// 直接保存日干
	self.setDayGan(self.PGan.Value())

	// 在这里计算藏干
	self.genCangGan()
	self.genShiShen()
	return self
}

func (self *TZhu) genHourGanZhi(nHour int) *TZhu {
	// 取出日干
	nGan := self.NDayGan

	// 24小时校验
	nHour %= 24
	if nHour < 0 {
		nHour += 24
	}

	nZhi := 0
	if nHour == 23 {
		// 次日子时
		nGan = (nGan + 1) % 10
	} else {
		nZhi = (nHour + 1) / 2
	}

	// Gan 此时是本日干数，根据规则换算成本日首时辰干数
	if nGan >= 5 {
		nGan -= 5
	}

	// 计算此时辰干数
	nGan = (2*nGan + nZhi) % 10

	self.PGan = NewGan(nGan)
	self.PZhi = NewZhi(nZhi)

	// 组合干支
	self.PGanZhi = CombineGanZhi(self.PGan, self.PZhi)

	// 在这里计算藏干
	self.genCangGan()
	self.genShiShen()
	return self
}

// Gan 获取干
func (self *TZhu) Gan() *TGan {
	return self.PGan
}

// Zhi 获取支
func (self *TZhu) Zhi() *TZhi {
	return self.PZhi
}

// GanZhi 获取干支
func (self *TZhu) GanZhi() *TGanZhi {
	return self.PGanZhi
}

// ToYinYang 从柱里获取阴阳 (阴 == 0,  阳 == 1)
func (self *TZhu) ToYinYang() *TYinYang {
	return NewYinYangFromZhu(self)
}

// CangGan 获取藏干
func (self *TZhu) CangGan() *TCangGan {
	return self.PCangGan
}

// ShiShen 获取十神
func (self *TZhu) ShiShen() *TShiShen {
	return self.PShiShen
}
