package recordable

import (
	"bufio"
	"os"
	"sync"

	"github.com/snowmerak/msgbuf/log"
	"github.com/snowmerak/msgbuf/log/loglevel"
)

type Stdout struct {
	sync.Mutex
	level  loglevel.LogLevel
	writer *bufio.Writer
}

func NewStdout(level loglevel.LogLevel) log.Writable {
	return &Stdout{
		writer: bufio.NewWriter(os.Stdout),
		level:  level,
	}
}

func (s *Stdout) Write(level loglevel.LogLevel, value []byte) error {
	s.Lock()
	defer s.Unlock()
	if loglevel.Available(s.level, level) {
		s.writer.Write(value)
		s.writer.WriteByte('\n')
		return s.writer.Flush()
	}
	return nil
}

func (s *Stdout) Close() error {
	return nil
}