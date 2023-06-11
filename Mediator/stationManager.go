package Mediator

import "sync"

type StationManager struct {
	isPlatformFree bool
	lock           *sync.Mutex
	trainQueue     []Train
}

func NewStationManager() *StationManager {
	return &StationManager{
		isPlatformFree: true,
		lock:           &sync.Mutex{},
	}
}

func (s *StationManager) CanLand(t Train) bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.isPlatformFree {
		s.isPlatformFree = false
		return true
	}
	s.trainQueue = append(s.trainQueue, t)
	return false
}

func (s *StationManager) NotifyFree() {
	s.lock.Lock()
	defer s.lock.Unlock()
	if !s.isPlatformFree {
		s.isPlatformFree = true
	}

	if len(s.trainQueue) > 0 {
		first := s.trainQueue[0]
		s.trainQueue = s.trainQueue[1:]
		first.PermitArrival()

	}
}
