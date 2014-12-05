// +build OMIT

package sample // OMIT
type Stream struct {
	// some fields
	isConnClosed     bool
	connClosedCond   *sync.Cond
	connClosedLocker sync.Mutex
}

func (s *Stream) Wait() error {
	s.connClosedCond.L.Lock()
	for !s.isConnClosed {
		s.connClosedCond.Wait()
	}
	s.connClosedCond.L.Unlock()
	// some code
}
func (s *Stream) Close() {
	// some code
	s.connClosedCond.L.Lock()
	s.isConnClosed = true
	s.connClosedCond.L.Unlock()
	s.connClosedCond.Broadcast()
}
func (s *Stream) IsClosed() bool {
	return s.isConnClosed
}
