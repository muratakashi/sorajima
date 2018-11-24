package monthly

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq" // PostgreSQL Driver
)

// WorkHourStruct 勤務時間情報
type WorkHourStruct struct {
	ID        int64  `json:"id"`
	UserID    int64  `json:"user_id"`
	WorkDay   string `json:"work_day"`
	Hello     string `json:"hello"`
	Goodbye   string `json:"goodbye"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

const (
	dbUser     = "postgres"
	dbPassword = "postgres"
	dbName     = "ohara"
)

func findMonthly(userID int, dateFrom time.Time, dateTo time.Time) (workhours []WorkHourStruct, err error) {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", dbUser, dbPassword, dbName)
	db, err := sql.Open("postgres", connStr)
	defer db.Close()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query(sqlFindMonthly, dateFrom, dateTo, userID)
	defer rows.Close()

	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, err
	default:
	}

	for rows.Next() {
		var element WorkHourStruct
		rows.Scan(&element.ID, &element.UserID, &element.WorkDay, &element.Hello, &element.Goodbye, &element.CreatedAt, &element.UpdatedAt)

		workhours = append(workhours, element)
	}

	return workhours, nil
}
