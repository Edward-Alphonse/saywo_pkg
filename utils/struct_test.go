package utils

import (
	"testing"
)

func TestTrans2Map(t *testing.T) {
	type args struct {
		model any
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]any
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				map[string]any{
					"user_id": 505140469949274820,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Trans2Map(tt.args.model)
			if (err != nil) != tt.wantErr {
				t.Errorf("Trans2Map() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("value: %v", got)
		})
	}
}

type model struct {
	UserId uint64 `json:"user_id"`
}

func TestTrans2Struct(t *testing.T) {
	type args struct {
		dict map[string]any
	}
	type testCase[T any] struct {
		name    string
		args    args
		want    *T
		wantErr bool
	}
	tests := []testCase[model]{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				map[string]any{
					"user_id": 505140469949274820,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Trans2Struct[model](tt.args.dict)
			if (err != nil) != tt.wantErr {
				t.Errorf("Trans2Struct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("value: %v", got)
		})
	}
}
