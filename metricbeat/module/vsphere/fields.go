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

package vsphere

import (
	"github.com/T0kii/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "vsphere", asset.ModuleFieldsPri, AssetVsphere); err != nil {
		panic(err)
	}
}

// AssetVsphere returns asset data.
// This is the base64 encoded zlib format compressed contents of module/vsphere.
func AssetVsphere() string {
	return "eJzsl8+OmzAQxu95itHedx+AQ6Vqq20vaSttt9fIgSG4MQyyh6zo01e2IeKP002KaS/1ASnYfN8PzHxD7uGIbQInUxeocQPAkhUmcHd6dmfuNgAZmlTLmiVVCbzbAAB0s1BS1ih7mUaFwmACB7EByCWqzCRu6T1UosShhR3c1naxpqbuzgRcxkJDsUywMExnubDkRdluKiAyvo9+TDGGKPY4muhJjti+ks4mc7/hseNDzzTX7Q1zY/XjWT5JhaY1jCXMhHvPVNQildw+MLFQD/uW0QQJFFWH2+y/WUVwikA5cIHBjbEjJ10KTmBuP+PMNWJUzCeNGJ2yMZhFpXwxmK1DWaccZDSpUJjtckViuuAa1hp1ihVfS9stP8/OgqEgw0syYXL9v46DT2T4chKkdeM3pyx+xnyBHr++gKxgO1Ed2vociOfrU+AKY1fY8XxdWb9hW2JJeq1i3Tpxax+SfrtKO7i1cjkS3jpxvBSuQn4lfdzZX2GwPyraz14WxrLn7yCpuRGqFGkhq0XfLxeVbk8tG3sPcnqfS4NrIjgys4eVsrJvJpMsXyujv/ttgK3fh8txTRHfsanplxq1YFkd4Nl/yv1vGus1jfcnIZXYq5s6x6FBw6v1D8rhozVYnNSO1VXoiqg+HaK0vPiPddT4oj1X9wrGhx32weWsaWOYyp3vFEFC2v/A2T8Bf3K3IM4enTEEjP9qr/4VAAD//9YITTs="
}
