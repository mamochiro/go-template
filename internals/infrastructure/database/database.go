package database

import (
	"fmt"
	"go-boilerplate/internals/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// this file should be a connect Database only

// DB is struct of this file.
type DB struct {
	Connection *gorm.DB
	env        config.Configuration
}

// MigrateDB ...
func (db *DB) MigrateDB() {
	fmt.Println("Start migrate db")
	// if err := db.Connection.Debug().AutoMigrate(
	// 	&entity.Accounts{},
	// ).Error; err != nil {
	// 	panic(err)
	// }
}

// Seeder ...
func (db *DB) Seeder() {
	fmt.Println("Start seeder")
}

// NewServerBase is start connection database.
func NewServerBase(config config.Configuration) *DB {
	// Uncomment code below when want to connect database

	// ADD CONDITION BY ENV.DBDRIVER IF YOU WANT TO CHANGE DATABASE.
	// db, err := gorm.Open(config.DbDriver, "host="+config.DbHost+" port="+config.DbPort+" user="+config.DbUser+" dbname="+config.DbName+" password="+config.DbPassword+" sslmode=disable")
	// if err != nil {
	// 	panic(err)
	// }

	// if config.NodeEnv != "production" {
	// 	db.LogMode(true)
	// }

	// return &DB{
	// 	Connection: db,
	// 	Config:     config,
	// }

	return &DB{}
}
