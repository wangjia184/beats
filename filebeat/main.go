package main

import (
	"os"
	"runtime"
	"strconv"
	"syscall"
	"time"

	"github.com/wangjia184/beats/filebeat/beater"
	"github.com/wangjia184/beats/libbeat/beat"
)

var Name = "filebeat"

// The basic model of execution:
// - prospector: finds files in paths/globs to harvest, starts harvesters
// - harvester: reads a file, sends events to the spooler
// - spooler: buffers events until ready to flush to the publisher
// - publisher: writes to the network, notifies registrar
// - registrar: records positions of files read
// Finally, prospector uses the registrar information, on restart, to
// determine where in each file to restart a harvester.

const PROCESS_MODE_BACKGROUND_BEGIN = 0x00100000

func main() {

	runtime.GOMAXPROCS(1)

	// https://bitsum.com/pl_io_priority.php
	dll, err := syscall.LoadLibrary("kernel32.dll")
	if err == nil {
		getCurrentProcess, err := syscall.GetProcAddress(dll, "GetCurrentProcess")
		if err == nil {
			handle, _, _ := syscall.Syscall(getCurrentProcess, 0, 0, 0, 0)
			setPriorityClass, err := syscall.GetProcAddress(dll, "SetPriorityClass")
			if err == nil {
				syscall.Syscall(setPriorityClass, 2, handle, PROCESS_MODE_BACKGROUND_BEGIN, 0)
			}

			setProcessAffinityMask, err := syscall.GetProcAddress(dll, "SetProcessAffinityMask")
			if err == nil {
				syscall.Syscall(setProcessAffinityMask, 2, handle, 1, 0)
			}
		}
	}

	processId, err := strconv.Atoi(os.Getenv("ParentProcess.ID"))

	if err == nil && processId > 0 {
		process, err := os.FindProcess(processId)
		if err == nil {
			go (func() {
				process.Wait()
				time.Sleep(15 * time.Second)
				os.Exit(0)
			})()
		}
	}

	if err := beat.Run(Name, "", beater.New); err != nil {
		os.Exit(1)
	}
}
