package pfcsqlite

import (
	"testing"

	"github.com/picfight/pfcdata/testutil"
)

func TestRetrieveBlockSummary(t *testing.T) {
	testutil.BindCurrentTestSetup(t)
	db := InitTestDB(DBPathForTest())
	testEmptyDBRetrieveBlockSummary(db)
}

func testEmptyDBRetrieveBlockSummary(db *DB) {
	summary, err := db.RetrieveBlockSummary(0)
	if err == nil {
		testutil.ReportTestFailed(
			"RetrieveBlockSummary() failed: error expected")
	}
	if summary != nil {
		testutil.ReportTestFailed(
			"RetrieveBlockSummary() failed: "+
				"nil expected, "+
				"%v returned", summary)
	}
}
