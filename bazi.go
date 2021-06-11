package bazi

import (
	"encoding/json"
	"fmt"
)

// NewBazi 新建八字
func NewBazi(pSolarDate *TSolarDate, nSex int) *TBazi {
	//
	pBazi := &TBazi{
		PSolarDate: pSolarDate,
		NSex:       nSex,
	}
	return pBazi.init()
}

// NewBaziFromLunarDate 新建八字 从农历
func NewBaziFromLunarDate(pLunarDate *TLunarDate, nSex int) *TBazi {
	pBazi := &TBazi{
		PLunarDate: pLunarDate,
		NSex:       nSex,
	}

	return pBazi.init()
}

// GetBazi 旧版八字接口, 八字入口
func GetBazi(nYear, nMonth, nDay, nHour, nMinute, nSecond, nSex int) *TBazi {
	// 先解决时间问题. 然后开始处理八字问题
	pSolarDate := NewSolarDate(nYear, nMonth, nDay, nHour, nMinute, nSecond)
	if pSolarDate == nil {
		return nil
	}

	return NewBazi(pSolarDate, nSex)
}

// TBazi 八字大类
type TBazi struct {
	PSolarDate *TSolarDate `json:"pSolarDate"` // 新历的日期
	PLunarDate *TLunarDate `json:"pLunarDate"` // 农历日期
	PBaziDate  *TBaziDate  `json:"pBaziDate"`  // 八字历
	PSiZhu     *TSiZhu     `json:"pSiZhu"`     // 四柱嗯
	NSex       int         `json:"nSex"`       // 性别1男其他女
	PDaYun     *TDaYun     `json:"pDaYun"`     // 大运
	PQiYunDate *TSolarDate `json:"pQiYunDate"` // 起运时间XX年XX月开始起运
}

// 八字初始化
func (self *TBazi) init() *TBazi {
	// 1. 新农互转
	if self.PSolarDate == nil {
		if self.PLunarDate == nil {
			return nil
		}

		// todo 这里进行新农互转
		// self.pSolarDate = self.pLunarDate
	} else {
		// todo 这里进行新农互转
		self.PLunarDate = self.PSolarDate.ToLunarDate()
	}

	// 1. 拿到新历的情况下, 需要计算八字历
	self.PBaziDate = self.PSolarDate.ToBaziDate()

	// 2. 根据八字历, 准备计算四柱了
	self.PSiZhu = NewSiZhu(self.PSolarDate, self.PBaziDate)

	// 3. 计算大运
	self.PDaYun = NewDaYun(self.PSiZhu, self.NSex)

	// 4. 计算起运时间
	self.PQiYunDate = NewQiYun(self.PDaYun.ShunNi(), self.PBaziDate.PreviousJie().ToSolarDate(), self.PBaziDate.NextJie().ToSolarDate(), self.PSolarDate)

	return self
}

func (self *TBazi) String() string {
	return fmt.Sprintf("%v\n %v\n %v\n%v\n%v \n起运时间%v", self.PSolarDate, self.PLunarDate, self.PBaziDate, self.PSiZhu, self.PDaYun, self.PQiYunDate)
}

func (self *TBazi) Data() string {
	return ObjecToString(self)
}

// SiZhu 四柱
func (self *TBazi) SiZhu() *TSiZhu {
	return self.PSiZhu
}

func ObjecToString(obj interface{}) string {

	jsonByte, _ := json.MarshalIndent(obj, "", " ")
	jsonStr := string(jsonByte)
	return jsonStr
}
