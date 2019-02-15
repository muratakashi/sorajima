package entity

// WorkHourStruct 勤務時間情報
type WorkHourStruct struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	WorkDay   string `json:"work_day"`
	Hello     string `json:"hello"`
	Goodbye   string `json:"goodbye"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
