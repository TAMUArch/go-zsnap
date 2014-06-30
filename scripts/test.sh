#!/bin/bash

DAYS="20140620 20140621 20140622 20140623"
HOURS="1340 1430 1520 1750"
WEEKS="22 23 24"
ZSNAP_BIN="/opt/zfsnap/bin/zfsnap"
BEG_SNAPSHOT_COUNT=23

function build_test {
  zfs create rpool/test
  for day in $DAYS
  do
    zfs snapshot rpool/test@${day}

    for hour in $HOURS
    do
      zfs snapshot rpool/test@${day}-${hour}
    done
  done

  for week in $WEEKS
  do
    zfs snapshot rpool/test@201406-week${week}
  done
}

function count_snapshots {
  count=$(zfs list -H -t snapshot | grep rpool/test | wc -l)

  echo $count
  return $count
}

build_test
