package utils

import (
	"fmt"
	"math"
)

// FormatINR formats amount in Indian Rupees
func FormatINR(amount float64) string {
	return fmt.Sprintf("₹%.2f", amount)
}

// FormatPrice formats price with currency symbol
func FormatPrice(amount float64, currency string) string {
	symbols := map[string]string{
		"INR": "₹",
		"USD": "$",
		"EUR": "€",
		"GBP": "£",
	}

	symbol, exists := symbols[currency]
	if !exists {
		symbol = currency + " "
	}

	return fmt.Sprintf("%s%.2f", symbol, amount)
}

// CalculateDiscount calculates discount amount
func CalculateDiscount(price, discountPercent float64) float64 {
	return price * (discountPercent / 100)
}

// CalculateFinalPrice calculates price after discount
func CalculateFinalPrice(price, discountPercent float64) float64 {
	discount := CalculateDiscount(price, discountPercent)
	return price - discount
}

// CalculateGST calculates GST amount
func CalculateGST(price, gstPercent float64) float64 {
	return price * (gstPercent / 100)
}

// RoundToDecimal rounds float to n decimal places
func RoundToDecimal(value float64, places int) float64 {
	multiplier := math.Pow(10, float64(places))
	return math.Round(value*multiplier) / multiplier
}

// ConvertPaisaToRupees converts paisa to rupees
func ConvertPaisaToRupees(paisa int64) float64 {
	return float64(paisa) / 100.0
}

// ConvertRupeesToPaisa converts rupees to paisa
func ConvertRupeesToPaisa(rupees float64) int64 {
	return int64(rupees * 100)
}