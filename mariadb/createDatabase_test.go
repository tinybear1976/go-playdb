package mariadb

import (
	"testing"
)

func TestGenerateCreateDatabaseSQL(t *testing.T) {
	testCases := []string{
		"dbName1",
		"dbName2",
	}

	for _, tc := range testCases {
		sql := GenerateCreateDatabaseSQL(tc)
		t.Logf("%s, Generated SQL: %v", tc, sql)
	}
}
