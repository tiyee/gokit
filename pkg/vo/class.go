package vo

type Class struct {
	ClassID      string `json:"class_id"`
	ClassName    string `json:"class_name"`
	TeacherName  string `json:"teacher_name"`
	EndDate      string `json:"end_date"`
	Duration     string `json:"duration"`
	StartTime    int64  `json:"start_time"`
	EndTime      int64  `json:"end_time"`
	TotalLessons int    `json:"total_lessons"`
	MTime        int64  `json:"mtime"`
	CTime        int64  `json:"ctime"`
	Price        int    `json:"price"`
	Tag          int8   `json:"tag"`
}
