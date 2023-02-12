package logger

import (
	"io"
	"sync"
)

// GatedWriter  is an io.Writer implementation that buffers all of its
// data into an internal buffer until it is told to let data through.
// it's used to log a daemon's stdout/stderr in a more orderly fashion
type GatedWriter struct {
	writer io.Writer
	buf    [][]byte
	flush  bool
	mutex  sync.RWMutex
	// lock semaphore.Semaphore
}

// NewGatedWriter returns a new gated writer
func NewGatedWriter(writer io.Writer) *GatedWriter {
	return &GatedWriter{
		writer: writer,
		// lock:   semaphore.NewBinarySemaphore(),
	}
}

func (w *GatedWriter) Flush() {
	w.mutex.Lock()
	// w.mutex.Wait()
	w.flush = true
	w.mutex.Unlock()
	// w.mutex.Signal()
	for _, p := range w.buf {
		w.Write(p) //nolint:errcheck
	}
	w.buf = nil
}

func (w *GatedWriter) Write(p []byte) (n int, err error) {
	w.mutex.RLock()
	// w.mutex.Wait()
	defer w.mutex.RUnlock()
	if w.flush {
		// w.mutex.Signal()
		return w.writer.Write(p)
	}
	p2 := make([]byte, len(p))
	copy(p2, p)
	w.buf = append(w.buf, p2)
	// w.mutex.Signal()
	return len(p), nil
}
