package support

import (
	"regexp"
	"strings"
)

var romanNumerals = map[int]string{
	1: "I", 2: "II", 3: "III", 4: "IV", 5: "V",
	6: "VI", 7: "VII", 8: "VIII", 9: "IX", 10: "X",
	40: "XL", 50: "L", 90: "XC", 100: "C",
	400: "CD", 500: "D", 900: "CM", 1000: "M",
}

var arabicToRoman = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
	"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
}

func ToRoman(num int) string {
	if num < 1 {
		panic("Result is less than 1, cannot represent as Roman numeral")
	}

	var roman strings.Builder

	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	for _, value := range values {
		for num >= value {
			num -= value

			roman.WriteString(romanNumerals[value])
		}
	}
	return roman.String()
}

func ToArabic(roman string) int {
	if val, ok := arabicToRoman[roman]; ok {
		return val
	}

	panic("Invalid Roman numeral input")
}

func IsRoman(s string) bool {
	return regexp.MustCompile(`^[IVXLCDM]+$`).MatchString(s)
}

func IsComplexFormat(a, b string) bool {
	return IsRoman(a) && !IsRoman(b)
}
