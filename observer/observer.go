package observer

type Observer interface {
	Update(m MeteorologicalData) Observer
}

type Subject interface {
	RegisterObserver(o Observer)
	RemoveObserver(o Observer)
	NotifyObservers() []Observer
}
