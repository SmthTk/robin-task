package database

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

// InitDB initializes the database connection using Viper configuration.
func InitDB() (*gorm.DB, error) {
	// Read database configuration from Viper
	host := viper.GetString("database.host")
	port := viper.GetString("database.port")
	dbName := viper.GetString("database.database")
	username := viper.GetString("database.username")
	password := viper.GetString("database.password")
	maxIdleConns := viper.GetInt("database.max_idle_conns")
	maxOpenConns := viper.GetInt("database.max_open_conns")

	// Build the DSN (Data Source Name) string for MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", username, password, host, port, dbName)

	// Connect to the database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Set connection pool options
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(maxIdleConns) // Set maximum idle connections
	sqlDB.SetMaxOpenConns(maxOpenConns) // Set maximum open connections

	log.Printf("Successfully connected to the database: %s", dbName)
	return db, nil
}
