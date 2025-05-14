package config

import (
	"reflect"
	"testing"
)

func Test_getPoemConfig(t *testing.T) {
	tests := []struct {
		name string
		want Poem
	}{
		{
			name: "test",
			want: Poem{
				[]PoemDetails{{
					Day:   1,
					Lines: []string{"this is test", "line 2", "line3"},
				},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getPoemConfig(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getPoemConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
