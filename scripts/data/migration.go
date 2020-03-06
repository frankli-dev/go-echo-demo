package data

import (
	"log"

	"github.com/ivan-marquez/golang-demo/pkg/storage/pq"
	"github.com/jinzhu/gorm"
	gormbulk "github.com/t-tiger/gorm-bulk-insert"
)

// TODO: add comment
func Migrate(db *gorm.DB) {
	log.Println("creating table…")
	db.AutoMigrate(&pq.RenewableResource{})
	log.Println("table successfully created.")

	log.Println("requesting data to be inserted…")
	res, err := DataSetRequest()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("data retrieved successfully.")

	log.Println("parsing retrieved data…")
	// Process Json data
	values, err := ParseResponse(res.Body())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("inserting data…")
	var records []interface{}
	for _, v := range values {
		records = append(records, v)
	}
	// WARN: using gorm-bulk-insert until the following GORM issue gets closed:
	// https://github.com/jinzhu/gorm/issues/255
	gormbulk.BulkInsert(db, records, 1000)
	log.Println("data inserted successfully.")
}
