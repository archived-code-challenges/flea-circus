package main

import (
	"io/ioutil"
	"log"
	"testing"
)

func testEq(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	contains := func(s []string, elem string) bool {
		for i := range s {
			if s[i] == elem {
				return true
			}
		}
		return false
	}

	for i := range a {
		if !contains(b, a[i]) {
			return false
		}
	}
	return true
}

func TestFlee_possibleJumps(t *testing.T) {
	tests := []struct {
		name string
		x    int
		y    int
		want []string
	}{
		{
			name: "top left corner",
			x:    0, y: 0,
			want: []string{"down", "right"},
		},
		{
			name: "top right corner",
			x:    gridXSize - 1, y: 0,
			want: []string{"down", "left"},
		},
		{
			name: "bottom left corner",
			x:    0, y: gridYSize - 1,
			want: []string{"up", "right"},
		},
		{
			name: "bottom right corner",
			x:    gridXSize - 1, y: gridYSize - 1,
			want: []string{"left", "up"},
		},
		{
			name: "top side",
			x:    gridXSize / 2, y: 0,
			want: []string{"down", "right", "left"},
		},
		{
			name: "bottom side",
			x:    gridXSize / 2, y: gridYSize - 1,
			want: []string{"up", "left", "right"},
		},
		{
			name: "left side",
			x:    0, y: gridYSize / 2,
			want: []string{"up", "down", "right"},
		},
		{
			name: "right side",
			x:    gridXSize - 1, y: gridYSize / 2,
			want: []string{"up", "down", "left"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Flee{X: tt.x, Y: tt.y}
			if got := f.possibleJumps(); !testEq(got, tt.want) {
				t.Errorf("Flee.possibleJumps([%d, %d]) = %v, want %v", f.X, f.Y, got, tt.want)
			}
		})
	}
}

func benchmarkRun(b *testing.B, i int) {
	log.SetOutput(ioutil.Discard) // disable log output when running benchmarks
	for n := 1; n < b.N; n++ {
		run(n, i)
	}
}

func BenchmarkRun1Worker(b *testing.B)    { benchmarkRun(b, 1) }
func BenchmarkRun10Workers(b *testing.B)  { benchmarkRun(b, 10) }
func BenchmarkRun100Workers(b *testing.B) { benchmarkRun(b, 100) }
