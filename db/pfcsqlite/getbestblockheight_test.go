package pfcsqlite

import (
	"testing"

	"github.com/picfight/pfcdata/testutil"
)

func TestGetBestBlockHeight(t *testing.T) {
	testutil.BindCurrentTestSetup(t)
	db := InitTestDB(DBPathForTest())
	testEmptyDBGetBestBlockHeight(db)
}

// Empty DB, should return -1
func testEmptyDBGetBestBlockHeight(db *DB) {
	h := db.GetBestBlockHeight()
	if h != -1 {
		testutil.ReportTestFailed(
			"db.GetBestBlockHeight() is %v,"+
				" should be -1",
			h)
	}
}
