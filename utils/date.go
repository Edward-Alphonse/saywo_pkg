package utils

import (
	"fmt"
	"time"
)

// GetLastWeekDate 获取上周开始和结束日期
func GetLastWeekDate() (startDate, endDate string) {
	now := time.Now()
	currentWeekDay := time.Now().Weekday()
	switch currentWeekDay {
	case time.Sunday:
		return now.Format("2006-01-02"), now.AddDate(0, 0, 7).Format("2006-01-02")
	case time.Monday:
		return now.AddDate(0, 0, -1).Format("2006-01-02"), now.AddDate(0, 0, 6).Format("2006-01-02")
	case time.Tuesday:
		return now.AddDate(0, 0, -2).Format("2006-01-02"), now.AddDate(0, 0, 5).Format("2006-01-02")
	case time.Wednesday:
		return now.AddDate(0, 0, -3).Format("2006-01-02"), now.AddDate(0, 0, 4).Format("2006-01-02")
	case time.Thursday:
		return now.AddDate(0, 0, -4).Format("2006-01-02"), now.AddDate(0, 0, 3).Format("2006-01-02")
	case time.Friday:
		return now.AddDate(0, 0, -5).Format("2006-01-02"), now.AddDate(0, 0, 2).Format("2006-01-02")
	case time.Saturday:
		return now.AddDate(0, 0, -6).Format("2006-01-02"), now.AddDate(0, 0, 1).Format("2006-01-02")
	default:
		return now.AddDate(0, 0, -6).Format("2006-01-02"), now.AddDate(0, 0, 1).Format("2006-01-02")
	}
}

// GetLastMonthDate 获取当前月开始和结束日期
func GetLastMonthDate() (string, string) {
	now := time.Now()
	startDate := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location()).Format("2006-01-02")
	endDate := time.Date(now.Year(), now.Month()+1, 1, 0, 0, 0, 0, now.Location()).Format("2006-01-02")
	return startDate, endDate
}

// GetStartDateAndEndDate 获取开始\结束时间
// time：day、week、month
func GetStartDateAndEndDate(timeType string) (s, e string) {
	var startDate, endDate string
	if timeType == "day" {
		startDate, endDate = time.Now().AddDate(0, 0, -1).Format("2006-01-02"), time.Now().Format("2006-01-02")
	} else if timeType == "week" {
		startDate, endDate = GetLastWeekDate()
	} else {
		startDate, endDate = GetLastMonthDate()
	}
	return startDate, endDate
}

// DateStart 指定日期开始时间
func DateStart(st time.Time) time.Time {
	return time.Date(st.Year(), st.Month(), st.Day(),
		0, 0, 0, 0, time.Now().Location())
}

// DateEnd 指定日期结束时间
func DateEnd(st time.Time) time.Time {
	return time.Date(st.Year(), st.Month(), st.Day(),
		23, 59, 59, 0, time.Now().Location())
}

// GetWeekDay 获得当前周的初始和结束日期
func GetWeekDay() (string, string) {
	now := time.Now()
	offset := int(time.Monday - now.Weekday())
	//周日做特殊判断 因为time.Monday = 0
	if offset > 0 {
		offset = -6
	}

	lastoffset := int(time.Saturday - now.Weekday())
	//周日做特殊判断 因为time.Monday = 0
	if lastoffset == 6 {
		lastoffset = -1
	}

	firstOfWeek := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	lastOfWeeK := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, lastoffset+1)
	f := firstOfWeek.Unix()
	l := lastOfWeeK.Unix()
	return time.Unix(f, 0).Format("2006-01-02") + " 00:00:00", time.Unix(l, 0).Format("2006-01-02") + " 23:59:59"
}

// GetWeekStartDay 获得当前周的初始日期(yyyyMMdd)
func GetWeekStartDay(t *time.Time) time.Time {
	offset := int(time.Monday - t.Weekday())
	//周日做特殊判断 因为time.Monday = 0
	if offset > 0 {
		offset = -6
	}

	firstOfWeek := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	f := firstOfWeek.Unix()
	return time.Unix(f, 0)
}

// GetLastWeekMondayMidnight 获取上周一凌晨的时间
func GetLastWeekMondayMidnight() time.Time {
	// 获取当前日期时间
	now := time.Now()

	// 获取当前日期所在的星期几（Sunday = 0, Monday = 1, ..., Saturday = 6）
	weekday := int(now.Weekday())

	// 计算当前日期与上周一之间的天数差
	daysSinceLastMonday := (weekday + 6) % 7

	// 计算上一周的周一日期
	lastWeekMonday := now.AddDate(0, 0, -daysSinceLastMonday-7)

	// 设置时间为凌晨 00:00:00
	lastWeekMondayMidnight := time.Date(lastWeekMonday.Year(), lastWeekMonday.Month(), lastWeekMonday.Day(), 0, 0, 0, 0, lastWeekMonday.Location())

	return lastWeekMondayMidnight
}

func GetTodayLeftTime() int {
	// 获取当前时间
	now := time.Now()

	// 获取当天的最后一秒时间
	endOfDay := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location())

	// 计算当前时间到当天结束的剩余秒数
	remainingSeconds := int(endOfDay.Sub(now).Seconds())

	return remainingSeconds
}

func GetTodayDateStr() string {
	// 获取当前时间
	now := time.Now()

	// 获取当天的年月日
	year := now.Year()
	month := now.Month()
	day := now.Day()

	// 构造年月日字符串
	dateString := fmt.Sprintf("%d-%02d-%02d", year, month, day)
	return dateString
}

func IsSameDay(date1, date2 time.Time) bool {
	return date1.Year() == date2.Year() && date1.Month() == date2.Month() && date1.Day() == date2.Day()
}

func DaysFromNow(targetTime time.Time) int {
	currentTime := time.Now()
	duration := currentTime.Sub(targetTime)
	days := int(duration.Hours() / 24)
	return days
}

func GetWeeksFromTimeUntilNextTime(date time.Time, nextDate time.Time) int {
	// 计算时间间隔
	duration := nextDate.Sub(date)
	// 将时间间隔转换为天数
	days := int(duration.Hours() / 24)
	// 计算周数
	weeks := days / 7
	return weeks
}
