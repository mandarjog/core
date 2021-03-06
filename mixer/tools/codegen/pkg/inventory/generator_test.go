// Copyright 2017 Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package inventory

import (
	"bytes"
	"testing"
)

var header = `// Copyright 2017 Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// THIS FILE IS AUTOMATICALLY GENERATED.

package adapter
`

var empty = header + `
import (
	adptr "istio.io/core/mixer/pkg/adapter"
)

// Inventory returns the inventory of all available adapters.
func Inventory() []adptr.InfoFn {
	return []adptr.InfoFn{}
}
`

var example = header + `
import (
	kubernetes "istio.io/core/mixer/adapter/kubernetes"
	noop "istio.io/core/mixer/adapter/noop"
	prometheus "istio.io/core/mixer/adapter/prometheus"
	adptr "istio.io/core/mixer/pkg/adapter"
)

// Inventory returns the inventory of all available adapters.
func Inventory() []adptr.InfoFn {
	return []adptr.InfoFn{
		kubernetes.GetInfo,
		noop.GetInfo,
		prometheus.GetInfo,
	}
}
`

func TestGenerate(t *testing.T) {
	exampleDeps := map[string]string{
		"noop":       "istio.io/core/mixer/adapter/noop",
		"prometheus": "istio.io/core/mixer/adapter/prometheus",
		"kubernetes": "istio.io/core/mixer/adapter/kubernetes",
	}
	tests := []struct {
		name    string
		args    map[string]string
		wantOut string
		wantErr bool
	}{
		{"no adapters", map[string]string{}, empty, false},
		{"example deps", exampleDeps, example, false},
		{"missing import", map[string]string{"missing": ""}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			if err := Generate(tt.args, out); (err != nil) != tt.wantErr {
				t.Errorf("Generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOut := out.String(); gotOut != tt.wantOut {
				t.Errorf("Generate() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}
