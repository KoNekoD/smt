package smt

import (
	"reflect"
	"testing"
)

func TestSliceReverse(t *testing.T) {
	type args[T any] struct {
		s []T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		{
			name: "reverse test odd",
			args: args[int]{
				[]int{1, 2, 3, 4, 5},
			},
			want: []int{5, 4, 3, 2, 1},
		},

		{
			name: "reverse test even",
			args: args[int]{
				[]int{1, 2, 3, 4, 5, 6},
			},
			want: []int{6, 5, 4, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := SliceReverse(tt.args.s); !reflect.DeepEqual(
					got,
					tt.want,
				) {
					t.Errorf("SliceReverse() = %v, want %v", got, tt.want)
				}

				if tt.args.s[0] != 1 {
					t.Errorf("Original slice should not be modified")
				}
			},
		)
	}
}
