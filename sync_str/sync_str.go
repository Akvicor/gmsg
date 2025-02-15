package sync_str

import "sync"

type String struct {
	l   sync.RWMutex
	str string
}

func (s *String) String() string {
	s.l.RLock()
	defer s.l.RUnlock()
	return s.str
}

func (s *String) Append(str string) string {
	s.l.RLock()
	defer s.l.RUnlock()
	return s.str + str
}

func (s *String) Set(str string) {
	s.l.Lock()
	defer s.l.Unlock()
	s.str = str
}

func New(str string) *String {
	return &String{str: str}
}
