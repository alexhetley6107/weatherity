package weather

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func GetWeather(city, apiKey string) (string, error) {
	baseURL := "https://api.openweathermap.org/data/2.5/weather"

	params := url.Values{}
	params.Add("q", city)

	params.Add("appid", apiKey)
	params.Add("units", "metric")
	params.Add("lang", "ru")

	resp, err := http.Get(fmt.Sprintf("%s?%s", baseURL, params.Encode()))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("город '%s' не найден", city)
	}

	var data WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", err
	}

	return fmt.Sprintf(
		"🌍 %s\n🌡 Температура: %.1f°C\n💧 Влажность: %d%%\n☁️  %s",
		data.Name,
		data.Main.Temp,
		data.Main.Humidity,
		data.Weather[0].Description,
	), nil
}
