package jalaali

import (
	"encoding/json"
	"time"
)

// JalaaliTehran jalaali date with tehran locale
type JalaaliTehran struct {
	Jalaali
}

// Parse parse from jalaali string date
func (tTime *JalaaliTehran) Parse(str string) {
	tTime.time = ParseForLocale(str, TehranTz())
}

// SetTime set jalaali date from time
func (tTime *JalaaliTehran) SetTime(t time.Time) {
	t = t.In(TehranTz())
	tTime.time = &t
}

// UnmarshalJSON parse jalaali from bytes
func (tTime *JalaaliTehran) UnmarshalJSON(data []byte) error {
	var strData string
	if err := json.Unmarshal(data, &strData); err != nil {
		return err
	}
	tTime.Parse(strData)
	return nil
}
