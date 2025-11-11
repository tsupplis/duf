//go:build solaris || illumos
// +build solaris illumos

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/sys/unix"
)

func (m *Mount) Stat() unix.Statvfs_t {
	return m.Metadata.(unix.Statvfs_t)
}

func mounts() ([]Mount, []string, error) {
	var ret []Mount
	var warnings []string

	// Read /etc/mnttab on illumos/Solaris
	file, err := os.Open("/etc/mnttab")
	if err != nil {
		return nil, nil, fmt.Errorf("opening /etc/mnttab: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		// mnttab format: special mount_point fstype options time
		if len(fields) < 4 {
			continue
		}

		device := fields[0]
		mountPoint := fields[1]
		fsType := fields[2]
		opts := fields[3]

		if len(device) == 0 {
			continue
		}

		var stat unix.Statvfs_t
		err := unix.Statvfs(mountPoint, &stat)
		if err != nil {
			if err != os.ErrPermission {
				warnings = append(warnings, fmt.Sprintf("%s: %s", mountPoint, err))
			}
			continue
		}

		d := Mount{
			Device:     device,
			Mountpoint: mountPoint,
			Fstype:     fsType,
			Type:       fsType,
			Opts:       opts,
			Metadata:   stat,
			Total:      stat.Blocks * stat.Frsize,
			Free:       stat.Bavail * stat.Frsize,
			Used:       (stat.Blocks - stat.Bfree) * stat.Frsize,
			Inodes:     stat.Files,
			InodesFree: stat.Ffree,
			InodesUsed: stat.Files - stat.Ffree,
			Blocks:     stat.Blocks,
			BlockSize:  stat.Frsize,
		}
		d.DeviceType = deviceType(d)

		ret = append(ret, d)
	}

	if err := scanner.Err(); err != nil {
		return nil, warnings, fmt.Errorf("reading /etc/mnttab: %w", err)
	}

	return ret, warnings, nil
}
