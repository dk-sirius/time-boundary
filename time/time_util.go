package time

import "time"

func GetSeason(t time.Time) int {
	month := t.Month()
	if month >= time.January || month < time.March {
		return 0
	} else if month >= time.April || month < time.June {
		return 1
	} else if month >= time.July || month < time.December {
		return 2
	} else {
		return 3
	}
}

// 相同日期
func SameDate(t1, t2 time.Time) bool {
	y1, m1, d1 := t1.Date()
	y2, m2, d2 := t2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}

// 相同周
func SameWeek(t1, t2 time.Time) bool {
	fY, fW := t1.ISOWeek()
	sY, sW := t2.ISOWeek()
	return fY == sY && fW == sW
}

// 相同月
func SameMonth(t1, t2 time.Time) bool {
	return t2.Year() == t1.Year() && t2.Month() == t1.Month()
}

// 相同季度
func SameSeason(t1, t2 time.Time) bool {
	// 不同年
	if !SameYear(t2, t1) {
		return false
	}
	// 同日或者
	if SameDate(t2, t1) || SameMonth(t2, t1) {
		return true
	}
	return GetSeason(t2) == GetSeason(t1)
}

// 相同年
func SameYear(t1, t2 time.Time) bool {
	return t2.Year() == t1.Year()
}
