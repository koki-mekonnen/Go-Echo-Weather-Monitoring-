package configs

import (
	"fmt"
	models "go-weathermonitor/internal/app/entities"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/spf13/viper"
)


var database *gorm.DB
var e error


func Databaseinit(){

 viper.SetConfigFile(".env")
 viper.ReadInConfig()

host :=viper.Get("Host")
user:=viper.Get("User")
password:=viper.Get("Password")
dbName:=viper.Get("dbName")

port:=viper.GetInt("Port")


fmt.Println(host, user, password, dbName, port)

dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable ", host, user, password, dbName, port)
	database, e = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if e != nil {
		panic(e)
}

if e = database.AutoMigrate(
		&models.Weather{},

	); e != nil {
		panic(e)
	}

}

func DB() *gorm.DB {
	return database
}





