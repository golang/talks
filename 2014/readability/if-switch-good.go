// +build OMIT

package sample // OMIT

func BrowserHeightBucket(s *session.Event) string {
	size := sizeFromSession(s)
	h := size.GetHeight()
	switch {
	case h <= 0: // HL
		return "null"
	case h <= 480: // HL
		return "small"
	case h <= 640: // HL
		return "medium"
	default: // HL
		return "large"
	}
}
