package utils

import (
	"time"
)

type Date struct {
	timestamp uint64

	Year          uint32
	Month         uint32
	Day           uint32
	Zodiac        string
	Constellation string
}

func NewDate(timestamp uint64) *Date {
	year, month, day := getYMDFromTimestamp(timestamp)
	return &Date{
		Year:          year,
		Month:         month,
		Day:           day,
		Zodiac:        getZodiac(year),
		Constellation: getConstellation(month, day),
	}
}

func getYMDFromTimestamp(timestamp uint64) (year, month, day uint32) {
	date := time.Unix(int64(timestamp), 0)
	return uint32(date.Year()), uint32(date.Month()), uint32(date.Day())
}

//func getTimeFromStrDate(date string) (year, month, day uint32) {
//	const shortForm = "2006-01-02"
//	d, err := time.Parse(shortForm, date)
//	if err != nil {
//		fmt.Println("出生日期解析错误！", err)
//		return 0, 0, 0
//	}
//	year = uint32(d.Year())
//	month = uint32(d.Month())
//	day = uint32(d.Day())
//	return
//}

func getZodiac(year uint32) (zodiac string) {
	if year <= 0 {
		zodiac = "-1"
	}
	start := uint32(1901)
	x := (start - year) % 12
	if x < 0 {
		x += 12
	}

	switch x {
	case 1:
		zodiac = "鼠"
	case 0:
		zodiac = "牛"
	case 11:
		zodiac = "虎"
	case 10:
		zodiac = "兔"
	case 9:
		zodiac = "龙"
	case 8:
		zodiac = "蛇"
	case 7:
		zodiac = "马"
	case 6:
		zodiac = "羊"
	case 5:
		zodiac = "猴"
	case 4:
		zodiac = "鸡"
	case 3:
		zodiac = "狗"
	case 2:
		zodiac = "猪"
	}
	return
}

func GetAge(year uint32) (age uint32) {
	if year <= 0 {
		age = 0
	}
	nowyear := uint32(time.Now().Year())
	age = nowyear - year
	return
}

func getConstellation(month, day uint32) (star string) {
	switch {
	case month <= 0, month >= 13, day <= 0, day >= 32:
		star = "-1"
	case month == 1 && day >= 20, month == 2 && day <= 18:
		star = "水瓶座"
	case month == 2 && day >= 19, month == 3 && day <= 20:
		star = "双鱼座"
	case month == 3 && day >= 21, month == 4 && day <= 19:
		star = "白羊座"
	case month == 4 && day >= 20, month == 5 && day <= 20:
		star = "金牛座"
	case month == 5 && day >= 21, month == 6 && day <= 21:
		star = "双子座"
	case month == 6 && day >= 22, month == 7 && day <= 22:
		star = "巨蟹座"
	case month == 7 && day >= 23, month == 8 && day <= 22:
		star = "狮子座"
	case month == 8 && day >= 23, month == 9 && day <= 22:
		star = "处女座"
	case month == 9 && day >= 23, month == 10 && day <= 22:
		star = "天秤座"
	case month == 10 && day >= 23, month == 11 && day <= 21:
		star = "天蝎座"
	case month == 11 && day >= 22, month == 12 && day <= 21:
		star = "射手座"
	case month == 12 && day >= 22, month == 1 && day <= 19:
		star = "魔蝎座"
	}
	return
}

// GetMondayDateOfWeek 获取当周第一天的字符串
func GetMondayDateOfWeek() string {
	now := time.Now()
	zeroTime := GetZeroTime(now)
	// 获取当前时间所在地点的周一作为一周的开始
	offset := int(time.Monday - zeroTime.Weekday())
	startOfWeek := zeroTime.AddDate(0, 0, offset)
	// 格式化输出日期
	return startOfWeek.Format("2006-01-02")
}

// GetNextMondayTs 距离下周一0点的时间
func GetNextMondayTs() time.Duration {
	now := time.Now()
	zeroTime := GetZeroTime(now)

	offset := int(time.Monday-zeroTime.Weekday()) + 7

	nextMonday := zeroTime.AddDate(0, 0, offset)

	return nextMonday.Sub(now)
}

func GetZeroTime(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
}

func GetNextDayTs() time.Duration {
	now := time.Now()
	zeroTime := GetZeroTime(now)

	return zeroTime.Sub(now)
}
