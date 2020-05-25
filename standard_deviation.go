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
StandardDeviation returns the standard deviation of the last n values.

# Formula

![SD formula](https://wikimedia.org/api/rest_v1/media/math/render/svg/2845de27edc898d2a2a4320eda5f57e0dac6f650)

Where:

* _σ_ - value of standard deviation for N given probes.
* _N_ - number of probes in observation.
* _x<sub>i</sub>_ - i-th observed value from N elements observation.

# Parameters

* _n_ - number of periods (integer greater than 0)

# Example
```
sd, _ := NewStandardDeviation(4)
sd.Next(10.)
sd.Next(20.)
sd.Next(30.)
sd.Next(20.)
sd.Next(10.)
```
*/
type StandardDeviation struct {
	// number of periods (must be an integer greater than 0)
	n int

	// internal parameters for calculations
	index int
	count int

	m  float64
	m2 float64

	// slice of data needed for calculation
	data []float64
}

// NewStandardDeviation creates a new StandardDeviation with the given number of periods
// Example: NewStandardDeviation(9)
func NewStandardDeviation(n int) (*StandardDeviation, error) {
	if n <= 0 {
		return nil, ErrInvalidParameters
	}

	return &StandardDeviation{
		n: n,

		index: 0,
		count: 0,

		m:  0,
		m2: 0,

		data: make([]float64, n),
	}, nil
}

// Next takes the next input and returns the next StandardDeviation value
func (sd *StandardDeviation) Next(input float64) float64 {
	// add input to data
	sd.index = (sd.index + 1) % sd.n
	oldValue := sd.data[sd.index]
	sd.data[sd.index] = input

	if sd.count < sd.n {
		// not enough data for n periods yet
		sd.count++
		delta := input - sd.m
		sd.m += delta / float64(sd.count)
		delta2 := input - sd.m
		sd.m2 += delta * delta2
	} else {
		oldM := sd.m
		delta := input - oldValue
		sd.m += delta / float64(sd.n)

		delta2 := input - sd.m + oldValue - oldM
		sd.m2 += delta * delta2
	}

	return math.Sqrt(sd.m2 / float64(sd.count))
}

// Reset resets the indicators to a clean state
func (sd *StandardDeviation) Reset() {
	sd.index = 0
	sd.count = 0

	sd.m = 0
	sd.m2 = 0

	sd.data = make([]float64, sd.n)
}

func (sd *StandardDeviation) String() string {
	return fmt.Sprintf("SD(%d)", sd.n)
}
