// Copyright 2023 Google LLC
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

package impl

import (
	"reflect"
	"testing"
)

func Test_topologicalSort(t *testing.T) {
	type args struct {
		edges [][2]string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"empty graph", args{edges: nil}, nil},
		{"single graph of two nodes", args{edges: [][2]string{{"A", "B"}}}, []string{"B", "A"}},
		{"simple graph of three nodes", args{edges: [][2]string{{"A", "B"}, {"B", "C"}}}, []string{"C", "B", "A"}},
		{"a way more complex graph of five nodes", args{edges: [][2]string{{"A", "B"}, {"B", "C"}, {"A", "C"}, {"D", "E"}, {"A", "E"}}}, []string{"C", "B", "E", "A", "D"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := topologicalSort(tt.args.edges)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("topoSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
