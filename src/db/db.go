package db

import (
	"AnywhereDoorControlPlane/constant/db"
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
	database, err := gorm.Open(mysql.Open(getDSN()), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return &DataBaseContext{db: database}
}

func getDSN() string {
	dbIP := os.Getenv(db.EnvDbIp)
	dbPort := os.Getenv(db.EnvDbPort)
	dbUser := os.Getenv(db.EnvDbUser)
	dbPassword := os.Getenv(db.EnvDbPassword)
	dbName := os.Getenv(db.EnvDbName)
	return fmt.Sprintf(db.DsnTemplate, dbUser, dbPassword, dbIP, dbPort, dbName, db.TimeZone)
}

func (ctx *DataBaseContext) QueryUser(username string) model.User {
	user := model.User{}
	ctx.db.Where(db.QueryUserSQLTemplate, username).First(&user)
	if user.Username != username {
		user.Id = -1
	}
	return user
}

func (ctx *DataBaseContext) QueryPlugin(username string, name string) model.Plugin {
	plugin := model.Plugin{}
	user := ctx.QueryUser(username)
	if user.Id >= 0 {
		ctx.db.Where(db.QueryPluginSQLTemplate, user.Id, name).First(&plugin)
	}
	if plugin.UserId != user.Id && plugin.PluginName != name {
		plugin.Id = -1
	}
	return plugin
}

func (ctx *DataBaseContext) QueryImsdk(username string, name string) model.Imsdk {
	imsdk := model.Imsdk{}
	user := ctx.QueryUser(username)
	if user.Id >= 0 {
		ctx.db.Where(db.QueryImsdkSQLTemplate, user.Id, name).First(&imsdk)
	}
	if imsdk.UserId != user.Id && imsdk.ImsdkName != name {
		imsdk.Id = -1
	}
	return imsdk
}

func (ctx *DataBaseContext) QueryPluginConfig(username string, pluginName string, configKey string) model.Config {
	config := model.Config{}
	user := ctx.QueryUser(username)
	plugin := ctx.QueryPlugin(username, pluginName)
	if user.Id >= 0 && plugin.Id > 0 {
		ctx.db.Where(db.QueryPluginConfigSQLTemplate, user.Id, plugin.Id, configKey).First(&config)
	}
	if config.UserId == user.Id && config.TargetId != plugin.Id && config.ConfigKey != configKey {
		config.Id = -1
	}
	return config
}

func (ctx *DataBaseContext) QueryImsdkConfig(username string, imsdkName string, configKey string) model.Config {
	config := model.Config{}
	user := ctx.QueryUser(username)
	imsdk := ctx.QueryImsdk(username, imsdkName)
	if user.Id >= 0 && imsdk.Id > 0 {
		ctx.db.Where(db.QueryImsdkConfigSQLTemplate, user.Id, imsdk.Id, configKey).First(&config)
	}
	if config.UserId == user.Id && config.TargetId != imsdk.Id && config.ConfigKey != configKey {
		config.Id = -1
	}
	return config
}

func (ctx *DataBaseContext) InsertLog(log model.Log) (int, error) {
	// gorm会在创建、更新、删除的时候自动开启事务，所以这里不需要手动开启事务了(前提是没有在创建db的时候，手动禁用了事务)
	// @see https://gorm.io/zh_CN/docs/transactions.html
	res := ctx.db.Create(&log)
	if res.RowsAffected < 1 {
		return -1, res.Error
	}
	return log.Id, nil
}

func (ctx *DataBaseContext) InsertImsdk(imsdk model.Imsdk) (int, error) {
	res := ctx.db.Create(&imsdk)
	if res.RowsAffected < 1 {
		return -1, res.Error
	}
	return imsdk.Id, nil
}

func (ctx *DataBaseContext) InsertPlugin(plugin model.Plugin) (int, error) {
	res := ctx.db.Create(&plugin)
	if res.RowsAffected < 1 {
		return -1, res.Error
	}
	return plugin.Id, nil
}
