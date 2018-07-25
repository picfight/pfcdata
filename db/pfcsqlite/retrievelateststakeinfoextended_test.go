package pfcsqlite

import (
	"testing"

	"github.com/picfight/pfcdata/testutil"
)

func TestRetrieveLatestStakeInfoExtended(t *testing.T) {
	testutil.BindCurrentTestSetup(t)
	db := InitTestDB(DBPathForTest())
	testEmptyDBRetrieveLatestStakeInfoExtended(db)
}

func testEmptyDBRetrieveLatestStakeInfoExtended(db *DB) {
	result, err := db.RetrieveLatestStakeInfoExtended()
	if err == nil {
		testutil.ReportTestFailed(
			"RetrieveLatestStakeInfoExtended() failed: error expected")
	}
	if result != nil {
		testutil.ReportTestFailed(
			"RetrieveLatestStakeInfoExtended() failed:"+
				" nil  expected, %v provided",
			result)
	}
}
