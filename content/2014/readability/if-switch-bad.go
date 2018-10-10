// +build OMIT

package sample // OMIT

func BrowserHeightBucket(s *session.Event) string {
	browserSize := sizeFromSession(s)
	if h := browserSize.GetHeight(); h > 0 { // HL
		browserHeight := int(h)
		if browserHeight <= 480 { // HL
			return "small"
		} else if browserHeight <= 640 { // HL
			return "medium"
		} else {
			return "large"
		}
	} else {
		return "null"
	}
}
