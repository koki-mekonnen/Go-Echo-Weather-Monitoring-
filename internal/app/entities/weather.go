package entities

import (
	"time"
"github.com/google/uuid"
	"gorm.io/gorm"
)

type Weather struct {
	ID        string    `gorm:"primary_key" json:"id,omitempty"`
    Temp float64 `json:"temp"`
   Name string `json:"name"`
    CreatedAt time.Time `gorm:"not null" json:"created_at" sql:"DEFAULT:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at" sql:"DEFAULT:CURRENT_TIMESTAMP"`
}


func (base *Weather) BeforeCreate(scope *gorm.DB) error {
	uuid, err := uuid.New().MarshalText()
	if err != nil {
		return err
	}
	base.ID = string(uuid)
	return nil
}

type WeatherData struct {
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
	Name string `json:"name"`
}
