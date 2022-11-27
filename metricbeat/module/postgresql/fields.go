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

package postgresql

import (
	"github.com/T0kii/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "postgresql", asset.ModuleFieldsPri, AssetPostgresql); err != nil {
		panic(err)
	}
}

// AssetPostgresql returns asset data.
// This is the base64 encoded zlib format compressed contents of module/postgresql.
func AssetPostgresql() string {
	return "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"
}
