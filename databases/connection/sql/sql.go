package sql

import (
	"fmt"
	"time"

	"hotel/commonHelpers"
	"hotel/utils"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

// DB Object for connection
var DB *gorm.DB

// NewDB inital connection db
func NewDB(host, database, user, password, port string) {

	dsn := fmt.Sprintf(
		"sqlserver://%s:%s@%s:%s?database=%s",
		user,
		password,
		host,
		port,
		database,
	)
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	if err != nil {
		utils.SaveErrorToApplicationInsight("failed connect to DB", "unexpected_error", err.Error(), "", 0)
		panic(err)
	}

	isDatabaseOnDebugMode := commonHelpers.GetConfigurationBoolValue("database.sql_server.debug")
	if isDatabaseOnDebugMode {
		db = db.Debug()
	}
	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetConnMaxLifetime(60 * time.Minute)
	DB = db
}
