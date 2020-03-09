package pq

import (
	"fmt"

	gormbulk "github.com/t-tiger/gorm-bulk-insert"
)

// Migrate func creates db table
func (s *Storage) Migrate() {
	s.db.AutoMigrate(&RenewableResource{})
}

// Populate populates db table with passed records
func (s *Storage) Populate(r []*RenewableResource) error {
	var records []interface{}
	for _, v := range r {
		records = append(records, v)
	}
	// WARN: using gorm-bulk-insert until the following GORM issue gets closed:
	// https://github.com/jinzhu/gorm/issues/255
	err := gormbulk.BulkInsert(s.db, records, 1000)
	if err != nil {
		return fmt.Errorf("Error while inserting the data: %v", err)
	}

	return nil
}
