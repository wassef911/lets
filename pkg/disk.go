package pkg

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"syscall"
)

type DiskServiceInterface interface {
	ShowDiskSpace() ([]string, error)
	ShowFolderSize(path string) (string, error)
	ShowFolderSizeWithLimit(dir string, minSize float64) ([]string, error)
}

var _ DiskServiceInterface = &DiskService{}

type DiskService struct{}

func NewDisk() *DiskService {
	return &DiskService{}
}

// displays disk usage for all mounts (equivalent to `df -h`).
func (d *DiskService) ShowDiskSpace() ([]string, error) {
	var stat syscall.Statfs_t
	var diskUsage []string
	// mounted filesystems
	mounts, err := os.ReadFile("/proc/mounts")
	if err != nil {
		return nil, errors.New("failed to read mounts: " + err.Error())
	}
	diskUsage = append(diskUsage, "Filesystem      Size  Used  Avail  Use%  Mounted on")
	for _, line := range strings.Split(string(mounts), "\n") {
		fields := strings.Fields(line)
		if len(fields) < 2 {
			continue
		}
		mountPoint := fields[1]

		// get filesystem stats
		err := syscall.Statfs(mountPoint, &stat)
		if err != nil {
			continue
		}

		// calc disk usage
		total := stat.Blocks * uint64(stat.Bsize)
		avail := stat.Bavail * uint64(stat.Bsize)
		used := total - avail
		usePercent := (float64(used) / float64(total)) * 100

		totalHR := humanReadableSize(total)
		usedHR := humanReadableSize(used)
		availHR := humanReadableSize(avail)

		diskUsage = append(diskUsage, fmt.Sprintf("%-15s %5s %5s %5s %3.0f%%  %s", fields[0], totalHR, usedHR, availHR, usePercent, mountPoint))
	}

	return diskUsage, nil
}

// calculates the size of a directory (equivalent to `du -sh /var`).
func (d *DiskService) ShowFolderSize(path string) (string, error) {
	var totalSize int64

	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			totalSize += info.Size()
		}
		return nil
	})
	if err != nil {
		return "", errors.New("failed to calculate folder size: " + err.Error())
	}

	var format string = humanReadableSize(uint64(totalSize))
	return fmt.Sprintf("%s %s\n", format, path), nil
}

// finds files larger than a specified size in a directory (equivalent to `find /home -type f -size +100M`).
func (d *DiskService) ShowFolderSizeWithLimit(dir string, minSize float64) ([]string, error) {
	var largeFiles []string

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && float64(info.Size()) > minSize {
			largeFiles = append(largeFiles, path)
		}
		return nil
	})
	if err != nil {
		return nil, errors.New("failed to list large files: " + err.Error())
	}

	return largeFiles, nil
}

// humanReadableSize converts a size in bytes to a human-readable format.
func humanReadableSize(size uint64) string {
	const unit = 1024
	if size < unit {
		return fmt.Sprintf("%d B", size)
	}
	div, exp := uint64(unit), 0
	for n := size / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(size)/float64(div), "KMGTPE"[exp])
}
