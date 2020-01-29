package module01

// BaseToDec takes in a number and the base it is currently
// in and returns the decimal equivalent as an integer.
//
// Eg:
//
//   BaseToDec("E", 16) => 14
//   BaseToDec("1110", 2) => 14
//
func BaseToDec(value string, base int) int {
	var result, digit int
	multi := 1
	for i := len(value) - 1; i >= 0; i-- {
		elem := value[i]
		if elem < 65 {
			digit = int(elem - '0')
		} else {
			digit = int(elem - 55)
		}
		result += digit * multi
		multi *= base

	}

	return result
}
