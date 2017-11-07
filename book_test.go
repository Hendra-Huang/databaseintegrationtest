// +build integration

package databaseintegrationtest

import (
	"context"
	"testing"

	"github.com/Hendra-Huang/databaseintegrationtest/database"
	"github.com/Hendra-Huang/databaseintegrationtest/testingutil"
)

func TestGetBooks(t *testing.T) {
	t.Parallel()
	testDb, schemaName, cleanup := database.CreateTestDatabase(t)
	defer cleanup()

	loadTestData(t, testDb, schemaName, "book")
	getBooksQuery = addSchemaPrefix(schemaName, getBooksQuery) // override the query and add schema to the query

	books, err := GetBooks(context.Background(), testDb) // will execute getBooksQuery
	testingutil.Ok(t, err)
	testingutil.Equals(t, 2, len(books))
}
