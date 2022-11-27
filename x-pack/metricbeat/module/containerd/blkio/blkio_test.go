// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

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

//go:build !integration
// +build !integration

package blkio

import (
	"testing"

	mbtest "github.com/T0kii/beats/v7/metricbeat/mb/testing"

	"github.com/T0kii/beats/v7/metricbeat/helper/prometheus/ptest"
	_ "github.com/T0kii/beats/v7/x-pack/metricbeat/module/containerd"
)

func TestEventMapping(t *testing.T) {
	ptest.TestMetricSet(t, "containerd", "blkio",
		ptest.TestCases{
			{
				MetricsFile:  "../_meta/test/containerd.v1.5.2",
				ExpectedFile: "./_meta/test/containerd.v1.5.2.expected",
			},
		},
	)
}

func TestData(t *testing.T) {
	mbtest.TestDataFiles(t, "containerd", "blkio")
}
