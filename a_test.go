package main

import (
	"strings"
	"testing"
)

func TestSolveSmoke(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
		wantErr  bool
	}{
		{
			name: "Example 1",
			input: `2
1 2
2 1`,
			expected: "yes",
			wantErr:  false,
		},
		{
			name: "Example 2",
			input: `3
10 20 30
1 1 1
0 0 1`,
			expected: "no",
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := strings.NewReader(tt.input)
			got, err := solve(r)

			if (err != nil) != tt.wantErr {
				t.Errorf("solve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != tt.expected {
				t.Errorf("solve() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestSolveExtra(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
		wantErr  bool
	}{

		{
			name: "Already sorted",
			input: `2
5 0
0 3`,
			expected: "yes",
			wantErr:  false,
		},
		{
			name: "Invalid format",
			input: `abc
1 2
3 4`,
			expected: "no",
			wantErr:  true,
		},
		{
			name: "Sum mismatch",
			input: `2
3 1
2 2`,
			expected: "no",
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := strings.NewReader(tt.input)
			got, err := solve(r)

			if (err != nil) != tt.wantErr {
				t.Errorf("solve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != tt.expected {
				t.Errorf("solve() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestSolveBigNumbers(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
		wantErr  bool
	}{

		{
			name: "Large numbers match",
			input: `2
10000000000
0 1000000000`,
			expected: "yes",
			wantErr:  false,
		},
		{
			name: "Large numbers mismatch",
			input: `2
1000000000 0 
1 999999999`,
			expected: "no",
			wantErr:  false,
		},
		{
			name: "Max n=100 with large numbers",
			input: `100
` + strings.Repeat("1000000000"+strings.Repeat(" 1000000000", 99)+"\n", 100),
			expected: "yes",
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := strings.NewReader(tt.input)
			got, err := solve(r)

			if (err != nil) != tt.wantErr {
				t.Errorf("solve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != tt.expected {
				t.Errorf("solve() = %v, want %v", got, tt.expected)
			}
		})
	}
}
