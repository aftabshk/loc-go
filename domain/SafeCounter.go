package domain

import "sync"

type SafeCounter struct {
	lock  sync.Mutex
	count int
}

func (s *SafeCounter) Inc() {
	s.lock.Lock()
	s.count++
	s.lock.Unlock()
}

func (s *SafeCounter) Dec() {
	s.lock.Lock()
	s.count = s.count - 1
	s.lock.Unlock()
}

func (s *SafeCounter) Count() int {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.count
}
