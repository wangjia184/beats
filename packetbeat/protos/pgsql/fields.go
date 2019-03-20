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

package pgsql

import (
	"github.com/wangjia184/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("packetbeat", "pgsql", asset.ModuleFieldsPri, AssetPgsql); err != nil {
		panic(err)
	}
}

// AssetPgsql returns asset data.
// This is the base64 encoded gzipped contents of protos/pgsql.
func AssetPgsql() string {
	return "eJzEkcFu6jAURPf5ihHrBx+QBRJCPAkJtRTYI0gmwapjB1+bKn9fhaTF0GbBqllF1z4zR9djvLNJUZdy1gnglddMMVpb8aXj9m01SoCckjlVe2VNimkCALcLY6mZqUJl4IXGo1DUuUwS9H/p9f4Y5lDxVtR+vqmZonQ21P0kJmKKzlm3z2zO76MHq92JkVMHoAUmEdEVamvKZKCiosihfK6lZyZDmcILnfLNU6FfUKxfWxF11NxfDjowWlPXudhsXjcPs/+z3Wz1MFvPXpbzn7ImVPtu/0Oi07ugZQF/IraL1WK+wznQNVAFJGQZRYqg/8GflHRvCiUQenh7hUyojnR3cbboXx+OPjjD/JeFto7OfvyZYdsd+30GAAD//yg57bo="
}
