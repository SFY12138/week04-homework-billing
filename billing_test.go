package main

import (
	"testing"
)

// TestCalculateTotalBill 测试计算最终电费
func TestCalculateTotalBill(t *testing.T) {
	testCases := []struct {
		amount      float64
		hour        int
		expected    float64
		description string
	}{
		{100.0, 14, 55.0, "第一档电量，高峰时段"},
		{300.0, 14, 198.0, "第二档电量，高峰时段"},
		{500.0, 14, 418.0, "第三档电量，高峰时段"},
		{200.0, 14, 110.0, "第一档上限，高峰时段"},
	}

	for _, tc := range testCases {
		actual := calculateTotalBill(tc.amount, tc.hour)
		// 使用近似比较，避免浮点数精度问题
		if actual < tc.expected-0.001 || actual > tc.expected+0.001 {
			t.Errorf("%s: expected %.2f, got %.2f", tc.description, tc.expected, actual)
		}
	}
}

// TestCalculateTierBill 测试计算阶梯电费
func TestCalculateTierBill(t *testing.T) {
	testCases := []struct {
		amount   float64
		expected float64
		name     string
	}{
		{100.0, 100.0 * 0.5, "第一档电量"},
		{200.0, 200.0 * 0.5, "第一档上限"},
		{300.0, 200.0*0.5 + 100.0*0.8, "第二档电量"},
		{400.0, 200.0*0.5 + 200.0*0.8, "第二档上限"},
		{500.0, 200.0*0.5 + 200.0*0.8 + 100.0*1.2, "第三档电量"},
	}

	for _, tc := range testCases {
		actual := calculateTierBill(tc.amount)
		if actual != tc.expected {
			t.Errorf("%s: expected %.2f, got %.2f", tc.name, tc.expected, actual)
		}
	}
}

// TestCalculateTimeFactor 测试计算峰谷平调节因子
func TestCalculateTimeFactor(t *testing.T) {
	testCases := []struct {
		hour     int
		expected float64
		name     string
	}{
		{8, 1.1, "高峰时段开始"},
		{14, 1.1, "高峰时段中间"},
		{21, 1.1, "高峰时段结束"},
		{22, 0.8, "低谷时段开始"},
		{0, 0.8, "低谷时段中间"},
		{7, 0.8, "低谷时段结束"},
	}

	for _, tc := range testCases {
		actual := calculateTimeFactor(tc.hour)
		if actual != tc.expected {
			t.Errorf("%s: expected %.2f, got %.2f", tc.name, tc.expected, actual)
		}
	}
}
