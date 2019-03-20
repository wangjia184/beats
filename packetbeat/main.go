package main

import (
	"os"

	"github.com/wangjia184/beats/libbeat/beat"
	"github.com/wangjia184/beats/packetbeat/beater"

	// import protocol modules
	_ "github.com/wangjia184/beats/packetbeat/include"
)

var Name = "packetbeat"

// Setups and Runs Packetbeat
func main() {
	if err := beat.Run(Name, "", beater.New); err != nil {
		os.Exit(1)
	}
}
