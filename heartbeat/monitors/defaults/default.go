package defaults

import (
	// register standard active monitors
	_ "github.com/wangjia184/beats/heartbeat/monitors/active/http"
	_ "github.com/wangjia184/beats/heartbeat/monitors/active/icmp"
	_ "github.com/wangjia184/beats/heartbeat/monitors/active/tcp"
)
