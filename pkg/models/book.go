package models

import(
	"github.com/jinzhou/gorm"
	"github.com/te70/stunning-octo-tribble/pkg/config"
)

var db *gorm.DB

type Book struct(
	gorm.model
	Name string `gorm:""json:"name"`
	Author string `json:"author"`
	Publication string `json:"pulication"`
)

func init(){
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&book{})
}

 