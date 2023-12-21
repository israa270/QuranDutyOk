package utils

import (
	"fmt"
	"time"
)

// ConvertStringTOTimeUTC
func ConvertStringTOTimeUTC(timeString string) (*time.Time, error) {
	// st := strings.NewReader(timeString)

	// var t time.Time

	// err := json.NewDecoder(st).Decode(&t)
	// if err != nil {
	// 	return nil, err
	// }

	t, err := time.Parse("2006-01-02T15:04:05.000Z", timeString)
	fmt.Println("Error time parse: ", err)
	return &t, nil
}

// ConvertIntToTime
func ConvertIntToTime(timeStamp int64) time.Time {
	return time.Unix(timeStamp, 0)
}
