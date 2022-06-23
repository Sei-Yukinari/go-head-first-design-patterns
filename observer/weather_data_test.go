package observer

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWeatherData(t *testing.T) {
	wd := NewWeatherData()
	NewCurrentConditionsDisplay(wd)
	NewStatisticsDisplay(wd)
	wd.SetMeasurements(
		MeteorologicalData{
			temperature: 80,
			humidity:    65,
			pressure:    30.4,
		},
	)
	wd.SetMeasurements(
		MeteorologicalData{
			temperature: 82,
			humidity:    70,
			pressure:    29.2,
		},
	)
	wd.SetMeasurements(
		MeteorologicalData{
			temperature: 78,
			humidity:    90,
			pressure:    29.2,
		},
	)

	for _, n := range wd.observers {
		switch v := n.(type) {
		case *CurrentConditionsDisplay:
			assert.Equal(t, v.temperature, float64(78))
			assert.Equal(t, v.humidity, float64(90))
			fmt.Println(v.Display())
		case *StatisticsDisplay:
			assert.Equal(t, v.temperatures, []float64{80, 82, 78})
			fmt.Println(v.Display())
		}
	}
}

type MockObserver struct {
}

func (mo MockObserver) Update(me MeteorologicalData) Observer {
	return mo
}

func TestRegisterObserver(t *testing.T) {
	wd := NewWeatherData()
	m := MockObserver{}
	t.Run("register", func(t *testing.T) {
		before := len(wd.observers)
		wd.RegisterObserver(m)
		assert.Equal(t, before+1, len(wd.observers))
	})
	t.Run("remove", func(t *testing.T) {
		before := len(wd.observers)
		wd.RemoveObserver(m)
		assert.Equal(t, before-1, len(wd.observers))
	})
}
