package module01

//Euclidean algorithm iterative

func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

/* //Euclidean algorithm recursive
func GCD(a, b int) int {
	if b == 0 {
		return a
	}
	a, b = b, a%b
	return GCD(a, b)
}
*/

/* var primes = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199}
 //Using factor function, and list of primes
func GCD(a, b int) int {
	factorsA := Factor(primes, a)
	factorsB := Factor(primes, b)
	result := 1
	for _, factorA := range factorsA {
		for i, factorB := range factorsB {
			if factorA == factorB {
				result *= factorB
				factorsB[i] = -1
				break
			}
		}
	}
	return result
} */
