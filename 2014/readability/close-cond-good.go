// +build OMIT

package sample // OMIT
type Stream struct {
	// some fields
	cc chan struct{} // HL
}

func (s *Stream) Wait() error {
	<-s.cc
	// some code
}
func (s *Stream) Close() {
	// some code
	close(s.cc)
}
func (s *Stream) IsClosed() bool {
	select {
	case <-s.cc:
		return true
	default:
		return false
	}
}
