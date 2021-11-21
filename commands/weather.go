package Commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	Telegram "jbelbo/guidoBot/telegram"
	"log"
	"net/http"
	"os"
	"strings"
)

/*
   "weather": [
     {
       "id": 802,
       "main": "Clouds",
       "description": "scattered clouds",
       "icon": "03n" //https://openweathermap.org/img/wn/03n.png
     }
   ],
   "base": "stations",
   "main": {
     "temp": 300.15,
     "pressure": 1007,
     "humidity": 74,
     "temp_min": 300.15,
     "temp_max": 300.15
   },
*/

type WeatherResponse struct {
	//Coord struct {
	//	Lon float64 `json:"lon"`
	//	Lat float64 `json:"lat"`
	//} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	//Base string `json:"base"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
		Gust  float64 `json:"gust"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	//Dt  int `json:"dt"`
	//Sys struct {
	//	Type    int    `json:"type"`
	//	ID      int    `json:"id"`
	//	Country string `json:"country"`
	//	Sunrise int    `json:"sunrise"`
	//	Sunset  int    `json:"sunset"`
	//} `json:"sys"`
	//Timezone int    `json:"timezone"`
	//ID       int    `json:"id"`
	//Name     string `json:"name"`
	//Cod      int    `json:"cod"`
}

func GetWeather(reqBody *Telegram.WebhookReqBody, responseBody *Telegram.MessageResponse) error {

	city := strings.TrimSpace(strings.TrimPrefix(reqBody.Message.Text, "/weather"))
	apiKey := os.Getenv("OPEN_WEATHER_API_KEY")

	if apiKey == "" {
		log.Fatal("No Open Weather API KEY")
	}

	resp, err := http.Get("https://api.openweathermap.org/data/2.5/weather?units=metric&q=" + city + "&appid=" + apiKey)
	if err != nil {
		fmt.Println("No response from request")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal("Error while reading Open Weather response")
	}
	fmt.Println(string(body))

	var weather WeatherResponse
	err = json.Unmarshal(body, &weather)
	if err != nil {
		fmt.Println("error unmarshalling open weather response: ", err)
	}
	responseBody.Text = formatResponse(weather)

	return nil
}

func formatResponse(response WeatherResponse) string {
	prettyJSON, err := json.MarshalIndent(response, "", "    ")
	if err != nil {
		log.Fatal("Failed to generate json", err)
	}
	return fmt.Sprintf("%s\n", string(prettyJSON))
}
