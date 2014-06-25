package zsnap

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	hourlyLayout = "2006012-1504"
	dailyLayout  = "2006012"
	weeklyLayout = "200601"
)

type Snapshot struct {
	Name     string
	Creation time.Time
	Type     string
	Week     int
}

// Parses a snapshot for the timestamp and snapshot type
func (s *Snapshot) Parse() error {
	tmpStr := strings.Split(s.Name, "@")[1]

	var err error
	if strings.Contains(tmpStr, "week") {
		//discarding the week number
		splitStr := strings.Split(tmpStr, "-week")
		dateStr := splitStr[0]
		s.Week, _ = strconv.Atoi(splitStr[1])
		s.Creation, err = time.Parse(weeklyLayout, dateStr)
		s.Type = "weekly"
		if err != nil {
			return err
		}
		// Assume daily if timestamp no greater than 8 chars
	} else if len(tmpStr) <= 8 {
		s.Creation, err = time.Parse(dailyLayout, tmpStr)
		s.Type = "daily"
		if err != nil {
			return err
		}
	} else {
		s.Creation, err = time.Parse(hourlyLayout, tmpStr)
		s.Type = "hourly"
		if err != nil {
			return err
		}
	}
	return nil
}

// returns a time stamp based on hourly, daily, or weekly
func TimeStamp(stampType string) (string, error) {
	t := time.Now()
	switch stampType {
	case "hourly":
		return t.Format(hourlyLayout), nil
	case "daily":
		return t.Format(dailyLayout), nil
	case "weekly":
		_, week := t.ISOWeek()
		weekStr := strconv.Itoa(week)
		return (t.Format(weeklyLayout) + "-week" + weekStr), nil
	}
	return "", errors.New(fmt.Sprintf("timestamp \"%s\" not supported", stampType))
}
