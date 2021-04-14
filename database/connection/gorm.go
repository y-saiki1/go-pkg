package connection

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	LIMIT = 30
)

var db *gorm.DB = nil

// ConnectionPoolの設定
// Source/Replicaの設定
// https://gorm.io/ja_JP/docs/dbresolver.html
func OpenDB() error {
	if db != nil {
		return nil
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
	// replica := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", os.Getenv("READ_DB_USER"), os.Getenv("READ_DB_PASS"), os.Getenv("READ_DB_HOST"), os.Getenv("READ_DB_NAME"))

	isDebugMode, err := strconv.ParseBool(os.Getenv("GORM_DEBUG_MODE"))
	if err != nil {
		return err
	}

	var conf *gorm.Config
	if isDebugMode {
		logConf := logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold: time.Second,
				LogLevel:      logger.Info,
				Colorful:      true,
			},
		)
		conf = &gorm.Config{Logger: logConf}
	} else {
		conf = &gorm.Config{}
	}

	conn, err := gorm.Open(mysql.Open(dsn), conf)
	if err != nil {
		return err
	}
	// resolverConfig := dbresolver.Config{Replicas: []gorm.Dialector{mysql.Open(replica)}}
	// db.Use(
	// 	dbresolver.Register(resolverConfig).
	// 		SetConnMaxIdleTime(time.Hour).
	// 		SetConnMaxLifetime(24 * time.Hour).
	// 		SetMaxIdleConns(100).
	// 		SetMaxOpenConns(200),
	// )
	db = conn
	return nil
}

func GetDB() (*gorm.DB, error) {
	if db == nil {
		return nil, errors.New("not opened connection to db")
	}

	return db, nil
}
