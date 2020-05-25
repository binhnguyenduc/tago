/*
Copyright 2020 Binh Nguyen
Licensed under terms of MIT license (see LICENSE)
*/

package tago

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

/*
Median returns the median value of the last n values.

Where:

* _n_ - number of probes in observation.

# Parameters

* _n_ - number of periods (integer greater than 0)

# Example
```
```
*/
type Median struct {
	// number of periods (must be an integer greater than 0)
	n int

	// internal parameters for calculations
	index int

	// slice of data needed for calculation
	data []float64
}

// NewMedian creates a new Median with the given number of periods
// Example: NewMedian(9)
func NewMedian(n int) (*Median, error) {
	if n <= 0 {
		return nil, ErrInvalidParameters
	}

	return &Median{
		n: n,

		index: 0,

		data: make([]float64, 0, n),
	}, nil
}

// Next takes the next input and returns the next Median value
func (m *Median) Next(input float64) float64 {
	// add input to data
	if len(m.data) < m.n {
		m.data = append(m.data, input)
	} else {
		m.index = m.index % m.n
		m.data[m.index] = input
	}
	m.index++
	return quickselectMedian(m.data, defaultPivotFunc)
}

// Reset resets the indicators to a clean state
func (m *Median) Reset() {
	m.index = 0

	m.data = make([]float64, 0, m.n)
}

func (m *Median) String() string {
	return fmt.Sprintf("Median(%d)", m.n)
}

func quickselect(l []float64, k int, pivotFn func([]float64) float64) float64 {
	/*
	   Select the kth element in l (0 based)
	   :param l: List of numerics
	   :param k: Index
	   :param pivot_fn: Function to choose a pivot, defaults to random.choice
	   :return: The kth element of l
	*/
	if len(l) == 1 {
		if k != 0 {
			// only here for completeness. should never happen
			log.Fatal("invalid")
		}
		return l[0]
	}

	pivot := pivotFn(l)
	var lows, highs, pivots []float64
	for _, v := range l {
		if v < pivot {
			lows = append(lows, v)
			continue
		}
		if v > pivot {
			highs = append(highs, v)
			continue
		}
		if v == pivot {
			pivots = append(pivots, v)
			continue
		}
	}

	if k < len(lows) {
		return quickselect(lows, k, pivotFn)
	} else if k < len(lows)+len(pivots) {
		// We got lucky and guessed the median
		return pivots[0]
	} else {
		return quickselect(highs, k-len(lows)-len(pivots), pivotFn)
	}
}

func quickselectMedian(l []float64, pivotFn func([]float64) float64) float64 {
	if len(l)%2 == 1 {
		return quickselect(l, len(l)/2, pivotFn)
	}

	return 0.5 * (quickselect(l, len(l)/2-1, pivotFn) + quickselect(l, len(l)/2, pivotFn))
}

func defaultPivotFunc(l []float64) float64 {
	rand.Seed(time.Now().Unix())
	return l[rand.Intn(len(l))]
}
