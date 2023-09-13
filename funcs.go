package jalaali

import (
	"regexp"
	"strconv"
	"strings"
	"time"

	ptime "github.com/yaa110/go-persian-calendar"
)

// TehranTz get tehran time zone
func TehranTz() *time.Location {
	return time.FixedZone("Asia/Tehran", 12600)
}

// KabulTz get kabul time zone
func KabulTz() *time.Location {
	l, err := time.LoadLocation("Asia/Kabul")
	if err != nil {
		return nil
	}
	return l
}

// Parse parse jalaali date from string
//
// this function return nil if invalid date passed
func Parse(str string) *time.Time {
	return ParseForLocale(str, time.Now().Location())
}

// ParseForLocale parse jalaali date from string for locale
//
// this function return nil if invalid date passed
func ParseForLocale(str string, loc *time.Location) *time.Time {
	isNumber := func(v string) bool { return regexp.MustCompile(`^[0-9]+$`).MatchString(v) }
	pattern := regexp.MustCompile(`[^\d]+`)
	_parts := strings.Split(pattern.ReplaceAllString(str, "-"), "-")
	parts := make([]int, 6)
	for i := 0; i < 6; i++ {
		item := "0"
		if len(_parts) > i && isNumber(_parts[i]) {
			item = _parts[i]
		}
		t, _ := strconv.Atoi(item)
		parts[i] = t
	}
	jDate := ptime.Date(parts[0], ptime.Month(parts[1]), parts[2], parts[3], parts[4], parts[5], 0, loc)
	if jDate.Year() == parts[0] && int(jDate.Month()) == parts[1] && jDate.Day() == parts[2] {
		res := jDate.Time()
		return &res
	}
	return nil
}

// New create new jalaali date from time
//
// this function use current time zone
func New(t time.Time) Jalaali {
	var jd Jalaali
	jd.SetTime(t)
	return jd
}

// New create new jalaali date from time
//
// this function use tehran time zone
func NewTehran(t time.Time) JalaaliTehran {
	var jd JalaaliTehran
	jd.SetTime(t)
	return jd
}

// New create new jalaali date from time
//
// this function use kabul time zone
func NewKabul(t time.Time) JalaaliKabul {
	var jd JalaaliKabul
	jd.SetTime(t)
	return jd
}
