package zsnap

import (
	"testing"
)

func GenerateVolume() Volume {
	volName := "rpool/test"
	snapNames := []string{
		"rpool/test@20140620",
		"rpool/test@20140621",
		"rpool/test@20140622",
		"rpool/test@20140620-1340",
		"rpool/test@20140620-1440",
		"rpool/test@201406-week22",
		"rpool/test@201406-week23",
		"rpool/test@201406-week24"}

	var snaps []Snapshot
	for _, snap := range snapNames {
		snaps = append(snaps, Snapshot{Name: snap})
	}

	return Volume{Name: volName, Snapshots: snaps}
}

func TestParseSnapshots(t *testing.T) {
	volume := GenerateVolume()
	volume.ParseSnapshots()

	for _, snapshot := range volume.Snapshots {
		if snapshot.Type == "" {
			t.Errorf("Expected snapshot %s type other than nil", snapshot.Name)
		}
	}
}

func TestSnapshotsOfType(t *testing.T) {
	volume := GenerateVolume()
	volume.ParseSnapshots()

	hourlySnapshots := volume.SnapshotsOfType("hourly")
	if len(hourlySnapshots) != 2 {
		t.Errorf("Expected 2 hourly snapshots but returned %s", len(hourlySnapshots))
	}

	dailySnapshots := volume.SnapshotsOfType("daily")
	if len(dailySnapshots) != 3 {
		t.Errorf("Expected 3 daily snapshots but returned %i", len(dailySnapshots))
	}

	weeklySnapshots := volume.SnapshotsOfType("weekly")
	if len(weeklySnapshots) != 3 {
		t.Errorf("Expected 3 weekly snapshots but returned %i", len(weeklySnapshots))
	}

	noSnapshots := volume.SnapshotsOfType("nonreal")
	if len(noSnapshots) != 0 {
		t.Errorf("Expected 0 fake snapshots but returned %i", len(noSnapshots))
	}
}
