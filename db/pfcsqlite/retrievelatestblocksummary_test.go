package pfcsqlite

/*
This file contains checks performed by tests
*/
import (
	"testing"

	"github.com/picfight/pfcdata/testutil"
)

func TestRetrieveLatestBlockSummary(t *testing.T) {
	testutil.BindCurrentTestSetup(t)
	db := InitTestDB(DBPathForTest())
	testEmptyDBRetrieveLatestBlockSummary(db)
}

func testEmptyDBRetrieveLatestBlockSummary(db *DB) {
	summary, err := db.RetrieveLatestBlockSummary()
	// expected "sql: no rows in result set"
	if err == nil {
		testutil.ReportTestFailed(
			"RetrieveLatestBlockSummary() failed: error expected")
	}
	if summary != nil {
		testutil.ReportTestFailed(
			"RetrieveLatestBlockSummary() failed: "+
				"nil expected, "+
				"%v returned", summary)
	}
}
