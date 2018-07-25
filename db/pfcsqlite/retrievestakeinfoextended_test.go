package pfcsqlite

import (
	"testing"

	"github.com/picfight/pfcdata/testutil"
)

func TestRetrieveStakeInfoExtended(t *testing.T) {
	testutil.BindCurrentTestSetup(t)
	db := InitTestDB(DBPathForTest())
	testEmptyDBRetrieveStakeInfoExtended(db)
}

func testEmptyDBRetrieveStakeInfoExtended(db *DB) {
	info, err := db.RetrieveStakeInfoExtended(0)
	if err == nil {
		testutil.ReportTestFailed(
			"RetrieveStakeInfoExtended() failed: error expected")
	}
	if info != nil {
		testutil.ReportTestFailed(
			"RetrieveStakeInfoExtended() failed:"+
				" nil expected, %v provided",
			info)
	}
}
