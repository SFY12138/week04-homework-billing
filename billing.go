package main

import (
	"fmt"
	"time"
)

// 定义阶梯阈值和单价常量
const (
	// 阶梯阈值
	FirstTierThreshold  = 200.0
	SecondTierThreshold = 400.0

	// 阶梯单价
	FirstTierPrice  = 0.5
	SecondTierPrice = 0.8
	ThirdTierPrice  = 1.2

	// 峰谷平调节因子
	PeakFactor   = 1.1 // 高峰时段总费率增加10%
	ValleyFactor = 0.8 // 低谷时段总费率减少20%

	// 计费规则版本号
	Version = "1.0.0"
)

// init函数初始化系统
func init() {
	fmt.Printf("计费规则版本号: %s\n", Version)
	fmt.Printf("系统初始化时间: %s\n", time.Now().Format("2006-01-02 15:04:05"))
}

// calculateTierBill 计算阶梯电费
func calculateTierBill(amount float64) float64 {
	var bill float64

	if amount <= FirstTierThreshold {
		// 第一档：0-200度
		bill = amount * FirstTierPrice
	} else if amount <= SecondTierThreshold {
		// 第二档：200-400度
		bill = FirstTierThreshold*FirstTierPrice + (amount-FirstTierThreshold)*SecondTierPrice
	} else {
		// 第三档：400度以上
		bill = FirstTierThreshold*FirstTierPrice + (SecondTierThreshold-FirstTierThreshold)*SecondTierPrice + (amount-SecondTierThreshold)*ThirdTierPrice
	}

	return bill
}

// calculateTimeFactor 计算峰谷平调节因子
func calculateTimeFactor(hour int) float64 {
	if hour >= 8 && hour < 22 {
		// 高峰时段 (8:00-22:00)
		return PeakFactor
	} else {
		// 低谷时段 (22:00-次日8:00)
		return ValleyFactor
	}
}

// calculateTotalBill 计算最终电费
func calculateTotalBill(amount float64, hour int) float64 {
	tierBill := calculateTierBill(amount)
	timeFactor := calculateTimeFactor(hour)
	return tierBill * timeFactor
}

// getHourFromTimeString 从时间字符串中提取小时
func getHourFromTimeString(timeStr string) (int, error) {
	// 解析时间字符串，格式为 "14:00"
	var hour int
	_, err := fmt.Sscanf(timeStr, "%d:%d", &hour, new(int))
	return hour, err
}

// runBilling 运行计费流程
func runBilling() {
	var amount float64
	var timeStr string

	// 引导用户输入用电量
	fmt.Print("请输入用电量（比如 400.00）: ")
	fmt.Scanln(&amount)

	// 验证用电量是否为非负数
	if amount < 0 {
		fmt.Println("输入的用电量不能为负数，请重新运行程序并输入正确的用电量")
		return
	}

	// 引导用户输入用电时段
	fmt.Print("请输入用电时段（比如 14:00）: ")
	fmt.Scanln(&timeStr)

	// 提取小时和分钟
	hour, minute, err := getTimeFromTimeString(timeStr)
	if err != nil {
		fmt.Println("输入的时间格式不正确，请重新运行程序并输入正确的时间格式（比如 14:00）")
		return
	}

	// 验证时间是否在有效范围内
	if hour < 0 || hour > 23 || minute < 0 || minute > 59 {
		fmt.Println("输入的时间不在有效范围内（0:00~23:59），请重新运行程序并输入正确的时间")
		return
	}

	// 计算最终电费
	finalBill := calculateTotalBill(amount, hour)

	// 打印账单明细
	fmt.Println("--- 账单明细 ---")
	fmt.Printf("当前用电: %.2f 度\n", amount)
	fmt.Printf("当前时段: %s 点\n", timeStr)
	fmt.Printf("最终电费: %.2f 元\n", finalBill)
}

// getTimeFromTimeString 从时间字符串中提取小时和分钟
func getTimeFromTimeString(timeStr string) (int, int, error) {
	// 解析时间字符串，格式为 "14:00"
	var hour, minute int
	_, err := fmt.Sscanf(timeStr, "%d:%d", &hour, &minute)
	return hour, minute, err
}

func main() {
	runBilling()
}
