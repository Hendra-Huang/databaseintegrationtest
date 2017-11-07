package databaseintegrationtest

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"testing"

	"database/sql"

	"github.com/Hendra-Huang/databaseintegrationtest/testingutil" // helper function for testing
)

// list of regexp pattern for adding schema to the query
var schemaPrefixRegexps = [...]*regexp.Regexp{
	regexp.MustCompile(`(?i)(^CREATE SEQUENCE\s)(["\w]+)(.*)`),
	regexp.MustCompile(`(?i)(^CREATE TABLE\s)(["\w]+)(\s.+)`),
	regexp.MustCompile(`(?i)(^ALTER TABLE\s)(["\w]+)(\s.+)`),
	regexp.MustCompile(`(?i)(^UPDATE\s)(["\w]+)(\s.+)`),
	regexp.MustCompile(`(?i)(^INSERT INTO\s)(["\w]+)(\s.+)`),
	regexp.MustCompile(`(?i)(^DELETE FROM\s)(["\w]+)(.*)`),
	regexp.MustCompile(`(?i)(.+\sFROM\s)(["\w]+)(.*)`),
	regexp.MustCompile(`(?i)(\sJOIN\s)(["\w]+)(.*)`),
}

// adding schema before the table name
func addSchemaPrefix(schemaName, query string) string {
	prefixedQuery := query
	for _, re := range schemaPrefixRegexps {
		prefixedQuery = re.ReplaceAllString(prefixedQuery, fmt.Sprintf("${1}%s.${2}${3}", schemaName))
	}
	return prefixedQuery
}

func loadTestData(t *testing.T, db *sql.DB, schemaName string, testDataNames ...string) {
	for _, testDataName := range testDataNames {
		file, err := os.Open(fmt.Sprintf("./testdata/%s.sql", testDataName))
		testingutil.Ok(t, err)
		reader := bufio.NewReader(file)
		var query string
		for {
			line, err := reader.ReadString('\n')
			if err == io.EOF {
				break
			}
			testingutil.Ok(t, err)
			line = line[:len(line)-1]
			if line == "" {
				query = addSchemaPrefix(schemaName, query)
				_, err := db.Exec(query)
				testingutil.Ok(t, err)
				query = ""
			}
			query += line
		}
		file.Close()
	}
}
