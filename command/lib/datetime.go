package lib

import (
	"fmt"
	"time"
)

func FormatDate(dt time.Time) string {
	return dt.Format("2006-01-02")
}

func FormatDateTime(dt time.Time) string {
	return dt.Format("2006-01-02 15:01:02")
}

func ParseDateTime(dt string) time.Time {
	dT, err := time.Parse("2006-01-02 15:01:02", dt)
	if err != nil {
		fmt.Println(err.Error())
	}
	return dT
}

func ParseDate(d string) time.Time {
	dT, err := time.Parse("2006-01-02", d)
	if err != nil {
		fmt.Println(err.Error())
	}
	return dT
}
