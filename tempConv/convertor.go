package tempconv

type celsius float32
type fahrenheit float32

type TemperatureUnit struct {
	Celsius celsius
	Fahrenheit fahrenheit
}

func (t *TemperatureUnit) ToFahrenheit() fahrenheit {
	
	t.Fahrenheit = fahrenheit(t.Celsius * 9.0/5.0 + 32.0)
	return t.Fahrenheit
}
