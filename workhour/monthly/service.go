package monthly

import (
	"time"
)

func getMonthly(userID int, dateFrom time.Time, dateTo time.Time) (b []byte, err error) {
	result, err := findMonthly(userID, dateFrom, dateTo)
	if err != nil {
		return nil, err
	}

	return marshal(result)
}
