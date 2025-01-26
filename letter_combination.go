// Example 1:
// Input: digits = "23"
// Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]

// Example 2:
// Input: digits = ""
// Output: []

// Example 3:
// Input: digits = "2"
// Output: ["a","b","c"]

func letterCombinations(digits string) []string {
    var result []string = []string{}

	for _, digit := range digits {
		if len(result) > 0 {
			addDigits := getLettersFromDigit(string(digit))
			for _, row := range result {
				result = result[1:]
				for _, col := range addDigits {
					result = append(result, row+col)
				}
			}

		} else {
			result = getLettersFromDigit(string(digit))
		}
	}

	return result
}

func getLettersFromDigit(digit string) []string {
	switch digit {
	case "2":
		return []string{"a", "b", "c"}
	case "3":
		return []string{"d", "e", "f"}
	case "4":
		return []string{"g", "h", "i"}
	case "5":
		return []string{"j", "k", "l"}
	case "6":
		return []string{"m", "n", "o"}
	case "7":
		return []string{"p", "q", "r", "s"}
	case "8":
		return []string{"t", "u", "v"}
	case "9":
		return []string{"w", "x", "y", "z"}
	default:
		return []string{}
	}
}