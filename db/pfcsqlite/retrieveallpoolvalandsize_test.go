package pfcsqlite

import (
	"testing"

	"github.com/picfight/pfcdata/testutil"
)

func TestRetrieveAllPoolValAndSize(t *testing.T) {
	testutil.BindCurrentTestSetup(t)
	db := InitTestDB(DBPathForTest())
	testEmptyDBRetrieveAllPoolValAndSize(db)
}

func testEmptyDBRetrieveAllPoolValAndSize(db *DB) {
	result, err := db.RetrieveAllPoolValAndSize()
	if err != nil {
		testutil.ReportTestFailed(
			"RetrieveAllPoolValAndSize() failed: default result expected:",
			err)
	}
	checkChartsDataIsDefault("RetrieveAllPoolValAndSize", result)
}
