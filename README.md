# Project Euler: Problem 213 "Flea Circus"

A 30×30 grid of squares contains 900 fleas, initially one flea per square.
When a bell is rung, each flea jumps to an adjacent square at random (usually 4 possibilities, except for fleas on the edge of the grid or at the corners).

What is the expected number of unoccupied squares after 50 rings of the bell? Give your answer rounded to six decimal places.

## Implementation plan

- [x] Implement single simulation
- [x] Run multiple simulations and calculate average
- [x] Run simulations in parallel
- [x] Optimize for better speed, less memory allocations

## Run

    go run main.go --workers=1 --runs=1000

*Output*

    2022/04/04 22:17:16 Running 1000 simulation(s) in 1 worker(s)...
    2022/04/04 22:17:19 Average unnoccupied squares after 1000 simulations: 330.531000
    2022/04/04 22:17:19 2.844869326s elapsed

## Benchmarking

    go test -bench=. -benchtime=50x -benchmem

> "-benchtime=50x" sets the explicit iteration count to 50 to shorten the benchmark execution time

*Output*

    goos: darwin
    goarch: amd64
    pkg: github.com/noelruault/flea-circus
    cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
    BenchmarkRun1Worker-12                50          74098303 ns/op        72206676 B/op    1125144 allocs/op
    BenchmarkRun10Workers-12              50         138845500 ns/op        72192206 B/op    1125151 allocs/op
    BenchmarkRun100Workers-12             50         186722515 ns/op        72189076 B/op    1125161 allocs/op

## Race condition checks

Running the program with the race detector enabled displays it's absence of data races

    go run -race main.go --workers=10 --runs=500
