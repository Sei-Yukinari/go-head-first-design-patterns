package observer

type WeatherData struct {
	observers          []Observer
	meteorologicalData MeteorologicalData
}

type MeteorologicalData struct {
	temperature float64
	humidity    float64
	pressure    float64
}

func NewWeatherData() *WeatherData {
	wd := WeatherData{}
	wd.observers = make([]Observer, 0)
	return &wd
}

func (wd *WeatherData) RegisterObserver(o Observer) {
	wd.observers = append(wd.observers, o)
}

func (wd *WeatherData) NotifyObservers() (obs []Observer) {
	for _, o := range wd.observers {
		obs = append(obs, o.Update(wd.meteorologicalData))
	}
	return wd.observers
}

func (wd *WeatherData) RemoveObserver(o Observer) {
	for idx, obs := range wd.observers {
		if obs == o {
			wd.observers = remove(wd.observers, idx)
		}
	}
}

func remove(s []Observer, i int) []Observer {
	return append(s[:i], s[i+1:]...)
}

func (wd *WeatherData) MeasurementsChanged() {
	wd.NotifyObservers()
}

func (wd *WeatherData) SetMeasurements(m MeteorologicalData) {
	wd.meteorologicalData.temperature = m.temperature
	wd.meteorologicalData.humidity = m.humidity
	wd.meteorologicalData.pressure = m.pressure
	wd.MeasurementsChanged()
}
