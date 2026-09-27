package p004

import "testing"

func TestSolve(t *testing.T) {
	tests := []struct {
		name   string
		digits int
		want   uint64
	}{
		{
			name:   "longest palindrome produced by prod of 2 digit numbers",
			digits: 2,
			want:   9009,
		},
		{
			name:   "longest palindrome produced by prod of 3 digit numbers",
			digits: 3,
			want:   906609,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Solve(tt.digits)
			if got != tt.want {
				t.Errorf("Solve(%d) = %d; want %d", tt.digits, got, tt.want)
			}
		})
	}
}

func BenchmarkSolve(b *testing.B) {
	for b.Loop() {
		_ = Solve(4)
	}
}
