package pool

import (
	s "github.com/wangjia184/beats/metricbeat/schema"
	c "github.com/wangjia184/beats/metricbeat/schema/mapstriface"
)

var (
	schema = s.Schema{
		"name": c.Str("pool"),
		"connections": s.Object{
			"accepted": c.Int("accepted conn"),
			"queued":   c.Int("listen queue"),
		},
		"processes": s.Object{
			"idle":   c.Int("idle processes"),
			"active": c.Int("active processes"),
		},
		"slow_requests": c.Int("slow requests"),
	}
)
