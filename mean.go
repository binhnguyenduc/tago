/*
Copyright 2020 Binh Nguyen
Licensed under terms of MIT license (see LICENSE)
*/

package tago

import (
	"fmt"
)

/*
Mean returns the arithmetic mean value of the last n values.

# Formula

![SD formula](https://wikimedia.org/api/rest_v1/media/math/render/svg/97821b8c43e3182faa22db06932846d1550866fb_

Where:

* _n_ - number of probes in observation.

# Parameters

* _n_ - number of periods (integer greater than 0)

# Example
```
```
*/
type Mean struct {
	// number of periods (must be an integer greater than 0)
	n int

	// internal parameters for calculations
	index int
	count int

	sum float64

	// slice of data needed for calculation
	data []float64
}

// NewMean creates a new Mean with the given number of periods
// Example: NewMean(9)
func NewMean(n int) (*Mean, error) {
	if n <= 0 {
		return nil, ErrInvalidParameters
	}

	return &Mean{
		n: n,

		index: 0,
		count: 0,

		sum: 0,

		data: make([]float64, n),
	}, nil
}

// Next takes the next input and returns the next Mean value
func (m *Mean) Next(input float64) float64 {
	// add input to data
	m.index = (m.index + 1) % m.n
	oldValue := m.data[m.index]
	m.data[m.index] = input

	if m.count < m.n {
		// not enough data for n periods yet
		m.count++
		m.sum += input
	} else {
		delta := input - oldValue
		m.sum += delta
	}

	return m.sum / float64(m.count)
}

// Reset resets the indicators to a clean state
func (m *Mean) Reset() {
	m.index = 0
	m.count = 0

	m.sum = 0

	m.data = make([]float64, m.n)
}

func (m *Mean) String() string {
	return fmt.Sprintf("Mean(%d)", m.n)
}
