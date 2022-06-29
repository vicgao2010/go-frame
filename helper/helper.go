package helper

import (
	"strconv"
	"time"
)


func SetDefaultInt(value, def int) int {
	if value == 0 {
		return def
	}
	return value
}

func SetDefaultBool(value, def bool) bool {
	if value == false {
		return def
	}
	return value
}

func SetDefaultString(value, def string) string {
	if value == "" {
		return def
	}
	return value
}

func SetDefaultDuration(value, def string) time.Duration {
	if value == "" {
		duration, _ := time.ParseDuration(def)
		return duration
	}
	duration, _ := time.ParseDuration(value)
	return duration
}

func GetTimeStamp() int64 {
	timestamp := time.Now().Unix()
	return timestamp
}

func GetDateTime(timestamp int64) string {
	datetime := time.Unix(timestamp, 0).Format("2006-01-02 15:04:05")
	return datetime
}

func GetFormatDateTime(timestamp int64, fmt string) string {
	datetime := time.Unix(timestamp, 0).Format(fmt)
	return datetime
}

func ParseTimeStamp(date string, fmt string) int64 {
	stamp, _ := time.ParseInLocation(fmt, date, time.Local)
	return stamp.Unix()
}

func ToInt(v interface{}) int {
	switch v := v.(type) {
	case nil:
		return 0
	case int:
		return v
	case float64:
		return int(v)
	case []byte:
		vInt, _ := strconv.ParseInt(string(v), 10, 0)
		return int(vInt)
	case string:
		vInt, _ := strconv.Atoi(v)
		return vInt
	default:
		panic("to int error.")
	}
}

func ToInt64(v interface{}) int64 {
	switch v := v.(type) {
	case nil:
		return 0
	case int64:
		return v
	case float64:
		return int64(v)
	case []byte:
		vInt, _ := strconv.ParseInt(string(v), 10, 64)
		return vInt
	case string:
		vInt, _ := strconv.Atoi(v)
		return int64(vInt)
	default:
		panic("to int64 error.")
	}
}

func ToFloat64(v interface{}) float64 {
	switch v := v.(type) {
	case nil:
		return 0
	case float64:
		return v
	case string:
		vF, _ := strconv.ParseFloat(v, 64)
		return vF
	case []byte:
		vF, _ := strconv.ParseFloat(string(v), 64)
		return vF
	default:
		panic("to float64 error.")
	}
}

func ToString(v interface{}) string {
	switch v := v.(type) {
	case []byte:
		return string(v)
	case string:
		return v
	case int:
		return strconv.Itoa(v)
	case int64:
		return strconv.FormatInt(v, 10)
	case float64:
		return strconv.FormatFloat(v, 'E', -1, 64)
	case nil:
		return ""
	default:
		panic("to string error.")
	}
}
func ToBool(v interface{}) bool {
	switch v := v.(type) {
	case int64:
		return v != 0
	case []byte:
		vb, _ := strconv.ParseBool(string(v))
		return vb
	case nil:
		return false
	default:
		panic("to bool error.")
	}
}