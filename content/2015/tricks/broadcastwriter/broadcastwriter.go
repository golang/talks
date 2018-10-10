package broadcastwriter

import (
	"io"
	"sync"
)

type BroadcastWriter struct {
	sync.Mutex
	writers map[io.WriteCloser]struct{}
}

func (w *BroadcastWriter) AddWriter(writer io.WriteCloser) {
	w.Lock()
	w.writers[writer] = struct{}{}
	w.Unlock()
}

func (w *BroadcastWriter) Write(p []byte) (n int, err error) {
	w.Lock()
	for sw := range w.writers {
		if n, err := sw.Write(p); err != nil || n != len(p) {
			delete(w.writers, sw)
		}
	}
	w.Unlock()
	return len(p), nil
}

// END OMIT

func (w *BroadcastWriter) Clean() error {
	w.Lock()
	for w := range w.writers {
		w.Close()
	}
	w.writers = make(map[io.WriteCloser]struct{})
	w.Unlock()
	return nil
}

func New() *BroadcastWriter {
	return &BroadcastWriter{
		writers: make(map[io.WriteCloser]struct{}),
	}
}
