package lb

import (
	"testing"

	"github.com/wangjia184/beats/libbeat/common"
	"github.com/wangjia184/beats/libbeat/logp"
	"github.com/wangjia184/beats/libbeat/outputs"
)

var (
	testNoOpts     = outputs.Options{}
	testGuaranteed = outputs.Options{Guaranteed: true}

	testEvent = common.MapStr{
		"msg": "hello world",
	}
)

func enableLogging(selectors []string) {
	if testing.Verbose() {
		logp.LogInit(logp.LOG_DEBUG, "", false, true, selectors)
	}
}
