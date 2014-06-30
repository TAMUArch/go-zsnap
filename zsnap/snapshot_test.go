package zsnap

import (
	"strings"
	"testing"
	"time"
)

func YearTest(actual time.Time, expectedYear int, t *testing.T) {
	actualYear := actual.Year()
	if actualYear != expectedYear {
		t.Errorf("Snapshot creation year expected %d instead of %v", expectedYear, actualYear)
	}
}

func MonthTest(actual time.Time, expectedMonth int, t *testing.T) {
	actualMonth := actual.Month()
	if int(actualMonth) != expectedMonth {
		t.Errorf("Snapshot creation month expected %d instead of %v", expectedMonth, actualMonth)
	}
}

func DayTest(actual time.Time, expectedDay int, t *testing.T) {
	actualDay := actual.Day()
	if actualDay != expectedDay {
		t.Errorf("Expected snapshot day to be %v instead of %v", expectedDay, actualDay)
	}
}

func SnapshotTypeTest(actual string, expected string, t *testing.T) {
	if actual != expected {
		t.Errorf("Snapshot type expected %v instead of %v", expected, actual)
	}
}

func TestParse_weekly(t *testing.T) {
	snapName := "tank/test_snapshot@201406-week22"

	snapshot := Snapshot{Name: snapName}
	err := snapshot.Parse()

	if err != nil {
		t.Errorf("Parsing snapshot %v returned with %v", snapName, err)
	} else {
		if snapshot.Week != 22 {
			t.Errorf("Snapshot week expected 22 instead of %v", snapshot.Week)
		}

		SnapshotTypeTest(snapshot.Type, "weekly", t)
		YearTest(snapshot.Creation, 2014, t)
		MonthTest(snapshot.Creation, 6, t)
	}
}

func TestParse_daily(t *testing.T) {
	snapName := "tank/test_snapshot@20140620"
	snapshot := Snapshot{Name: snapName}
	err := snapshot.Parse()

	if err != nil {
		t.Errorf("Parsing snapshot %v returned with %v", snapName, err)
	} else {
		if snapshot.Week != 0 {
			t.Errorf("Snapshot Week expected 0 but returned %v", snapshot.Week)
		}

		SnapshotTypeTest(snapshot.Type, "daily", t)
		YearTest(snapshot.Creation, 2014, t)
		MonthTest(snapshot.Creation, 6, t)
		DayTest(snapshot.Creation, 20, t)
	}
}

func TestParse_hourly(t *testing.T) {
	snapName := "tank/test_snapshot@20140620-1430"
	snapshot := Snapshot{Name: snapName}
	err := snapshot.Parse()

	if err != nil {
		t.Errorf("Parsing snapshot %v returned with %v", snapName, err)
	} else {
		SnapshotTypeTest(snapshot.Type, "hourly", t)
		YearTest(snapshot.Creation, 2014, t)
		MonthTest(snapshot.Creation, 6, t)
		DayTest(snapshot.Creation, 20, t)

		snapshotHour := snapshot.Creation.Hour()
		if snapshotHour != 14 {
			t.Errorf("Snapshot hour expected 14 instead of %v", snapshotHour)
		}

		snapshotMinute := snapshot.Creation.Minute()
		if snapshotMinute != 30 {
			t.Errorf("Snapshot minute expected 30 instead of %v", snapshotMinute)
		}
	}
}

func TestParse_invalid(t *testing.T) {
	snapName := "tank/test_snapshot@invalid-name"
	snapshot := Snapshot{Name: snapName}

	err := snapshot.Parse()

	if err == nil {
		t.Errorf("Expected non nil error")
	}
}

func TestTimeStamp_weekly(t *testing.T) {
	stamp, err := TimeStamp("weekly")
	if err != nil {
		t.Errorf("Timestamp returned with an error %v", err)
	} else {
		if !strings.Contains(stamp, "-week") {
			t.Errorf("Weekly timestamp returned incorrectly: %v", stamp)
		}
	}
}

func TestTimeStamp_daily(t *testing.T) {
	_, err := TimeStamp("daily")

	if err != nil {
		t.Errorf("Daily timestamp returned with error %v", err)
	}
}

func TestTimeStamp_hourly(t *testing.T) {
	_, err := TimeStamp("hourly")

	if err != nil {
		t.Errorf("Hourly timestamp returned with error %v", err)
	}
}

func TestTimeStamp_badInput(t *testing.T) {
	_, err := TimeStamp("should not work")

	if err == nil {
		t.Errorf("Expected timestamp to return error but returned nil")
	}
}
