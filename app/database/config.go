package database

import "fmt"

var (
	dbUsername = "postgres"
	dbPassword = "changeme"
	dbHost     = "postgres_container"
	dbTable    = "postgres"
	dbPort     = "5432"
	pgConnStr  = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUsername, dbTable, dbPassword)
	//pgConnStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",dbUsername,dbPassword,dbHost,dbPort,dbTable)
)
