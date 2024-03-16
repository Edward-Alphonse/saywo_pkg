package utils

import (
	"fmt"
	"time"
)

// GetStartAndEndDateOfMonth 获取指定年月的开始时间和结束时间 juexin
func GetStartAndEndDateOfMonth(year, month int) (startDate, endDate time.Time) {
	dateStr := fmt.Sprintf("%d-%d-1", year, month)
	t, _ := time.Parse("2006-1-2", dateStr)

	currentYear, currentMonth, _ := t.Date()
	loc, _ := time.LoadLocation("Local") //获取时区

	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, loc)
	lastDate := firstOfMonth.AddDate(0, 1, -1)
	lastDateStr := fmt.Sprintf("%d-%d-%d 23:59:59", lastDate.Year(), lastDate.Month(), lastDate.Day())
	a, _ := time.Parse("2006-1-2 15:04:05", lastDateStr)
	return firstOfMonth, a
}
