package mariadb

import (
	"fmt"
	"strings"
)

func GenerateCreateDatabaseSQL(dbName string) string {
	if LowercaseDatabaseName {
		dbName = strings.ToLower(dbName)
	}
	return fmt.Sprintf(`CREATE DATABASE IF NOT EXISTS %s
  CHARACTER SET utf8mb4
  COLLATE utf8mb4_general_ci;`, dbName)
}
