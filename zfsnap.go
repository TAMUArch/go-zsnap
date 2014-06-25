package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jarosser06/go-zsnap/zsnap"
)

func main() {
	var action = flag.String("action", "create", "whether to remove or create a new snapshot")
	var keep = flag.Int("keep", 5, "how many snapshots to keep")
	var snapshotType = flag.String("type", "hourly", "what type of snapshot to create")
	flag.Parse()

	for _, volume := range flag.Args() {
		vol := zsnap.Volume{Name: volume}
		err := vol.GetSnapshots()
		if err != nil {
			os.Exit(1)
		}

		if *action == "create" {
			err = vol.Snapshot(*snapshotType)

			if err != nil {
				fmt.Printf("Error creating snapshot: %s\n", err)
				os.Exit(1)
			}
		} else if *action == "remove" {
			vol.CleanupSnapshots(*snapshotType, *keep)
		} else {
			fmt.Println("Passed unsupported action: %s\n", *action)
			os.Exit(1)
		}
	}
}
