package signorm

type Singleton struct{}

var singleton *Singleton

func init() {
	singleton = new(Singleton)
}

func New() *Singleton {
	return singleton
}
