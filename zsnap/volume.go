package zsnap

import (
	"fmt"
	"os/exec"
	"strings"
)

type Volume struct {
	Name      string
	Snapshots []Snapshot
}

// Iterates through all snapshots in a volume and parses them
func (v *Volume) ParseSnapshots() {
	for i, snapshot := range v.Snapshots {
		snapshot.Parse()
		v.Snapshots[i] = snapshot
	}
}

// Returns snapshots of a specific type
func (v *Volume) SnapshotsOfType(snapshotType string) []Snapshot {
	var tmpSnapshots []Snapshot
	for _, snapshot := range v.Snapshots {
		if snapshot.Type == snapshotType {
			tmpSnapshots = append(tmpSnapshots, snapshot)
		}
	}

	return tmpSnapshots
}

// Creates a new snapshot based on the snapshot type (hourly, daily, weekly)
func (v *Volume) Snapshot(timeStampType string) error {
	timeStamp, err := TimeStamp(timeStampType)
	if err != nil {
		return err
	}

	snapshotName := strings.Join([]string{v.Name, timeStamp}, "@")
	cmd := exec.Command("/usr/sbin/zfs", "snapshot", snapshotName)
	fmt.Printf("Snapshotting -> %s as %s\n", v.Name, snapshotName)
	out, err := cmd.Output()

	if err != nil {
		fmt.Printf("Failed to create snapshot: %s\n", out)
		return err
	}

	return nil
}

// Lists all snapshots on a system and adds the relevant ones to the volume struct
// Gets snapshots and then calls parse snapshot
func (v *Volume) GetSnapshots() error {
	cmd := exec.Command("/usr/sbin/zfs", "list", "-H", "-o", "name", "-S", "name", "-t", "snapshot")
	cmdOut, err := cmd.Output()

	if err != nil {
		fmt.Printf("Error getting snapshots %s\n", cmdOut)
		return err
	}

	existingSnapshots := strings.Split(string(cmdOut), "\n")

	var volSnapshots []Snapshot
	for _, existingSnapshot := range existingSnapshots {
		if strings.Contains(existingSnapshot, v.Name) {
			volSnapshots = append(volSnapshots, Snapshot{Name: existingSnapshot})
		}
	}

	v.Snapshots = volSnapshots
	v.ParseSnapshots()

	return nil
}

/*
  get list of snapshots snapshots matching a specific type
  destroy snapshots past a certain number for instance:
  if you have 4 dailies but only want to hold onto 3 then
  delete the oldest daily
*/
func (v *Volume) CleanupSnapshots(snapshotType string, keep int) {
	snapshots := v.SnapshotsOfType(snapshotType)
	if len(snapshots) <= keep {
		fmt.Printf("Nothing to clean up for %s\n", v.Name)
	} else {

		for _, snapshot := range snapshots[keep:] {
			fmt.Printf("Destroying snapshot -> %s\n", snapshot.Name)
			err := DestroySnapshot(snapshot.Name)

			if err != nil {
				panic(err)
			}
		}
	}
}

//Destroys a snapshot
func DestroySnapshot(snapshot string) error {
	cmd := exec.Command("/usr/sbin/zfs", "destroy", snapshot)
	_, err := cmd.Output()

	if err != nil {
		return err
	}
	return nil
}
