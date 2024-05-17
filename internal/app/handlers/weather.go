package handlers

import (
	"encoding/json"
	"fmt"
	"go-weathermonitor/configs"
	models "go-weathermonitor/internal/app/entities"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)


func HomeHandler(c echo.Context) error{
	fmt.Println("hello")
	return c.String(http.StatusOK, "Server is Running")

}


func GetWeatherData(c echo.Context) error{

	db:=configs.DB()
	var weatherData *models.WeatherData

    city:=c.Param("city")

	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	apiKey:=viper.Get("Apikey")



	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching weather data:", err)
		return c.String(http.StatusInternalServerError, err.Error())
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return c.String(http.StatusInternalServerError, err.Error())
	}


	err = json.Unmarshal(body, &weatherData)
	if err != nil {
		fmt.Println("Error unmarshaling weather data:", err)
		return c.String(http.StatusInternalServerError, err.Error())
	}


   newweatherdata:=&models.Weather{
   Temp:weatherData.Main.Temp,
   Name :weatherData.Name,


   }

   if err:= db.Create(&newweatherdata).Error;err!=nil{
	return c.String(http.StatusInternalServerError, err.Error())
   }

  tempStr := fmt.Sprintf("The current temperature in %s is %.2f°C", weatherData.Name, weatherData.Main.Temp)
  return c.String(http.StatusOK, tempStr)
}
