package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var DBS *gorm.DB
func InitMysql()error {
	mysql := Cfg.Get("databases.mysql").(map[string]interface{})
	d, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&timeout=5s&loc=Asia%%2FShanghai", mysql["username"], mysql["password"], mysql["host"], mysql["dbname"]))
	if err != nil {
		return err
	}
	fmt.Println(d)
	d.SingularTable(true)
	d.LogMode(false)
	d.DB().SetMaxOpenConns(500)
	d.DB().SetMaxIdleConns(10)
	d.DB().SetConnMaxLifetime(time.Hour)
	DBS = d
	return nil
}