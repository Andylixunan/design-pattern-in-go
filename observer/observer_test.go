package observer

import "testing"

func TestObserverPattern(t *testing.T) {
	subject := &Observable{}
	ob1 := ObserverA{
		concreteSubject: subject,
	}
	ob2 := ObserverB{
		concreteSubject: subject,
	}
	subject.Add(ob1)
	subject.Add(ob2)
	subject.internalState = "new state"
	subject.Notify()
}
