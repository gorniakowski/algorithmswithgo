package module01

// Reverse will return the provided word in reverse
// order. Eg:
//
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"
//
/* func Reverse(word string) string {
	var result strings.Builder
	for i := len(word) - 1; i >= 0; i-- {
		result.WriteByte(word[i])
	}
	return result.String()
}
*/
//support runes
func Reverse(word string) string {
	var result string
	for _, char := range word {
		result = string(char) + result
	}
	return result
}
