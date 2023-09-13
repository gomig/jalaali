package jalaali_test

import (
	"testing"
	"time"

	"github.com/gomig/jalaali"
)

func TestIsNil(t *testing.T) {
	jDate := jalaali.Jalaali{}
	if !jDate.IsNil() {
		t.Error("failed nil test!")
	}
}

func TestSetTime(t *testing.T) {
	jDate := jalaali.Jalaali{}
	jDate.SetTime(time.Now())
	if jDate.IsNil() {
		t.Error("failed setTime!")
	}
}

func TestParse(t *testing.T) {
	jDate := jalaali.JalaaliTehran{}
	jDate.Parse("1400-12-29 15:40")
	if jDate.IsNil() || jDate.UTC().Format("15:04") != "12:10" {
		t.Error("failed Parse!")
	}
}

func TestTime(t *testing.T) {
	jDate := jalaali.JalaaliTehran{}
	jDate.Parse("1400-01-02 15:02:04")
	if jDate.Time() == nil {
		t.Error("failed get time!")
	}
}

func TestJalaali(t *testing.T) {
	jDate := jalaali.JalaaliTehran{}
	jDate.Parse("1400-01-02 15:02:04")
	if jDate.JTime() == nil {
		t.Error("failed get jTime!")
	}
}

func TestUTC(t *testing.T) {
	jDate := jalaali.JalaaliTehran{}
	jDate.Parse("1400-01-02 15:02:04")
	if jDate.UTC() == nil {
		t.Error("failed get time utc!")
	}
}

func TestFormat(t *testing.T) {
	jDate := jalaali.JalaaliTehran{}
	jDate.Parse("1400-01-02 15:02:04")
	if jDate.Format("2006-01-02 15:04:05") != "1400-01-02 15:02:04" {
		t.Error("failed format jalaali time!")
	}
}

func TestString(t *testing.T) {
	jDate := jalaali.JalaaliTehran{}
	jDate.Parse("1400-01-02 15:02:04")
	if jDate.String() != "1400-01-02 15:02:04 +0330 Asia/Tehran" {
		t.Error("failed get jalaali string!")
	}
}

func TestMarshalJSON(t *testing.T) {
	jDate := jalaali.JalaaliTehran{}
	jDate.Parse("1400-01-02 15:02:04")
	tJson, err := jDate.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}

	if string(tJson) != `"2021-03-22T15:02:04+03:30"` {
		t.Error("failed marshal as json!")
	}
}

func TestUnmarshalJSON(t *testing.T) {
	jDate := jalaali.JalaaliTehran{}
	err := jDate.UnmarshalJSON([]byte(`"1400-01-02 15:02:04"`))
	if err != nil {
		t.Fatal(err)
	}

	if jDate.String() != "1400-01-02 15:02:04 +0330 Asia/Tehran" {
		t.Log(jDate.String())
		t.Error("failed unmarshal json!")
	}
}
