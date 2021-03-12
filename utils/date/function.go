package date

import (
	"strings"
	"time"
)

const (
	TimeLayoutYear       = "2006"
	TimeLayoutYearSimple = "06"
	TimeLayoutMonth      = "01"
	TimeLayoutDay        = "02"
	TimeLayoutHour       = "15"
	TimeLayoutMinute     = "04"
	TimeLayoutSecond     = "05"
)

func TimeLayout(format string) string {
	var layouts = map[string]string{
		"YYYY": TimeLayoutYear,
		"MM":   TimeLayoutMonth,
		"DD":   TimeLayoutDay,
		"hh":   TimeLayoutHour,
		"ii":   TimeLayoutMinute,
		"ss":   TimeLayoutSecond,
	}
	for k, v := range layouts {
		format = strings.Replace(format, k, v, -1)
	}
	return format
}

func TimeFormat(time time.Time, format ...string) string {
	layout := "YYYY-MM-DD hh:ii:ss"
	if len(format) > 0 {
		layout = format[0]
	}
	return time.Format(TimeLayout(layout))
}
