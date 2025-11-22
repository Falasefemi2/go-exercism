package meteorology

import "fmt"

type TemperatureUnit int

const (
	Celsius TemperatureUnit = iota
	Fahrenheit
)

func (tu TemperatureUnit) String() string {
	switch tu {
	case Celsius:
		return "°C"
	case Fahrenheit:
		return "°F"
	default:
		return "Unknown"
	}
}

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

func (t Temperature) String() string {
	return fmt.Sprintf("%d %s", t.degree, t.unit.String())
}

type SpeedUnit int

const (
	KmPerHour SpeedUnit = iota
	MilesPerHour
)

func (su SpeedUnit) String() string {
	switch su {
	case KmPerHour:
		return "km/h"
	case MilesPerHour:
		return "mph"
	default:
		return "Unknown"
	}
}

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

func (s Speed) String() string {
	return fmt.Sprintf("%d %s", s.magnitude, s.unit.String())
}

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

func (m MeteorologyData) String() string {
	return fmt.Sprintf(
		"%s: %s, Wind %s at %s, %d%% Humidity",
		m.location,
		m.temperature.String(),
		m.windDirection,
		m.windSpeed.String(),
		m.humidity,
	)
}
