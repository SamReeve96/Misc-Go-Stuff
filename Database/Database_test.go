package database

import (
	"testing"
)

func TestMain(t *testing.T) {
	var testQueries = []struct {
		fieldName      string
		tableName      string
		searchCriteria string
		expResult      string
	}{
		{"firstname", "user", "id = 1", "Sam"},
		{"lastname", "user", "id = 1", "Reeve"},
		{"initials", "user", "id = 1", "SR"},
		{"firstname", "user", "id = 1", "Sam"},
		{"title", "document", "created like '2017-07-11%' Limit 1", "Go lang notes"},
	}
	//setup
	db := connectToDB()

	for _, c := range testQueries {
		Result := queryDb(db, c.fieldName, c.tableName, c.searchCriteria)

		if Result != c.expResult {
			t.Log("Should be: " + c.expResult + " got: " + Result)
			t.Fail()
		}
	}
}
