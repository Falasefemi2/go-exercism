package raindrops

import "fmt"

func Convert(number int) string {
	result := ""

	switch {
	case number%3 == 0:
		result += "Pling"
		fallthrough
	case number%5 == 0:
		if number%5 == 0 {
			result += "Plang"
		}
		fallthrough
	case number%7 == 0:
		if number%7 == 0 {
			result += "Plong"
		}
	}

	if result == "" {
		return fmt.Sprintf("%d", number)
	}

	return result
}
