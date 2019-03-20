// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package beater

import (
	"github.com/wangjia184/beats/libbeat/logp"
	"github.com/wangjia184/beats/x-pack/functionbeat/licenser"
)

func checkLicense(log *logp.Logger, license licenser.License) bool {
	return licenser.CheckBasic(log, license) || licenser.CheckTrial(log, license)
}
