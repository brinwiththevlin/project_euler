package p006

import "testing"

func TestSolve(t *testing.T) {
	tests := []struct {
		name  string
		limit int
		want  uint64
	}{
		{
			name:  "sum square difference up to 10",
			limit: 10,
			want:  2640,
		},
		{
			name:  "sum square difference up to 100",
			limit: 100,
			want:  25164150,
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

func BenchmarkSolve(b *testing.B) {
	b.Skip("TODO: set a real benchmark input for problem 6, then remove this skip")
	for b.Loop() {
		_ = Solve(0)
	}
}
