package siglazy

import "sync"

type Singleton struct{}

var (
	instantiated *Singleton
	once         sync.Once
)

func New() *Singleton {
	once.Do(func() {
		instantiated = &Singleton{}
	})
	return instantiated
}
