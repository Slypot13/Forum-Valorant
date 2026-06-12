package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// InitDB connecte l'application à MySQL.
func InitDB() *sql.DB {


	user := GetRequiredEnv("DB_USER")
	pwd := GetEnvWithDefault("DB_PWD", "")
	host := GetRequiredEnv("DB_HOST")
	port := GetRequiredEnv("DB_PORT")
	name := GetRequiredEnv("DB_NAME")

	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		user,
		pwd,
		host,
		port,
		name,
	)

	dbContext, err := sql.Open("mysql", connectionString)

	if err != nil {
		log.Fatal(err)
	}

	if err = dbContext.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("Connexion MySQL réussie")
	return dbContext
}