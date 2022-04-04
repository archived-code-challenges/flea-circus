package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

const (
	gridXSize = 30
	gridYSize = 30

	// bellRings sets the number of rings of the bell
	bellRings = 50
)

// remove an element of the slice using its index
func remove(s []string, index int) []string {
	s[index] = s[len(s)-1]
	return s[:len(s)-1]
}

type Flee struct {
	X int
	Y int
}

// possibleJumps returns the list of available moves that a Flee can perform.
func (f *Flee) possibleJumps() []string {
	var possibleJumps = []string{"up", "down", "left", "right"}
	if f.X == gridXSize-1 {
		possibleJumps = remove(possibleJumps, 3) // remove right
	}
	if f.X == 0 {
		possibleJumps = remove(possibleJumps, 2) // remove left
	}
	if f.Y == gridYSize-1 {
		possibleJumps = remove(possibleJumps, 1) // remove down
	}
	if f.Y == 0 {
		possibleJumps = remove(possibleJumps, 0) // remove up
	}

	return possibleJumps
}

// jump to one of the available cells randomly.
func (f *Flee) jump() {
	jumpTo := f.possibleJumps()
	switch jumpTo[rand.Intn(len(jumpTo))] {
	case "up":
		f.Y -= 1
	case "down":
		f.Y += 1
	case "left":
		f.X -= 1
	case "right":
		f.X += 1
	}
}

type Simulation struct {
	Flees []Flee
}

// initialize the simulation by adding one flee to each square. The size of the square is set by
// multiplying the X and Y axis (gridXSize * gridYSize).
func (s *Simulation) initialize() {
	s.Flees = make([]Flee, 0, gridXSize*gridYSize) // ayuda esto a velocidad
	for x := 0; x < gridXSize; x++ {
		for y := 0; y < gridYSize; y++ {
			s.Flees = append(s.Flees, Flee{X: x, Y: y})
		}
	}
}

// run the simulation "bellRings" times.
func (s *Simulation) run() {
	for ring := 0; ring < bellRings; ring++ {
		for i := range s.Flees {
			s.Flees[i].jump()
		}
	}
}

// unusedSquares returns the number of unused squares in a board of gridXSize * gridYSize by
// checking the position of all the flees contained in a simulation.
func (s *Simulation) unusedSquares() int {
	occupiedSquares := make(map[string]struct{})
	for i := range s.Flees {
		position := fmt.Sprintf("[%d,%d]", s.Flees[i].X, s.Flees[i].Y)
		occupiedSquares[position] = struct{}{}
	}
	return gridXSize*gridYSize - len(occupiedSquares)
}

func parseFlags() (runs, workers int) {
	numWorkers := flag.Int("workers", 1, "Number of concurrent workers")
	numRuns := flag.Int("runs", 1, "Number of runs")
	flag.Parse()

	return *numRuns, *numWorkers
}

// run the program "numRuns" times using the given num of "numRuns" and outputs the results.
func run(numRuns, numWorkers int) {
	wg := sync.WaitGroup{}
	wg.Add(numRuns)

	// workers is a limiting channel to control number of concurrent goroutines used
	workers := make(chan struct{}, numWorkers)

	var sum int64
	log.Printf("Running %d simulations on %d workers...", numRuns, numWorkers)
	start := time.Now()
	defer func() {
		log.Printf("%v elapsed", time.Since(start))
	}()

	for run := 0; run < numRuns; run++ {
		workers <- struct{}{}

		go func() { // perform simulations in parallel
			defer func() {
				<-workers
				wg.Add(-1)
			}()

			s := Simulation{}
			s.initialize()
			s.run() // single simulation run

			sum += int64(s.unusedSquares())
		}()
	}
	wg.Wait()

	// Calculate the average of the simulations performed
	log.Printf("Average unnoccupied squares after %d simulations: %f", numRuns, float64(sum)/float64(numRuns))
}

func main() {
	numWorkers, runs := parseFlags()
	run(numWorkers, runs)
}
