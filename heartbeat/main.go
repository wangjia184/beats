package main

import (
	"os"

	"github.com/wangjia184/beats/libbeat/beat"

	"github.com/wangjia184/beats/heartbeat/beater"

	// register default heartbeat monitors
	_ "github.com/wangjia184/beats/heartbeat/monitors/defaults"
)

func main() {
	err := beat.Run("heartbeat", "", beater.New)
	if err != nil {
		os.Exit(1)
	}
}
