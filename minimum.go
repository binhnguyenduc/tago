/*
Copyright 2020 Binh Nguyen
Licensed under terms of MIT license (see LICENSE)
*/

package tago

import (
	"fmt"
	"math"
)

/*
Minimum returns the lowest value in a given time frame

# Parameters

* _n_ - size of time frame (integer greater than 0)

# Example
```
```
*/
type Minimum struct {
	// number of periods (must be an integer greater than 0)
	n int

	// internal parameters for calculations
	minIndex int
	curIndex int

	// slice of data needed for calculation
	data []float64
}

// NewMinimum creates a new Minimum with the given number of periods
// Example: NewMinimum(9)
func NewMinimum(n int) (*Minimum, error) {
	if n <= 0 {
		return nil, ErrInvalidParameters
	}

	data := make([]float64, n)
	for i := range data {
		data[i] = math.Inf(1)
	}
	return &Minimum{
		n: n,

		minIndex: 0,
		curIndex: 0,

		data: data,
	}, nil
}

// Next takes the next input and returns the next Minimum value
func (m *Minimum) Next(input float64) float64 {
	// add input to data
	m.curIndex = (m.curIndex + 1) % m.n
	m.data[m.curIndex] = input

	if input < m.data[m.minIndex] {
		m.minIndex = m.curIndex
	} else if m.curIndex == m.minIndex {
		m.minIndex = findMinIndex(m.data)
	}

	return m.data[m.minIndex]
}

func findMinIndex(data []float64) int {
	min := math.Inf(1)
	index := 0

	for i, v := range data {
		if v < min {
			min = v
			index = i
		}
	}

	return index
}

// Reset resets the indicators to a clean state
func (m *Minimum) Reset() {
	data := make([]float64, m.n)
	for i := range data {
		data[i] = math.Inf(-1)
	}
	m.data = data
	m.curIndex = 0
	m.minIndex = 0
}

func (m *Minimum) String() string {
	return fmt.Sprintf("Min(%d)", m.n)
}
