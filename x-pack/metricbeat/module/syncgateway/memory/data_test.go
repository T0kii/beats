// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package memory

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	mbtest "github.com/T0kii/beats/v7/metricbeat/mb/testing"
	"github.com/T0kii/beats/v7/x-pack/metricbeat/module/syncgateway"
)

func TestData(t *testing.T) {
	mux := syncgateway.CreateTestMuxer()
	server := httptest.NewServer(mux)
	defer server.Close()

	config := syncgateway.GetConfig([]string{"memory"}, server.URL)

	f := mbtest.NewReportingMetricSetV2Error(t, config)
	if err := mbtest.WriteEventsReporterV2Error(f, t, ""); err != nil {
		t.Fatal("write", err)
	}
}

func TestFetch(t *testing.T) {
	mux := syncgateway.CreateTestMuxer()
	server := httptest.NewServer(mux)
	defer server.Close()

	config := syncgateway.GetConfig([]string{"memory"}, server.URL)
	metricSet := mbtest.NewReportingMetricSetV2Error(t, config)

	events, errs := mbtest.ReportingFetchV2Error(metricSet)
	if len(errs) > 0 {
		t.Fatalf("Expected 0 error, had %d. %v\n", len(errs), errs)
	}

	assert.NotEmpty(t, events)
	mbtest.TestMetricsetFieldsDocumented(t, metricSet, events)
}
