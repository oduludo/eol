package utils

type ZippedPair = struct {
	A any
	B any
}
type Zipped = []ZippedPair

// Zip together two arrays and return an array consisting of pairs and an 'ok' boolean.
func Zip[A any, B any](arrayA []A, arrayB []B) (Zipped, bool) {
	result := Zipped{}

	if len(arrayA) != len(arrayB) {
		return result, false
	}

	for i := 0; i < len(arrayA); i++ {
		result = append(result, ZippedPair{arrayA[i], arrayB[i]})
	}

	return result, true
}
