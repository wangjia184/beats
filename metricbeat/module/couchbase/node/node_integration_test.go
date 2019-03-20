// +build integration

package node

import (
	"testing"

	mbtest "github.com/wangjia184/beats/metricbeat/mb/testing"
	"github.com/wangjia184/beats/metricbeat/module/couchbase"
)

func TestData(t *testing.T) {
	f := mbtest.NewEventsFetcher(t, getConfig())

	err := mbtest.WriteEvents(f, t)
	if err != nil {
		t.Fatal("write", err)
	}
}

func getConfig() map[string]interface{} {
	return map[string]interface{}{
		"module":     "couchbase",
		"metricsets": []string{"node"},
		"hosts":      []string{couchbase.GetEnvDSN()},
	}
}
