package helper

import (
	"database/sql"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connection() *gorm.DB {

	var connString string = "host=localhost user=pranavrnayakn password=pass@1234 dbname=expense-monster port=5432 sslmode=disable "

	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{})

	if err != nil {
		log.Fatalln("Error connecting to database:", err)
	}
	return db
}

func HandleConnection() *sql.DB {
	db := Connection()

	sqlDB, err := db.DB()

	//Closing DB connection only on panic situation.
	defer func() {
		if r := recover(); r != nil {
			sqlDB.Close()
		}
	}()

	if err != nil {
		log.Fatalln("Error: on db connection", err)
	}

	if pingErr := sqlDB.Ping(); err != nil {
		log.Fatalln("Error: on db ping", pingErr)
	}

	return sqlDB
}
