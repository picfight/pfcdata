package pfcsqlite

import (
	"testing"

	"github.com/picfight/pfcdata/testutil"
)

func TestGetStakeInfoHeight(t *testing.T) {
	testutil.BindCurrentTestSetup(t)
	db := InitTestDB(DBPathForTest())
	testEmptyDBGetStakeInfoHeight(db)
}

func testEmptyDBGetStakeInfoHeight(db *DB) {
	endHeight, err := db.GetStakeInfoHeight()
	if err != nil {
		testutil.ReportTestFailed(
			"GetStakeInfoHeight() failed: %v",
			err)
	}
	if endHeight != -1 {
		testutil.ReportTestFailed(
			"GetStakeInfoHeight() failed: endHeight=%v,"+
				" should be -1",
			endHeight)
	}
}
