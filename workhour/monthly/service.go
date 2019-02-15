package monthly

import (
	"time"

	"github.com/sorajima/db/entity"
)

func getMonthly(userID int, year int, month int) (b []byte, err error) {

	dateFrom := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Local)
	dateTo := time.Date(year, time.Month(month+1), 1, 23, 59, 59, 0, time.Local).AddDate(0, 1, -1)

	results, err := findMonthly(userID, dateFrom, dateTo)
	if err != nil {
		return nil, err
	}

	// １月分の情報を作成
	var monthlyWorkHours []entity.WorkHourStruct
	var date time.Time
	for i := 1; i < dateTo.Day(); i++ {
		date = time.Date(year, time.Month(month), i, 0, 0, 0, 0, time.Local)
		var workhour = entity.WorkHourStruct{
			ID:        -1,
			UserID:    userID,
			WorkDay:   date.Format("2006-01-02"),
			Hello:     "",
			Goodbye:   "",
			CreatedAt: "",
			UpdatedAt: "",
		}
		for _, result := range results {
			workday, _ := time.ParseInLocation("2006-01-02", result.WorkDay, time.Local)
			if date.Equal(workday) {
				workhour.Hello = result.Hello
				workhour.Goodbye = result.Goodbye
				workhour.CreatedAt = result.CreatedAt
				workhour.UpdatedAt = result.UpdatedAt
				break
			}
		}

		monthlyWorkHours = append(monthlyWorkHours, workhour)
	}

	return marshal(monthlyWorkHours)
}
