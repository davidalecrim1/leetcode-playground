package main

import "testing"

func Test_missingNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				nums: []int{3, 0, 1},
			},
			want: 2,
		},
		{
			args: args{
				nums: []int{0, 1},
			},
			want: 2,
		},
		{
			args: args{
				nums: []int{9, 6, 4, 2, 3, 5, 7, 0, 1},
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := missingNumber(tt.args.nums); got != tt.want {
				t.Errorf("missingNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
