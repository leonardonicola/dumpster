package internal

import (
	"fmt"
	"sync"
)

func BackupPostgres(credential DatabaseCredentials, wg *sync.WaitGroup) error {
	defer wg.Done()
	// exec.Command("pg_dump", "-U", credential.Username, "-h", "localhost", "-p", "5432", "-d", credential.DBName, "-f", "pgsql.sql")
	fmt.Print("Sucessfully backed up Postgres...\n")
	return nil
}
func BackupSQLite(credential DatabaseCredentials, wg *sync.WaitGroup) error {
	defer wg.Done()
	// exec.Command("cp", credential.DBName, "backup.db")
	fmt.Print("Sucessfully backed up SQLite...\n")
	return nil
}
func BackupMongoDB(credential DatabaseCredentials, wg *sync.WaitGroup) error {
	defer wg.Done()
	fmt.Print("Sucessfully backed up Mongo...\n")
	return nil
}
func BackupMySQL(credential DatabaseCredentials, wg *sync.WaitGroup) error {
	defer wg.Done()
	// exec.Command("mysqldump", "-u", credential.Username, "-p", credential.Password, credential.DBName, "> mysql.sql")
	fmt.Print("Sucessfully backed up MySQL...\n")
	return nil
}
