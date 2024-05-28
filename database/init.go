package database

import (
	"fmt"
	"log"

	"github.com/pandhu/hehemock/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var connDB *gorm.DB
var err error

func Init(conf *config.Configuration) *gorm.DB {
	if connDB != nil {
		return connDB
	}

	configDB := conf.Database
	dbConString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=UTC&multiStatements=true",
		configDB.User,
		configDB.Password,
		configDB.Host,
		configDB.Port,
		configDB.Name,
	)

	connDB, err := gorm.Open(mysql.Open(dbConString), &gorm.Config{})

	if err != nil {
		log.Printf("Error when connecting to mysql db, %s", err)
	}
	return connDB
}
