package kuiperbelt

import (
	"bytes"
)

type TestSession struct {
	*bytes.Buffer
	send chan Message
	key  string
}

func (s *TestSession) Key() string {
	return s.key
}

func (s *TestSession) Send() chan<- Message {
	return s.send
}

func (s *TestSession) Close() error {
	return nil
}
