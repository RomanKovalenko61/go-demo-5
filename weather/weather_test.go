package weather_test

import (
	"demo/weather/geo"
	"demo/weather/weather"
	"strings"
	"testing"
)

func TestGetWeather(t *testing.T) {
	expected := "Moskow"
	geoData := geo.GeoData{
		City: expected,
	}
	format := 3
	result := weather.GetWeather(geoData, format)
	if !strings.Contains(result, expected) {
		t.Errorf("Ожидалось %v, получено %v", expected, result)
	}
}
