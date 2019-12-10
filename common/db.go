package common

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	Name         string `json:"name"`
	User         string `json:"user"`
	Password     string `json:"password"`
	Urls         string `json:"urls"`
	MaxIdleConns int    `json:"maxidleconns"`
	MaxOpenConns int    `json:"maxopenconns"`
	MaxLiftTime  int    `json:"maxlifttime"`
}

func dbInit() (err error) {
	var dbMap map[string]string
	dbMap, err = beego.AppConfig.GetSection("database")
	if err != nil {
		return
	}

	if v, ok := dbMap["default"]; ok {
		config := Database{}
		err = json.Unmarshal([]byte(v), &config)
		if err != nil {
			return
		}
		config.Name = "default"

		if _, err = RegisterOrm(config); err != nil {
			return
		}
	}

	for k, v := range dbMap {
		if k == "default" {
			continue
		}

		config := Database{}
		err = json.Unmarshal([]byte(v), &config)
		if err != nil {
			return
		}
		config.Name = k

		if _, err = RegisterOrm(config); err != nil {
			return
		}

	}

	return
}

func RegisterOrm(config Database) (o orm.Ormer, err error) {
	err = orm.RegisterDriver("mysql", orm.DRMySQL)
	if err != nil {
		return
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4", config.User, config.Password, config.Urls, config.Name)
	err = orm.RegisterDataBase(config.Name, "mysql", dsn, config.MaxIdleConns, config.MaxOpenConns)
	if err != nil {
		return
	}

	db, err := orm.GetDB(config.Name)
	if err != nil {
		return
	}

	if db == nil {
		return
	}

	db.SetConnMaxLifetime(time.Second * time.Duration(config.MaxLiftTime))

	o = orm.NewOrm()
	_ = o.Using(config.Name)

	dbOrms[config.Name] = o

	return
}

var (
	dbOrms map[string]orm.Ormer
)

func DBInit() (err error) {
	dbOrms = make(map[string]orm.Ormer)
	err = dbInit()
	return
}

func GetOrm(name string) orm.Ormer {
	o, exist := dbOrms[name]
	if !exist {
		return nil
	}

	return o
}
