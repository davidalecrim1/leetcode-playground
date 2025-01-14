package main

import "testing"

func Test_hammingWeight(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "11 has three set bits",
			args: args{
				n: 11,
			},
			want: 3,
		},
		{
			name: "one has one set bit",
			args: args{
				n: 1,
			},
			want: 1,
		},
		{
			name: "27 has four set bits",
			args: args{
				n: 27,
			},
			want: 4,
		},
		{
			name: "128 has four set bits",
			args: args{
				n: 128,
			},
			want: 1,
		},
		{
			name: "2147483645 has 30 set bit",
			args: args{
				n: 2147483645,
			},
			want: 30,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hammingWeight(tt.args.n); got != tt.want {
				t.Errorf("hammingWeight() = %v, want %v", got, tt.want)
			}
		})
	}
}
