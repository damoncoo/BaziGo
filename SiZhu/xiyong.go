package SiZhu

// 喜用神
import (
	. "github.com/warrially/BaziGo/Common"
	"log"
)

// 天干地支强度测试

// 天干强度表
var TIAN_GAN_QIANG_DU_LIST = [12][10]int{
	//甲   乙   丙   丁   戊   己   庚   辛   壬   癸
	{1200, 1200, 1000, 1000, 1000, 1000, 1000, 1000, 1200, 1200}, //子月
	{1060, 1060, 1000, 1000, 1100, 1100, 1140, 1140, 1100, 1100}, //丑月
	{1140, 1140, 1200, 1200, 1060, 1060, 1000, 1000, 1000, 1000}, //寅月
	{1200, 1200, 1200, 1200, 1000, 1000, 1000, 1000, 1000, 1000}, //卯月
	{1100, 1100, 1060, 1060, 1100, 1100, 1100, 1100, 1040, 1040}, //辰月
	{1000, 1000, 1140, 1140, 1140, 1140, 1060, 1060, 1060, 1060}, //巳月
	{1000, 1000, 1200, 1200, 1200, 1200, 1000, 1000, 1000, 1000}, //午月
	{1040, 1040, 1100, 1100, 1160, 1160, 1100, 1100, 1000, 1000}, //未月
	{1060, 1060, 1000, 1000, 1000, 1000, 1140, 1140, 1200, 1200}, //申月
	{1000, 1000, 1000, 1000, 1000, 1000, 1200, 1200, 1200, 1200}, //酉月
	{1000, 1000, 1040, 1040, 1140, 1140, 1160, 1160, 1060, 1060}, //戌月
	{1200, 1200, 1000, 1000, 1000, 1000, 1000, 1000, 1140, 1140}} //亥月

// 地支强度表
var DI_ZHI_QIANG_DU_LIST = [12][36]int{
	// 子 子 子  丑   丑   丑   寅   寅  寅  卯   卯  卯 辰   辰   辰   巳  巳  巳   午   午  午 未   未   未   申  申  申   酉   酉 酉  戌   戌   戌   亥   亥   亥
	// 癸        己   癸   辛   甲   丙      乙          戊   乙   癸   丙  戊  庚   丁   己     己   乙   丁   庚      壬   辛          戊   辛   丁   壬   甲
	{1000, 0, 0, 530, 300, 200, 798, 360, 0, 1140, 0, 0, 530, 342, 200, 840, 0, 300, 1200, 0, 0, 530, 228, 360, 700, 0, 300, 1000, 0, 0, 530, 300, 240, 700, 342, 0}, // 寅月
	{1000, 0, 0, 500, 300, 200, 840, 360, 0, 1200, 0, 0, 500, 360, 200, 840, 0, 300, 1200, 0, 0, 500, 240, 360, 700, 0, 300, 1000, 0, 0, 500, 300, 240, 700, 360, 0}, // 卯月
	{1040, 0, 0, 550, 312, 230, 770, 318, 0, 1100, 0, 0, 550, 330, 208, 742, 0, 330, 1060, 0, 0, 550, 220, 318, 770, 0, 312, 1100, 0, 0, 550, 330, 212, 728, 330, 0}, // 辰月
	{1060, 0, 0, 570, 318, 212, 700, 342, 0, 1000, 0, 0, 600, 300, 200, 840, 0, 300, 1140, 0, 0, 570, 200, 342, 742, 0, 318, 1060, 0, 0, 570, 318, 228, 742, 300, 0}, // 巳月
	{1000, 0, 0, 600, 300, 200, 700, 360, 0, 1000, 0, 0, 600, 300, 200, 840, 0, 300, 1200, 0, 0, 600, 200, 360, 700, 0, 300, 1000, 0, 0, 600, 300, 240, 700, 300, 0}, // 午月
	{1000, 0, 0, 580, 300, 220, 728, 330, 0, 1040, 0, 0, 580, 312, 200, 798, 0, 330, 1100, 0, 0, 580, 208, 330, 770, 0, 300, 1100, 0, 0, 580, 330, 220, 700, 312, 0}, // 未月
	{1200, 0, 0, 500, 360, 228, 742, 300, 0, 1060, 0, 0, 500, 318, 240, 700, 0, 342, 1000, 0, 0, 500, 212, 300, 798, 0, 360, 1140, 0, 0, 500, 342, 200, 840, 318, 0}, // 申月
	{1200, 0, 0, 500, 360, 248, 700, 300, 0, 1000, 0, 0, 500, 300, 240, 700, 0, 360, 1000, 0, 0, 500, 200, 300, 840, 0, 360, 1200, 0, 0, 500, 360, 200, 840, 300, 0}, // 酉月
	{1060, 0, 0, 570, 318, 232, 700, 342, 0, 1000, 0, 0, 570, 300, 212, 728, 0, 348, 1040, 0, 0, 570, 200, 312, 812, 0, 318, 1160, 0, 0, 570, 348, 208, 724, 300, 0}, // 戌月
	{1140, 0, 0, 500, 342, 200, 840, 318, 0, 1200, 0, 0, 500, 360, 228, 742, 0, 300, 1060, 0, 0, 500, 240, 318, 700, 0, 342, 1000, 0, 0, 500, 300, 212, 798, 360, 0}, // 亥月
	{1200, 0, 0, 500, 360, 200, 840, 300, 0, 1200, 0, 0, 500, 360, 240, 700, 0, 300, 1000, 0, 0, 500, 240, 300, 700, 0, 360, 1000, 0, 0, 500, 300, 200, 840, 360, 0}, // 子月
	{1100, 0, 0, 550, 330, 228, 742, 300, 0, 1060, 0, 0, 550, 318, 220, 700, 0, 342, 1000, 0, 0, 550, 212, 300, 798, 0, 330, 1140, 0, 0, 550, 342, 200, 770, 318, 0}} // 丑月

// 计算喜用神
func CalcXiYong(pSiZhu *TSiZhu) TXiYong {
	var xiyong TXiYong

	// 1. 通过四柱计算天干地支强度
	var wuxing = [5]int{0, 0, 0, 0, 0} // 金木水火土

	// 2. 拿到四柱的月支
	var nMonthZhi = pSiZhu.MonthZhu.Zhi.Value
	log.Println("月支是", nMonthZhi, pSiZhu.MonthZhu.Zhi.Str)

	// 3. 根据四柱天干, 换算强度
	wuxing[pSiZhu.YearZhu.Gan.WuXing.Value] += TIAN_GAN_QIANG_DU_LIST[nMonthZhi][pSiZhu.YearZhu.Gan.Value]
	wuxing[pSiZhu.MonthZhu.Gan.WuXing.Value] += TIAN_GAN_QIANG_DU_LIST[nMonthZhi][pSiZhu.MonthZhu.Gan.Value]
	wuxing[pSiZhu.DayZhu.Gan.WuXing.Value] += TIAN_GAN_QIANG_DU_LIST[nMonthZhi][pSiZhu.DayZhu.Gan.Value]
	wuxing[pSiZhu.HourZhu.Gan.WuXing.Value] += TIAN_GAN_QIANG_DU_LIST[nMonthZhi][pSiZhu.HourZhu.Gan.Value]

	log.Println("计算完毕天干后的五行权值是:", wuxing)

	return xiyong
}