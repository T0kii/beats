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

package state_storageclass

import (
	"fmt"

	p "github.com/T0kii/beats/v7/metricbeat/helper/prometheus"
	"github.com/T0kii/beats/v7/metricbeat/mb"
	k8smod "github.com/T0kii/beats/v7/metricbeat/module/kubernetes"
	"github.com/T0kii/beats/v7/metricbeat/module/kubernetes/util"
	"github.com/elastic/elastic-agent-autodiscover/kubernetes"
)

func init() {
	mb.Registry.MustAddMetricSet("kubernetes", "state_storageclass",
		NewStorageClassMetricSet,
		mb.WithHostParser(p.HostParser))
}

// StorageClassMetricSet is a prometheus based MetricSet that fetches
// and reports kube-state-metrics storage class metrics
type StorageClassMetricSet struct {
	mb.BaseMetricSet
	prometheus p.Prometheus
	mapping    *p.MetricsMapping
	mod        k8smod.Module
	enricher   util.Enricher
}

// NewStorageClassMetricSet returns a prometheus based metricset for Storage classes
func NewStorageClassMetricSet(base mb.BaseMetricSet) (mb.MetricSet, error) {
	prometheus, err := p.NewPrometheusClient(base)
	if err != nil {
		return nil, err
	}
	mod, ok := base.Module().(k8smod.Module)
	if !ok {
		return nil, fmt.Errorf("must be child of kubernetes module")
	}
	return &StorageClassMetricSet{
		BaseMetricSet: base,
		prometheus:    prometheus,
		mod:           mod,
		enricher:      util.NewResourceMetadataEnricher(base, &kubernetes.StorageClass{}, mod.GetMetricsRepo(), false),
		mapping: &p.MetricsMapping{
			Metrics: map[string]p.MetricMap{
				"kube_storageclass_info": p.InfoMetric(),
				"kube_storageclass_labels": p.ExtendedInfoMetric(
					p.Configuration{
						StoreNonMappedLabels:     true,
						NonMappedLabelsPlacement: mb.ModuleDataKey + ".labels",
						MetricProcessingOptions:  []p.MetricOption{p.OpLabelKeyPrefixRemover("label_")},
					}),
				"kube_storageclass_created": p.Metric("created", p.OpUnixTimestampValue()),
			},
			Labels: map[string]p.LabelMap{
				"storageclass":        p.KeyLabel("name"),
				"provisioner":         p.Label("provisioner"),
				"reclaimPolicy":       p.Label("reclaim_policy"),
				"reclaim_policy":      p.Label("reclaim_policy"),
				"volumeBindingMode":   p.Label("volume_binding_mode"),
				"volume_binding_mode": p.Label("volume_binding_mode"),
			},
		},
	}, nil
}

// Fetch prometheus metrics and treats those prefixed by mb.ModuleDataKey as
// module rooted fields at the event that gets reported
func (m *StorageClassMetricSet) Fetch(reporter mb.ReporterV2) {
	m.enricher.Start()

	families, err := m.mod.GetStateMetricsFamilies(m.prometheus)
	if err != nil {
		m.Logger().Error(err)
		reporter.Error(err)
		return
	}
	events, err := m.prometheus.ProcessMetrics(families, m.mapping)
	if err != nil {
		m.Logger().Error(err)
		reporter.Error(err)
		return
	}

	m.enricher.Enrich(events)

	for _, event := range events {

		e, err := util.CreateEvent(event, "kubernetes.storageclass")
		if err != nil {
			m.Logger().Error(err)
		}

		if reported := reporter.Event(e); !reported {
			m.Logger().Debug("error trying to emit event")
			return
		}
	}
}

// Close stops this metricset
func (m *StorageClassMetricSet) Close() error {
	m.enricher.Stop()
	return nil
}
