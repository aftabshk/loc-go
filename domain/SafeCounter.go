package domain

import "sync"

/*
SafeCounter is used to figure out when to close the channel on which the loc metadata is added.
We are using range on that channel, so we have to close it otherwise the range will keep on reading from
channel.
*/
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
