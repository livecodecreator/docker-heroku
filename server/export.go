package server

import (
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

// GetStatus is
func GetStatus() Status {

	return status
}

// UpdateStatus is
func UpdateStatus() {

	status.PostTime = time.Now()

	h, err := host.Info()
	if err == nil {
		status.BootTime = h.BootTime
	}

	c, err := cpu.Percent(time.Second, false)
	if err == nil {
		sum := 0.0
		cnt := 0.0
		for _, v := range c {
			sum += v
			cnt++
		}
		if cnt != 0 {
			status.CPU = sum / cnt
		}
	}

	parts, err := disk.Partitions(false)
	if err == nil {
		sum := 0.0
		cnt := 0.0
		for _, v := range parts {
			u, err := disk.Usage(v.Mountpoint)
			if err == nil {
				sum += u.UsedPercent
				cnt++
			}
		}
		if cnt != 0 {
			status.Disk = sum / cnt
		}
	}

	m, err := mem.VirtualMemory()
	if err == nil {
		status.Memory = m.UsedPercent
	}
}
