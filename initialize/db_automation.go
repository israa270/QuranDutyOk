package initialize

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

//ClearTable
//clear database table data
//db(database object ) *gorm.DB, tableName(table Name) string, compareField(compare field) string, interval(between interval) string
//error
func ClearTable(db *gorm.DB, tableName string, compareField string, interval string) error {
	if db == nil {
		return errors.New("db cannot be empty")
	}
	duration, err := time.ParseDuration(interval)
	if err != nil {
		return err
	}
	if duration < 0 {
		return errors.New("parse duration < 0")
	}
	return db.Debug().Exec(fmt.Sprintf("DELETE FROM %s WHERE %s < ?", tableName, compareField), time.Now().Add(-duration)).Error
}
