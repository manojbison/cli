// Copyright © 2019 The Tekton Authors.
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

package formatted

import (
	"reflect"
	"testing"

	v1 "k8s.io/api/core/v1"
)

func TestFormatDesc(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "Empty string",
			input: "",
			want:  "",
		},
		{
			name:  "String with 19 char",
			input: "test description to",
			want:  "test description to",
		},
		{
			name:  "String with 20 chars",
			input: "test description to ",
			want:  "test description to ",
		},
		{
			name:  "String with 21 chars",
			input: "test description to t",
			want:  "test description to...",
		},
		{
			name:  "Long string",
			input: "test description to test trimming",
			want:  "test description to...",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatDesc(tt.input); got != tt.want {
				t.Errorf("Input = %s, want %s", got, tt.want)
			}
		})
	}
}

func TestRemoveLastAppliedConfig(t *testing.T) {
	tests := []struct {
		name  string
		input map[string]string
		want  map[string]string
	}{
		{
			name:  "Empty Annotation",
			input: map[string]string{},
			want:  map[string]string{},
		},
		{
			name: "Annotations with last-applied-configuration",
			input: map[string]string{
				v1.LastAppliedConfigAnnotation: "JSON String",
				"tekton.dev/tags":              "game",
			},
			want: map[string]string{
				"tekton.dev/tags": "game",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !reflect.DeepEqual(tt.want, RemoveLastAppliedConfig(tt.input)) {
				t.Error("input = %w, want = %w", tt.input, tt.want)
			}
		})
	}
}
