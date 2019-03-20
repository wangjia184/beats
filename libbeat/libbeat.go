package main

import (
	"os"

	"github.com/wangjia184/beats/libbeat/beat"
	"github.com/wangjia184/beats/libbeat/mock"
)

func main() {
	if err := beat.Run(mock.Name, mock.Version, mock.New); err != nil {
		os.Exit(1)
	}
}
