// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package cluster_stats

import (
	"github.com/pkg/errors"

	"github.com/wangjia184/beats/metricbeat/helper/elastic"
	"github.com/wangjia184/beats/metricbeat/mb"
	"github.com/wangjia184/beats/metricbeat/module/elasticsearch"
)

func init() {
	mb.Registry.MustAddMetricSet(elasticsearch.ModuleName, "cluster_stats", New,
		mb.WithHostParser(elasticsearch.HostParser),
		mb.WithNamespace("elasticsearch.cluster.stats"),
	)
}

const (
	clusterStatsPath = "/_cluster/stats"
)

// MetricSet defines all fields of the MetricSet
type MetricSet struct {
	*elasticsearch.MetricSet
}

// New create a new instance of the MetricSet
func New(base mb.BaseMetricSet) (mb.MetricSet, error) {
	ms, err := elasticsearch.NewMetricSet(base, clusterStatsPath)
	if err != nil {
		return nil, err
	}
	return &MetricSet{MetricSet: ms}, nil
}

// Fetch methods implements the data gathering and data conversion to the right format
func (m *MetricSet) Fetch(r mb.ReporterV2) {
	isMaster, err := elasticsearch.IsMaster(m.HTTP, m.HostData().SanitizedURI+clusterStatsPath)
	if err != nil {
		err := errors.Wrap(err, "error determining if connected Elasticsearch node is master")
		elastic.ReportAndLogError(err, r, m.Log)
		return
	}

	// Not master, no event sent
	if !isMaster {
		m.Log.Debug("trying to fetch cluster stats from a non-master node")
		return
	}

	content, err := m.HTTP.FetchContent()
	if err != nil {
		elastic.ReportAndLogError(err, r, m.Log)
		return
	}

	info, err := elasticsearch.GetInfo(m.HTTP, m.HostData().SanitizedURI+clusterStatsPath)
	if err != nil {
		r.Error(errors.Wrap(err, "failed to get info from Elasticsearch"))
		return
	}

	if m.MetricSet.XPack {
		err = eventMappingXPack(r, m, *info, content)
	} else {
		err = eventMapping(r, *info, content)
	}

	if err != nil {
		m.Log.Error(err)
	}
}
