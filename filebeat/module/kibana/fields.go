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

package kibana

import (
	"github.com/T0kii/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "kibana", asset.ModuleFieldsPri, AssetKibana); err != nil {
		panic(err)
	}
}

// AssetKibana returns asset data.
// This is the base64 encoded zlib format compressed contents of module/kibana.
func AssetKibana() string {
	return "eJy8lk1z2jwQx+98ih0uXB4zQHgLh1ye5tDp9NZbp8Ms1tpWkSVHu4bm23dkk4wB4xCa1sd9+e/vL8uSI9jS8wq2eoMWewCixdAK+nWg3wPwZAiZVpCGvCKOvS5EO7uChx4AHHrhq1OloR5AoskoXlW5CCzm1NAPjzwXQc67sjhEWlTDU0tC4jwU6FnbFL7Uw4xLeXioa85rzmRi1s6utXpNnUz6lhF8/gQuAckISib/0gTI7GKNQgr2WjKQTDPQjqwM4RHjLBBoCyhCeSHgiUsjDCEEpdVPJb1KaTVsANAvzItqiceTO5rO5ouIlvebaDxRdxFOZ/NoOpnPx9PxYjoajfqNznrZtvS8d16duy0wpg6v/WBWqxezVXmXy34rs6IESyPvwMIdqbXb/KRYhqG0ky8UvBKGTqg7bwBFzjYOvboR9fqF/CPM+eR+tlFqFI1wRNF4TItoOZkl0XwxXcYKF2o5W17vAJVai1tX75Y7+ZkkGKh3gVYMkqEAHrvZIwNn6EmBuAsGvg8OW2LwHwxy9FsSbdPBj+uhFRkSWife5R9K7il3IRSE/xo8lpKRFR1jQFwX3u20It/p4LgHXnrONg8ezpiu/bNB1vH4Zt43P8hLrNWHejPwzbye0OSdwI8GWXTMhD7OTvGr9koZ9pmOM0hKk2hjSF3FblH0jq6HN85ty+Ld0HXbAfbfcKZntc3b+Qz6oZGAxpUMRlvi5m13ejc3xwqmfJS4xPkmwTFFGv4TgviwdSwLCn3c3P9L78lKLRtOphqkfXZOgq2j63PrJFUH1+1wF/WHnp6GGaEi3768iQl/LZZUlwRfKfE7AAD//8+r+Ws="
}
