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

func Test_topoSort(t *testing.T) {
	type args struct {
		edges [][2]string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{"empty graph", args{edges: nil}, nil, false},
		{"single graph of two nodes", args{edges: [][2]string{{"A", "B"}}}, []string{"B", "A"}, false},
		{"simple graph of three nodes", args{edges: [][2]string{{"A", "B"}, {"B", "C"}}}, []string{"C", "B", "A"}, false},
		{"a way more complex graph of five nodes", args{edges: [][2]string{{"A", "B"}, {"B", "C"}, {"A", "C"}, {"D", "E"}, {"A", "E"}}}, []string{"C", "B", "E", "A", "D"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := topoSort(tt.args.edges)
			if (err != nil) != tt.wantErr {
				t.Errorf("topoSort() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("topoSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
