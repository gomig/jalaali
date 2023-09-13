package jalaali

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	ptime "github.com/yaa110/go-persian-calendar"
)

// Jalaali data structure
type Jalaali struct {
	time *time.Time
}

// IsNil check if date is nil
func (jTime Jalaali) IsNil() bool {
	return jTime.time == nil
}

// SetTime set jalaali date from time
func (jTime *Jalaali) SetTime(t time.Time) {
	jTime.time = &t
}

// Parse parse from jalaali string date
func (jTime *Jalaali) Parse(str string) {
	jTime.time = Parse(str)
}

// Time get normal time
func (jTime Jalaali) Time() *time.Time {
	return jTime.time
}

// JTime get jalaali time
func (jTime Jalaali) JTime() *ptime.Time {
	if jTime.time != nil {
		r := ptime.New(*jTime.time)
		return &r
	}
	return nil
}

// UTC get normal time in UTC
func (jTime Jalaali) UTC() *time.Time {
	if jTime.time != nil {
		res := jTime.time.UTC()
		return &res
	}
	return nil
}

// Format format jalaali date using standard go time format
func (jTime Jalaali) Format(format string) string {
	if jTime.time != nil {
		jd := ptime.New(*jTime.time)
		r := strings.NewReplacer(
			"2006", "yyyy",
			"15", "HH",
			"06", "yy",
			"01", "MM",
			"1", "M",
			"January", "MMM",
			"Jan", "MMM",
			"02", "dd",
			"_2", fmt.Sprintf("%2d", jd.Day()),
			"2", "d",
			"Monday", "E",
			"Mon", "e",
			"03", "hh",
			"3", "h",
			"04", "mm",
			"4", "m",
			"05", "ss",
			"5", "s",
			"PM", "A",
			"pm", "a",
			"MST", "z",
			"Z0700", jd.ZoneOffset("Z0700"),
			"Z07:00", jd.ZoneOffset("Z07:00"),
			"-0700", jd.ZoneOffset("-0700"),
			"-07:00", jd.ZoneOffset("-07:00"),
			"-07", jd.ZoneOffset("-07"),
		)
		return jd.Format(r.Replace(format))
	}
	return ""
}

func (jTime Jalaali) String() string {
	return jTime.Format("2006-01-02 15:04:05 -0700 MST")
}

// MarshalJSON get value for json
func (jTime Jalaali) MarshalJSON() ([]byte, error) {
	return json.Marshal(jTime.time)
}

// UnmarshalJSON parse jalaali from bytes
func (jTime *Jalaali) UnmarshalJSON(data []byte) error {
	var strData string
	if err := json.Unmarshal(data, &strData); err != nil {
		return err
	}
	jTime.Parse(strData)
	return nil
}
