package db

import (
	"AnywhereDoorControlPlane/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

type DataBaseContext struct {
	db *gorm.DB
}

func CreateDataBaseContext() *DataBaseContext {
	db, err := gorm.Open(mysql.Open(getDSN()), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return &DataBaseContext{db: db}
}

func getDSN() string {
	dbIP := os.Getenv("DB_IP")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	if dbIP == "" {
		dbIP = "192.168.25.7"
	}
	if dbPort == "" {
		dbPort = "3306"
	}
	if dbUser == "" {
		dbUser = "root"
	}
	if dbPassword == "" {
		dbPassword = "09251205"
	}
	if dbName == "" {
		dbName = "anywhere_door"
	}

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", dbUser, dbPassword, dbIP, dbPort, dbName, "time_zone='Asia%2FShanghai'")
}

func (ctx *DataBaseContext) QueryPlugins(username string) []model.Plugin {
	var plugins []model.Plugin
	var user model.User
	ctx.db.Model(model.User{Username: username}).First(&user)
	if user.Id >= 0 {
		ctx.db.Where("user_id = ?", user.Id).Find(&plugins)
	}
	return plugins
}

func (ctx *DataBaseContext) QueryPlugin(username string, name string) model.Plugin {
	var plugin model.Plugin
	var user model.User
	ctx.db.Model(model.User{Username: username}).First(&user)
	if user.Id >= 0 {
		ctx.db.Where("user_id = ? AND plugin_name = ?", user.Id, name).First(&plugin)
	}
	return plugin
}

func (ctx *DataBaseContext) QueryImsdks(username string) []model.Imsdk {
	var imsdks []model.Imsdk
	var user model.User
	ctx.db.Model(model.User{Username: username}).First(&user)
	if user.Id >= 0 {
		ctx.db.Where("user_id = ?", user.Id).Find(&imsdks)
	}
	return imsdks
}

func (ctx *DataBaseContext) QueryImsdk(username string, name string) model.Imsdk {
	var imsdk model.Imsdk
	var user model.User
	ctx.db.Model(model.User{Username: username}).First(&user)
	if user.Id >= 0 {
		ctx.db.Where("user_id = ? AND imsdk_name = ?", user.Id, name).First(&imsdk)
	}
	return imsdk
}
