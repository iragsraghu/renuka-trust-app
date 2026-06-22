package helper

import (
	"github.com/iragsraghu/go-utils/numbertowords"
)

// func RupeesInWords(amount float64) string {
// 	return fmt.Sprintf("%.0f Rupees Only", amount)
// }

func RupeesInWords(amount float64) string {
	words := numbertowords.Rupees(amount)

	return words
}
