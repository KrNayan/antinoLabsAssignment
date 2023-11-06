package mysql

import (
	log "antinolabsassignment/pkg/common/utilities/logger"
	"testing"
)

func TestNewSQLDbInstance(t *testing.T) {
	logger := log.New()
	sqlDb, err := NewSQLDbInstance(logger)
	if err != nil {
		t.Fatal(err)
	}

	logger.Infof("sql connection created successfully ", sqlDb)
}
