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

package routes

import (
	"encoding/json"

	"github.com/wangjia184/beats/metricbeat/mb"

	"github.com/pkg/errors"

	s "github.com/wangjia184/beats/libbeat/common/schema"
	c "github.com/wangjia184/beats/libbeat/common/schema/mapstriface"
)

var (
	moduleSchema = s.Schema{
		"server": s.Object{
			"id":   c.Str("server_id"),
			"time": c.Str("now"),
		},
	}
	routesSchema = s.Schema{
		"total": c.Int("num_routes"),
	}
)

func eventMapping(r mb.ReporterV2, content []byte) error {
	var event mb.Event
	var inInterface map[string]interface{}

	err := json.Unmarshal(content, &inInterface)
	if err != nil {
		err = errors.Wrap(err, "failure parsing Nats routes API response")
		r.Error(err)
		return err
	}
	event.MetricSetFields, err = routesSchema.Apply(inInterface)
	if err != nil {
		err = errors.Wrap(err, "failure applying routes schema")
		r.Error(err)
		return err
	}

	event.ModuleFields, err = moduleSchema.Apply(inInterface)
	if err != nil {
		err = errors.Wrap(err, "failure applying module schema")
		r.Error(err)
		return err
	}
	r.Event(event)
	return nil
}
