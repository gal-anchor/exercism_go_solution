package paasio

import (
	"io"
	"sync"
)

type readCounter struct {
	reader     io.Reader
	totalBytes int64
	numReads   int
	mu         sync.Mutex
}

type writeCounter struct {
	writer     io.Writer
	totalBytes int64
	numWrites  int
	mu         sync.Mutex
}

type readWriteCounter struct {
	readCounter
	writeCounter
}

func NewWriteCounter(writer io.Writer) WriteCounter {
	return &writeCounter{writer: writer}
}

func NewReadCounter(reader io.Reader) ReadCounter {
	return &readCounter{reader: reader}
}

func NewReadWriteCounter(rw io.ReadWriter) ReadWriteCounter {
	return &readWriteCounter{
		readCounter:  readCounter{reader: rw},
		writeCounter: writeCounter{writer: rw},
	}
}

func (rc *readCounter) Read(p []byte) (int, error) {
	n, err := rc.reader.Read(p)
	rc.mu.Lock()
	rc.totalBytes += int64(n)
	rc.numReads++
	rc.mu.Unlock()
	return n, err
}

func (rc *readCounter) ReadCount() (int64, int) {
	rc.mu.Lock()
	defer rc.mu.Unlock()
	return rc.totalBytes, rc.numReads
}

func (wc *writeCounter) Write(p []byte) (int, error) {
	n, err := wc.writer.Write(p)
	wc.mu.Lock()
	wc.totalBytes += int64(n)
	wc.numWrites++
	wc.mu.Unlock()
	return n, err
}

func (wc *writeCounter) WriteCount() (int64, int) {
	wc.mu.Lock()
	defer wc.mu.Unlock()
	return wc.totalBytes, wc.numWrites
}
