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

package ceph

import (
	"github.com/wangjia184/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "ceph", asset.ModuleFieldsPri, AssetCeph); err != nil {
		panic(err)
	}
}

// AssetCeph returns asset data.
// This is the base64 encoded gzipped contents of ../metricbeat/module/ceph.
func AssetCeph() string {
	return "eJzEms9v27gSx+/5KwY5vQckxntXHx7Q1y3QYDdtsGlPi4VKkyOJa0okOFS8/u8XpGTZ1k/boV2fWtv5zsczo+HMSI+wxu0SOJr8DsBJp3AJ9x/R5Pd3AAKJW2mc1OUS/ncHAOA/gkKLSuEdAOXauoTrMpXZElKmyL9rUSEjXELG/HfQOVlmtIQ/7onU/QPc586Z+z/vAFKJStAyKD9CyQpsWfzLbY1XsboyzTsDRP71w//RD+C6dEyWBC5HKNBZyf2/mYMNWgTilhkUkFpdwMdPL58XjcAhxhGKqsihTYSkdfvhENYEmn+N6Bz7affqwhwCsTcmFVspXKy2DunoOzsupcus88EEmn992KlCUAWdBgc21J2vp9oWzC2hD7CDdNoxFRXwm1eMA1cRiqhs3wnF5WjdTMuRKZfHyLWe0vnZpt/QMqUScsxVw/5a43ajrTjPZV9rXah1p9zWppQskOfI17RAo3keKXbPzMAbWpK6nDNrdVWKxRtTFUYy3opD0D4NIHYoXo9CcAzSzc2e7eHcHM6zEzK2pz+Vm0NhuzwW4VRrfDGZEJalqeQLi0wkZ1WR8co0D1f7B7xRcLnVVZabyoFBC4Rcj+VOw7qx0uHNYYPVC2iDZ7VJjE8I5LHie+hCqQ2d57yr8dReOgWokGQU47gIp2tkjJ04mAzKqliNVOGWQa/+Qu5inaA9ikb+JBTLnNSDIMSZQpGkSjM3kt8GLcey++m5uH2EHaXAzDKB4ipB24nPBK1luE7QWooTgtai/LygtbjjQTPZQjDHbl4xTQbe7kTvarJFGABujsaOB4QxuJDlN4dz++FgDMw3/TfnqtrBYAQrNDwYWjpM/JuRLs2XLHTWePi9aQSuq94Fdan111xvCHK9gYKVWzAZAbMIsmygdDr703t8kZu9JkA03+9pEou0UsOle6W1QtalmjH+RF4UeqKHFktk9ipWvfCM6apINIlYp0QnGbyyz4S5cW9HUplrw/jU1DVQm57fXyahZPkzoJ6+TEJZLJgxKBKT3Zrs90/PH15ePv0yyhdzYg9a3R6jbQp1KZ0+eYly9qB6rB88QOHMPnFm3W/vDI9Vbve7u6Yv2s3yz1+HM6bnHHjfBuFz7YwZq/ufvl5F/+W//n/OfL2SjGa63kfOmw3LxmhWw6Zx3qhi5JLKCOawG8rasv/kzN8rC4RNjiVsGMGQ9s74aBtzUW59YQXO/Vxy2mLoE2ihdBZ1t/ubzvar3S7AKdvmQ7hCEo9K9yyJx8MjclHpXl+/xYO7+i2F9wLOXnQXJB8j17vWDs7VRKTxjziOJg+topC0hopY5lv32hfHTfLUQSdj+SA0rVpgV/E65aa1NjqRCHyTHBOuGEVcwrd2vcYDSKUwYyr8D2TJVSUQciEegEgAOr6YOON8ut5gyG1TpJ6+37SqRlzW3nC7JVWduBNUF95IjcG2X6ZM8JksKasi0lVE3ZEclOa+rNS9vAzD4XjwxprVa+3qOlFsNFiGg1XQWcTZOnhKxfNCodbN1s+blT5tYVXxNbobF8GO3dFy6M1c1fJEQdQkHiDX5B7Aau0m6uLWYHLNsHipp+Hxl+dSCYvDm6uLnNQY3gmDkt4DhIbZcEWvtsB1UbBhX3BbUZ5sUGb58DU9dDGfenoFcRgQ35+eZmTyPD8OwWJfcGcK/5Y0cq/jogWapNoiOalUre7ToNTuX/993CI9wH8eS/3v4VpuZcHsNmFpKkvptrEc76ei2tu+g7XIhCybewjhCafG7mhxtziRCBc8z5BjK+mBxszGfoRh/xRJKAnS+fOuUgJWCJXxURJ6M7zevU475x1RK0NQbtGUXJ/SyKXM5b3Hh97JY5jFsqlXYaCuD/1OKesdsEZr9e6n34ZEfuLB+iR2KwUPdoMT9XCLMWqyHiNv+nxfj+W0oddzxr2X/CWsc8Nl0txEPs1bN3iW7x0+usLyb7064vonAAD//8IHLlo="
}
