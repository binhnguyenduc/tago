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
Maximum returns the highest value in a given time frame

# Parameters

* _n_ - size of time frame (integer greater than 0)

# Example
```
```
*/
type Maximum struct {
	// number of periods (must be an integer greater than 0)
	n int

	// internal parameters for calculations
	maxIndex int
	curIndex int

	// slice of data needed for calculation
	data []float64
}

// NewMaximum creates a new Maximum with the given number of periods
// Example: NewMaximum(9)
func NewMaximum(n int) (*Maximum, error) {
	if n <= 0 {
		return nil, ErrInvalidParameters
	}

	data := make([]float64, n)
	for i := range data {
		data[i] = math.Inf(-1)
	}
	return &Maximum{
		n: n,

		maxIndex: 0,
		curIndex: 0,

		data: data,
	}, nil
}

// Next takes the next input and returns the next Maximum value
func (m *Maximum) Next(input float64) float64 {
	// add input to data
	m.curIndex = (m.curIndex + 1) % m.n
	m.data[m.curIndex] = input

	if input > m.data[m.maxIndex] {
		m.maxIndex = m.curIndex
	} else if m.curIndex == m.maxIndex {
		m.maxIndex = findMaxIndex(m.data)
	}

	return m.data[m.maxIndex]
}

func findMaxIndex(data []float64) int {
	max := math.Inf(-1)
	index := 0

	for i, v := range data {
		if v > max {
			max = v
			index = i
		}
	}

	return index
}

// Reset resets the indicators to a clean state
func (m *Maximum) Reset() {
	data := make([]float64, m.n)
	for i := range data {
		data[i] = math.Inf(-1)
	}
	m.data = data
	m.curIndex = 0
	m.maxIndex = 0
}

func (m *Maximum) String() string {
	return fmt.Sprintf("Max(%d)", m.n)
}
