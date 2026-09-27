package p005

import "testing"

func TestSolve(t *testing.T) {
	tests := []struct {
		name  string
		limit int
		want  uint64
	}{
		{
			name:  "smallest number divisible by all pos int less than 10",
			limit: 10,
			want:  2520,
		},
		{
			name:  "smallest number divisible by all pos int less than 20",
			limit: 20,
			want:  232792560,
		},
		// TODO: add the example from the problem statement and the real answer.
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

func BenchmarkSolve(b *testing.B) {
	b.Skip("TODO: set a real benchmark input for problem 5, then remove this skip")
	for b.Loop() {
		_ = Solve(20)
	}
}
