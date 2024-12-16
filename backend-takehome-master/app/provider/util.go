package provider

import (
	"fmt"
	"time"
)

func parseBytetoTime(timeByte []byte) time.Time {
	parsedTime, err := time.Parse("2006-01-02 15:04:05", string(timeByte))
	if err != nil {
		fmt.Println(fmt.Sprintf("Error parsing time: %s", string(timeByte)), err)
	}
	parsedTime = parsedTime.Truncate(24 * time.Hour)
	return parsedTime
}
