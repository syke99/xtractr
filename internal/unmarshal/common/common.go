package common

import (
	"github.com/syke99/xtractr/internal/pkg/resources/models"
	"strconv"
	"strings"
	"time"
)

func FormatTime(value string, format string) (time.Time, error) {
	var t time.Time
	var err error

	layout := ""

	if format == "" || format == "ISO8601" {
		if format == "ISO8601" {
			var year int
			var month time.Month
			var day int
			var er error

			tParts := strings.Split(value, "-")

			year, er = strconv.Atoi(tParts[0])
			if er != nil {
				return time.Time{}, er
			}

			m := 0
			m, er = strconv.Atoi(tParts[1])
			if er != nil {
				return time.Time{}, er
			}

			month = time.Month(m)

			day, er = strconv.Atoi(tParts[2])
			if er != nil {
				return time.Time{}, er
			}

			t = time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
		} else {
			layout = time.Layout
		}
	}

	if format != "" && format != "ISO8601" {
		if f, ok := models.TimeLayouts()[format]; ok {
			layout = f
		}

		t, err = time.Parse(layout, value)
		if err != nil {
			return time.Time{}, err
		}
	}

	return t, err
}
