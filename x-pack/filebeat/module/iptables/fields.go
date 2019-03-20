// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package iptables

import (
	"github.com/wangjia184/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "iptables", asset.ModuleFieldsPri, AssetIptables); err != nil {
		panic(err)
	}
}

// AssetIptables returns asset data.
// This is the base64 encoded gzipped contents of module/iptables.
func AssetIptables() string {
	return "eJy0l09vozwQxu/5FPMB3ub4HnJYqWq2UqV2W2nTvUaOPYAVY1N7SEQ//comUBKIQ0OWI3+e38x4/Ay+gy1WC5AFsY1CNwMgSQqP7gh03MqCpNEL+DEDAHgxolQIibGQMS2U1ClQhu1XoEzq5jOARKISbhE+ugPN8mNpf1FV4AJSa8ricGcA6K/HoAWJNfkwzF9dYBeKlKFde1b7qEEro9POzTN0f/1hqkQwScAHRY0UVGouSIGaZFI19dBIe2O3oFiFFgpryHCj5q1oL8xEmf1asQ2qXphSE6Zox0X69Lb7P4hBEIsRLUtz1LROFEtdj7rFam+sGEtt5SDIzeEeuMk3UjP/vi/dw8//YPkITAt4eRwTl0kShzRh1V6DQLNsvLTWx9eJNRKF5HnRQ3d79VJFHl7eDj057zw47dLOo4bMjcAjqcG0o/AD3isNpNhJUtwG9LSMYgpmWY501MITaK1cFGpRSIucBpiy+Cax0QImhEXnomCHH7fJ0+FHiZoj6DLfXMj2xN8mUP2HsX0hJmzIN8a3SI1ZysGcWpDmJi8UEq43FWHfoMZjf4X6eR/4EoUgGuUXJa0F7iTvT45vmeMyaABljMAiR7lDERypCOWIxFD414XvgfVGUr8ELjOWRjv0lxp4tQiXzJRyB9jKz0aTwG+0IftgfBGkQp1SNr23ap0IyJT0Lxa2lh23rMQnjZbVw5WT5XTQn0/6wvqGCMKYjzuwQ7tDcdq757p3BLSRPNfAt3VhT/yOCTO+vQnznm+12SsU9Q/VCPJeamH2k+H3YoeWpEMR4qhVwcnP2EQgmvjLupI5wsrAs9xd9olSTNo778sr907Pn66r8HOQaX5JfTQZMoE2/BYXrFKGRbPfyI9SkpxSguY0ZSy8H+Ta40rtimNrczosP43Gk3zPG8yFQgE8eU3wmoOt33P029Jfaz8fgbelwnW9R6/E9+CrDIPsYevDXlImdWiZcNsNjpdeUMenp9tE5A9UHjGf/Q0AAP//R0INxw=="
}
