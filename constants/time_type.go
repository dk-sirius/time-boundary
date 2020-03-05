package constants

type TimeType uint8

const (
	TIME_TYPE_UNKNOWN TimeType = iota
	TIME_TYPE__MIN              // 分钟
	TIME_TYPE__HOUR             // 小时
	TIME_TYPE__DAY              // 日
	TIME_TYPE__WEEK             // 周
	TIME_TYPE__MON              // 月
	TIME_TYPE__SEASON           // 季度
	TIME_TYPE__YEAR             // 年
)
