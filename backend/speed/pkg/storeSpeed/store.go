package storeSpeed

import (
	"fmt"
	"net"
	"sync"
)

type Train struct {
	Name string
	Id   int32
	Addr net.Addr
}

type TrainSpeed struct {
	Train Train
	Speed int32
}

type SpeedStore interface {
	Get(t *Train) (speed *TrainSpeed, err error)
	Update(t *TrainSpeed)
	GetAll() []TrainSpeed
}
type speedStore struct {
	speedList []TrainSpeed
	mutex     sync.Mutex
}

func NewStore() SpeedStore {
	return &speedStore{}
}

func (s *speedStore) Get(t *Train) (speed *TrainSpeed, err error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	for _, r := range s.speedList {
		if r.Train == *t {
			return &r, nil
		}
	}
	return nil, fmt.Errorf("Speed Record Not Found\n")
}

func (s *speedStore) Update(t *TrainSpeed) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	for i, r := range s.speedList {
		if r.Train == t.Train {
			s.speedList[i] = *t
			return
		}
	}
	s.speedList = append(s.speedList, *t)
}

func (s *speedStore) GetAll() []TrainSpeed {
	return s.speedList
}
