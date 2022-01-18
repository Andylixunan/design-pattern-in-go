package observer

import "fmt"

type IObservable interface {
	Add(o IObserver)
	Remove(o IObserver)
	Notify()
}

type IObserver interface {
	Update()
}

type Observable struct {
	observers     []IObserver
	internalState string
}

func (subject *Observable) Add(o IObserver) {
	subject.observers = append(subject.observers, o)
}

func (subject *Observable) Remove(o IObserver) {
	for i, observer := range subject.observers {
		if observer == o {
			subject.observers = append(subject.observers[:i], subject.observers[i:]...)
		}
	}
}

func (subject *Observable) Notify() {
	for _, observer := range subject.observers {
		observer.Update()
	}
}

type ObserverA struct {
	concreteSubject *Observable
}

func (o ObserverA) Update() {
	fmt.Printf("ObserverA observes new subject state: %v\n", o.concreteSubject.internalState)
}

type ObserverB struct {
	concreteSubject *Observable
}

func (o ObserverB) Update() {
	fmt.Printf("ObserverB observes new subject state: %v\n", o.concreteSubject.internalState)
}
