package pfcsqlite

import (
	"testing"

	"github.com/picfight/pfcdata/testutil"
)

func TestRetrieveBlockSummaryByHash(t *testing.T) {
	testutil.BindCurrentTestSetup(t)
	db := InitTestDB(DBPathForTest())
	testEmptyDBRetrieveBlockSummaryByHash(db)
}

func testEmptyDBRetrieveBlockSummaryByHash(db *DB) {
	summary, err := db.RetrieveBlockSummaryByHash("")
	if err == nil {
		testutil.ReportTestFailed(
			"RetrieveBlockSummaryByHash() failed: error expected")
	}
	if summary != nil {
		testutil.ReportTestFailed(
			"RetrieveBlockSummaryByHash() failed: "+
				"nil expected, "+
				"%v returned", summary)
	}
}
