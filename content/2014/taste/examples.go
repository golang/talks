// +build OMIT

package examples

// IndexOfAny START OMIT
func IndexOfAny(str string, chars []rune) int {
	if len(str) == 0 || len(chars) == 0 {
		return -1
	}
	for i, ch := range str {
		for _, match := range chars {
			if ch == match {
				return i
			}
		}
	}
	return -1
}

// IndexOfAny END OMIT
