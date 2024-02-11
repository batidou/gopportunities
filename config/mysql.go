package config

import (
	"github.com/batidou/gopportunities/schemas"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func InitializeMySQL() (*gorm.DB, error) {
	logger := GetLogger("mysql")
	dbPath := "./db/main.db"

	// check if the database file exists
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Errorf("Database file does not exist: %v", err)
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			logger.Errorf("Error creating database directory: %v", err)
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			logger.Errorf("Error creating database file: %v", err)
			return nil, err
		}
		defer file.Close()
		logger.Info("Database file created")
	}

	var (
		datetimePrecision = 2
		dns               = "root:root@tcp(127.0.0.1:3306)/gopportunities?charset=utf8mb4&parseTime=True&loc=Local"
	)

	// Initialize the MySQL database
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dns,
		DefaultStringSize:         256,                // add default size for string fields, by default, will use db type `longtext` for fields without size, not a primary key, no index defined and don't have default values
		DisableDatetimePrecision:  true,               // disable datetime precision support, which not supported before MySQL 5.6
		DefaultDatetimePrecision:  &datetimePrecision, // default datetime precision
		DontSupportRenameIndex:    true,               // drop & create index when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,               // use change when rename column, rename not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false,              // smart configure based on used version
	}), &gorm.Config{})

	if err != nil {
		logger.Errorf("Error initializing MySQL: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("Error migrating schema: %v", err)
		return nil, err
	}

	logger.Info("MySQL initialized")
	return db, nil
}
