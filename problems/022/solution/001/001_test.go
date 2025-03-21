package main

import (
	"reflect"
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	tests := []struct {
		scenario string
		given    int
		want     []string
	}{
		{
			scenario: "input one",
			given:    1,
			want:     []string{"()"},
		},
		{
			scenario: "input three",
			given:    3,
			want:     []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
		{
			scenario: "input six",
			given:    6,
			want:     []string{"(((((())))))", "((((()()))))", "((((())())))", "((((()))()))", "((((())))())", "((((()))))()", "(((()(()))))", "(((()()())))", "(((()())()))", "(((()()))())", "(((()())))()", "(((())(())))", "(((())()()))", "(((())())())", "(((())()))()", "(((()))(()))", "(((()))()())", "(((()))())()", "(((())))(())", "(((())))()()", "((()((()))))", "((()(()())))", "((()(())()))", "((()(()))())", "((()(())))()", "((()()(())))", "((()()()()))", "((()()())())", "((()()()))()", "((()())(()))", "((()())()())", "((()())())()", "((()()))(())", "((()()))()()", "((())((())))", "((())(()()))", "((())(())())", "((())(()))()", "((())()(()))", "((())()()())", "((())()())()", "((())())(())", "((())())()()", "((()))((()))", "((()))(()())", "((()))(())()", "((()))()(())", "((()))()()()", "(()(((()))))", "(()((()())))", "(()((())()))", "(()((()))())", "(()((())))()", "(()(()(())))", "(()(()()()))", "(()(()())())", "(()(()()))()", "(()(())(()))", "(()(())()())", "(()(())())()", "(()(()))(())", "(()(()))()()", "(()()((())))", "(()()(()()))", "(()()(())())", "(()()(()))()", "(()()()(()))", "(()()()()())", "(()()()())()", "(()()())(())", "(()()())()()", "(()())((()))", "(()())(()())", "(()())(())()", "(()())()(())", "(()())()()()", "(())(((())))", "(())((()()))", "(())((())())", "(())((()))()", "(())(()(()))", "(())(()()())", "(())(()())()", "(())(())(())", "(())(())()()", "(())()((()))", "(())()(()())", "(())()(())()", "(())()()(())", "(())()()()()", "()((((()))))", "()(((()())))", "()(((())()))", "()(((()))())", "()(((())))()", "()((()(())))", "()((()()()))", "()((()())())", "()((()()))()", "()((())(()))", "()((())()())", "()((())())()", "()((()))(())", "()((()))()()", "()(()((())))", "()(()(()()))", "()(()(())())", "()(()(()))()", "()(()()(()))", "()(()()()())", "()(()()())()", "()(()())(())", "()(()())()()", "()(())((()))", "()(())(()())", "()(())(())()", "()(())()(())", "()(())()()()", "()()(((())))", "()()((()()))", "()()((())())", "()()((()))()", "()()(()(()))", "()()(()()())", "()()(()())()", "()()(())(())", "()()(())()()", "()()()((()))", "()()()(()())", "()()()(())()", "()()()()(())", "()()()()()()"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.scenario, func(t *testing.T) {
			got := generateParenthesis(tt.given)

			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("want %v got %v", tt.want, got)
			}
		})
	}

}
