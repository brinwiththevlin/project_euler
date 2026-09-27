package p001

import "testing"

func TestSolve(t *testing.T) {
	tests := []struct {
		name  string
		limit int
		want  uint64
	}{
		{
			name:  "Project Euler Example (under 10)",
			limit: 10,
			want:  23,
		},
		{
			name:  "Project Euler Problem 1 (mulitiples of 3 or 5)",
			limit: 1000,
			want:  233168, // Update this with your verified answer once discovered
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Solve(tt.limit)
			if got != tt.want {
				t.Errorf("Solve(%d) = %d; want %d", tt.limit, got, tt.want)
			}
		})
	}
}

// BenchmarkSolve measures how efficient your sorting and map lookup logic is.
// Ideally, this algorithm should finish well within a few milliseconds.
func BenchmarkSolve(b *testing.B) {
	for b.Loop() {
		_ = Solve(1000)
	}
}
