package datastore

import (
	// "database/sql"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/boil"
	// "gorm.io/driver/postgres"
	// "gorm.io/gorm"
	// "gorm.io/gorm/logger"
)

func NewDB() *sql.DB {
	// if os.Getenv("GO_ENV") == "dev" {
	// 	err := godotenv.Load()
	// 	// err := godotenv.Load("../.env")
	// 	if err != nil {
	// 		log.Fatalln(err)
	// 	}
	// }

	// url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("POSTGRES_USER"),
	// 	os.Getenv("POSTGRES_PW"), os.Getenv("POSTGRES_HOST"),
	// 	os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_DB"))

	// newLogger := logger.New(
	// 	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	// 	logger.Config{
	// 		SlowThreshold:             time.Second, // Slow SQL threshold
	// 		LogLevel:                  logger.Info, // Log level
	// 		IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
	// 		ParameterizedQueries:      true,        // Don't include params in the SQL log
	// 		Colorful:                  true,        // Disable color
	// 	},
	// )

	fmt.Println("////////////////")
	db, err := sql.Open("postgres", "host=localhost port=25432  dbname=aaaa user=aaaa password=aaaa  sslmode=disable")
	if err != nil {
		fmt.Println("open database failure", err)
		log.Fatalln(err)
	}
	fmt.Println("Connceted", db)
	boil.SetDB(db)
	rows, err := db.Query("SELECT * FROM tasks")
	if err != nil {
		fmt.Println("ccc sadf", err)

	}
	fmt.Println("rrrrrrrrrrrrrrrrr", rows)
	return db
}

func CloseDB(db *sql.DB) {
	// sqlDB, _ := db.DB()
	// if err := sqlDB.Close(); err != nil {
	// 	log.Fatalln(err)
	// }
}
