package bazi

import "fmt"

// NewDaYun 新大运
func NewDaYun(pSiZhu *TSiZhu, nSex int) *TDaYun {
	p := &TDaYun{}
	p.init(pSiZhu, nSex)
	return p
}

// TDaYun 大运
type TDaYun struct {
	ZhuList  [12]*TZhu `json:"zhuList"`  // 12个大运柱列表
	IsShunNi bool      `json:"isShunNi"` //  顺转还是逆转(true 顺,  false 逆)
}

func (self *TDaYun) init(pSiZhu *TSiZhu, nSex int) *TDaYun {
	for i := 0; i < 12; i++ {
		self.ZhuList[i] = NewZhu() // 新建12个柱
	}

	// 第一判断年柱的阴阳
	yinyang := pSiZhu.YearZhu().ToYinYang()
	// ! 第二判断性别的男女

	// 月柱的干支
	nMonthGanZhi := pSiZhu.MonthZhu().GanZhi().Value()
	// 取出日干十神作为比较
	nDayGan := pSiZhu.DayZhu().Gan().Value()
	fmt.Println(nDayGan)

	//
	for i := 0; i < 12; i++ {
		if yinyang.Value() == nSex {
			self.IsShunNi = true
			self.ZhuList[i].genBaseGanZhi((nMonthGanZhi + 61 + i) % 60)
		} else {
			self.IsShunNi = false
			self.ZhuList[i].genBaseGanZhi((nMonthGanZhi + 59 - i) % 60)

		}
	}

	return self
}

// String
func (self *TDaYun) String() string {
	strResult := "大运:\n"

	for i := 0; i < 12; i++ {
		strResult += self.ZhuList[i].GanZhi().String() + " "
	}

	return strResult
}

// ShunNi 顺逆
func (self *TDaYun) ShunNi() bool {
	return self.IsShunNi
}
