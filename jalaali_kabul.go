package jalaali

import (
	"encoding/json"
	"time"
)

// JalaaliKabul jalaali date with kabul locale
type JalaaliKabul struct {
	Jalaali
}

// Parse parse from jalaali string date
func (kTime *JalaaliKabul) Parse(str string) {
	kTime.time = ParseForLocale(str, KabulTz())
}

// SetTime set jalaali date from time
func (kTime *JalaaliKabul) SetTime(t time.Time) {
	t = t.In(KabulTz())
	kTime.time = &t
}

// UnmarshalJSON parse jalaali from bytes
func (kTime *JalaaliKabul) UnmarshalJSON(data []byte) error {
	var strData string
	if err := json.Unmarshal(data, &strData); err != nil {
		return err
	}
	kTime.Parse(strData)
	return nil
}
