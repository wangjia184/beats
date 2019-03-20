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

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package windows

import (
	"github.com/wangjia184/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "windows", asset.ModuleFieldsPri, AssetWindows); err != nil {
		panic(err)
	}
}

// AssetWindows returns asset data.
// This is the base64 encoded gzipped contents of ../metricbeat/module/windows.
func AssetWindows() string {
	return "eJysVM1uIjkQvvMUpVxySdBm97QcVsqGyYiRMoomM+JIG3dBl+K2Pa4yhLcfuX9I00AgUXyInLL7+ytT1/CMmxGsyeZuzQMAITE4gotpXbkYAOTIOpAXcnYE/w0AAB5cHg3CwgWYbj/lwgWZaWcXtBzBQhnGAUBAg4pxBEs1AFgQmpxHFcg1WFVilzwt2fh0Objom8oB/l2gGmxLNEdRnXpNwhhWpHFbP0R0lKxeWYORgXZWFFkGKRBYlETuZtGS8bDz/W4O7erb6EqmfKfcKn7GzdqF/hm+qNKnzhXTf75/u9c3t7534w1nad1CtPQ7IkzGlZfKWu1jCBMBYlBQKC7ALarDUumCLF4yfP01GYOyeSrv4TYYlafXPA4aTn8/YnmKmlf6fXZ/vto7R1pO7I3azD4ssXkZX1ZoBe6cMajFhfdrboRUstpOtG162wKLCjJLat9h4JwUE270FUxfUXXBO2aaG4SVMhEZVEDIbqO4Ugnp7GoPNfvfOcmuIBsTq7nBPO0flI3KZFfVQ8ueNixYZqctyye7VVqiMjXy+W7vnBWykezykN1HFbk6qre13x/R2qb4lBLe7p339b4OIv2P+akk8IXSaM4/MY17F17ptyMPpCBOwyIFgyG4AIkWpFCy/cEF9C4I70GuC7T1eyK7BHHAjdsULTGsyRiYY4W9RIuBdH/m7mF2NERrkLnbMvDBrShPbWpL1+xR04J058sT4fojk9o4uzw2EG7++vfvD+TdvorDeStmp0kJ5smYTmYfJ+MT6qMXKnFY9ttx1MPChVLJCPIYVBLbOybro8zaSyUZQ4za2bxPcP6EvuRGJTTNwRzI7mAPB38CAAD//+JTaOY="
}
