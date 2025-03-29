package main

import "testing"

func Test_climbStairs(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test case 1",
			args: args{
				n: 1,
			},
			want: 1,
		},
		{
			name: "test case 2",
			args: args{
				n: 2,
			},
			want: 2,
		},
		{
			name: "test case 3",
			args: args{
				n: 3,
			},
			want: 3,
		},
		{
			name: "test case 5",
			args: args{
				n: 5,
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairs(tt.args.n); got != tt.want {
				t.Errorf("climbStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
