package weather

import (
	"demo/weather/geo"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

var ErrorNotValidURL = errors.New("ERROR_URL_IS_NOT_VALID")
var ErrorNotResponce = errors.New("ERROR_NOT_RESPONCE")
var ErrorCantReadResponce = errors.New("ERROR_CAN'T_READ_RESPONCE")
var ErrorWrongFormat = errors.New("ERROR_WRONG_FORMAT")

func GetWeather(geo geo.GeoData, format int) (string, error) {
	if format < 1 || format > 4 {
		return "", ErrorWrongFormat
	}
	baseUrl, err := url.Parse("https://wttr.in/" + geo.City)
	if err != nil {
		return "", ErrorNotValidURL
	}
	params := url.Values{}
	params.Add("format", fmt.Sprint(format))
	baseUrl.RawQuery = params.Encode()
	resp, err := http.Get(baseUrl.String())
	if err != nil {
		return "", ErrorNotResponce
	}
	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return "", ErrorCantReadResponce
	}
	return string(body), nil
}
