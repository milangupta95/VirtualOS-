package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func showWeatherApp() {
	w := myapp.NewWindow("Weather ForeCast")
	var selection string
	combo := widget.NewSelect([]string{"Chennai", "Noida", "Delhi", "Varanasi", "Mumbai", "Kolkata"}, func(value string) {
		selection = value
	})
	label2 := widget.NewLabel(" ")
	label3 := widget.NewLabel(" ")
	label4 := widget.NewLabel(" ")
	label5 := widget.NewLabel("Wind Speed ")
	label6 := widget.NewLabel("Rain ")
	label7 := widget.NewLabel("Temprature ")
	label1 := widget.NewLabelWithStyle("Weather Details", fyne.TextAlignCenter, fyne.TextStyle{Monospace: true, Bold: true})
	img := canvas.NewImageFromFile("images.png")
	img.FillMode = canvas.ImageFillOriginal
	clickMe := widget.NewButton("Click To Get Weather", func() {
		resp, err := http.Get("https://api.openweathermap.org/data/2.5/weather?q=" + selection + "&appid=270a4cb52a56646c4a7f98278f3c19eb")
		if err != nil {
			fmt.Println(err.Error())
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err.Error())
		}
		weather, err := UnmarshalWeather(body)
		if err != nil {
			fmt.Println(err.Error())
		}
		label2.SetText(fmt.Sprintf("%0.2f", weather.Wind.Speed))
		label3.SetText(fmt.Sprintf("%0.2f", weather.Rain.The1H))
		label4.SetText(fmt.Sprintf("%0.2f", weather.Main.Temp))
	})

	w.SetContent(container.NewVBox(
		label1,
		img,
		container.NewGridWithColumns(2,
			combo,
			clickMe,
		),
		container.NewGridWithColumns(2,
			label5,
			label2,
		),
		container.NewGridWithColumns(2,
			label6,
			label3,
		),
		container.NewGridWithColumns(2,
			label7,
			label4,
		),
	),
	)
	w.Show()
}

// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    weather, err := UnmarshalWeather(bytes)
//    bytes, err = weather.Marshal()

func UnmarshalWeather(data []byte) (Weather, error) {
	var r Weather
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Weather) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Weather struct {
	Coord      Coord            `json:"coord"`
	Weather    []WeatherElement `json:"weather"`
	Base       string           `json:"base"`
	Main       Main             `json:"main"`
	Visibility int64            `json:"visibility"`
	Wind       Wind             `json:"wind"`
	Rain       Rain             `json:"rain"`
	Clouds     Clouds           `json:"clouds"`
	Dt         int64            `json:"dt"`
	Sys        Sys              `json:"sys"`
	Timezone   int64            `json:"timezone"`
	ID         int64            `json:"id"`
	Name       string           `json:"name"`
	Cod        int64            `json:"cod"`
}

type Clouds struct {
	All int64 `json:"all"`
}

type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int64   `json:"pressure"`
	Humidity  int64   `json:"humidity"`
}

type Rain struct {
	The1H float64 `json:"1h"`
}

type Sys struct {
	Type    int64  `json:"type"`
	ID      int64  `json:"id"`
	Country string `json:"country"`
	Sunrise int64  `json:"sunrise"`
	Sunset  int64  `json:"sunset"`
}

type WeatherElement struct {
	ID          int64  `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int64   `json:"deg"`
}
