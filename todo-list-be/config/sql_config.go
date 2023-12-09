package config

import (
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewSQL(log *logrus.Logger) (*gorm.DB){
	dbString := os.Getenv("DB_STRING")
	if dbString == "" {
		log.Fatalln("failed to connect to database, sql connection string is empty or not loaded")
	}
	
	dsn := dbString
	prefix := "://"
	index := strings.Index(dsn, prefix)
	dsn = dsn[(index + len(prefix)):]


  	sql, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		TranslateError: true,
		// Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalln("failed to connect to database: ", err)
	}
	return sql
}