package main

import "testing"

func Test_getSum(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test case 001",
			args: args{
				a: 2,
				b: 3,
			},
			want: 5,
		},
		{
			name: "test case 002",
			args: args{
				a: 3,
				b: 3,
			},
			want: 6,
		},
		{
			name: "test case 003",
			args: args{
				a: 1,
				b: 2,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSum(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("getSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
