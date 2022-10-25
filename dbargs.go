package main

import (
	"github.com/radiochild/utils"
)

const dbHostEnvName = "REPORT_DB_HOST"
const defaultDBHost = "localhost"

const databaseEnvName = "REPORT_DB_DATABASE"
const defaultDatabase = "Datawarehouse"

const userEnvName = "REPORT_DB_USER"
const defaultUsername = "youser"

const xwordEnvName = "REPORT_DB_PASSWORD"
const defaultPassword = "top-secret"

const dbPortEnvName = "REPORT_DB_PORT"
const defaultDBPort = 5432

// --------------------------------------------------------------------------------
// DBParams (from Env or Vault)
// --------------------------------------------------------------------------------
type DBParams struct {
	Host     string
	Database string
	User     string
	Port     int
	Password string
}

func NewDBParamsFromEnv() *DBParams {
	host := utils.StringFromEnv(dbHostEnvName, defaultDBHost)
	database := utils.StringFromEnv(databaseEnvName, defaultDatabase)
	user := utils.StringFromEnv(userEnvName, utils.Username(defaultUsername))
	password := utils.StringFromEnv(xwordEnvName, defaultPassword)
	port := utils.IntFromEnv(dbPortEnvName, defaultDBPort)

	dbp := DBParams{
		Host:     host,
		Database: database,
		User:     user,
		Password: password,
		Port:     port,
	}
	return &dbp
}

// func NewDBParamsFromVault() *DBParams {
// 	host := "host_from_vault"
// 	database := "database_from_vault"
// 	user := "user_from_vault"
// 	password := "password_from_vault"
// 	portVal := defaultDBPort
//
// 	dbp := DBParams{
// 		Host:     host,
// 		Database: database,
// 		User:     user,
// 		Password: password,
// 		Port:     portVal,
// 	}
// 	return &dbp
// }
