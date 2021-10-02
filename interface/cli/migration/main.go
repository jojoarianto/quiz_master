package main

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/jojoarianto/quiz_master/config"
	"github.com/jojoarianto/quiz_master/domain/model"
)

func main() {
	conf := config.NewConfig("sqlite3", "quiz_master.sqlite3")
	db, _ := conf.ConnectDB()

	DBMigrate(db)
}

// DBMigrate will create and migrate the tables
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&model.Question{})
	log.Println("Schema migration has been procceed")

	return db
}
