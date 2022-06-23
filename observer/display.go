package observer

import "fmt"

type DisplayElement interface {
	Display() string
}

type CurrentConditionsDisplay struct {
	subject     Subject
	temperature float64
	humidity    float64
}

func NewCurrentConditionsDisplay(subj Subject) *CurrentConditionsDisplay {
	ccd := CurrentConditionsDisplay{subject: subj}
	ccd.subject.RegisterObserver(&ccd)
	return &ccd
}

func (ccd *CurrentConditionsDisplay) Update(m MeteorologicalData) Observer {
	ccd.temperature = m.temperature
	ccd.humidity = m.humidity
	return ccd
}

func (ccd *CurrentConditionsDisplay) Display() string {
	return fmt.Sprintf("Current Conditions: %.1f degrees F with %.1f%% humidity\n",
		ccd.temperature,
		ccd.humidity,
	)
}

type StatisticsDisplay struct {
	subject      Subject
	temperatures []float64
}

func NewStatisticsDisplay(subj Subject) *StatisticsDisplay {
	sd := StatisticsDisplay{subject: subj}
	sd.temperatures = make([]float64, 0)
	sd.subject.RegisterObserver(&sd)
	return &sd
}

func (sd *StatisticsDisplay) Update(m MeteorologicalData) Observer {
	sd.temperatures = append(sd.temperatures, m.temperature)
	return sd
}

func (sd *StatisticsDisplay) Display() string {
	min := 100000000.0
	max := -100000000.0

	tot := 0.0
	for _, t := range sd.temperatures {
		tot += t
		if t < min {
			min = t
		}
		if t > max {
			max = t
		}
	}
	avg := tot / float64(len(sd.temperatures))
	return fmt.Sprintf("Stats: Avg/Max/Min: %.1f / %.1f / %.1f\n",
		avg,
		max,
		min,
	)
}
