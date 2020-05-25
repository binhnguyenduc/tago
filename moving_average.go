/*
Copyright 2020 Binh Nguyen
Licensed under terms of MIT license (see LICENSE)
*/

package tago

import "fmt"

/*
MovingAverage returns the average in _n_ number of periods

# Formula

![SMA](https://wikimedia.org/api/rest_v1/media/math/render/svg/e2bf09dc6deaf86b3607040585fac6078f9c7c89)

Where:

* _SMA<sub>t</sub>_ - value of simple moving average at a point of time _t_
* _n_ - number of periods (length)
* _p<sub>t</sub>_ - input value at a point of time _t_

# Parameters

* _n_ - number of periods (integer greater than 0)

# Example
```
```
*/
type MovingAverage struct {
	// number of periods (must be an integer greater than 0)
	n int

	// internal parameters for calculations
	index int
	count int

	sum float64

	// slice of data needed for calculation
	data []float64
}

// NewMovingAverage creates a new MovingAverage with the given number of periods
// Example: NewMovingAverage(9)
func NewMovingAverage(n int) (*MovingAverage, error) {
	if n <= 0 {
		return nil, ErrInvalidParameters
	}

	return &MovingAverage{
		n: n,

		index: 0,
		count: 0,

		sum: 0,

		data: make([]float64, n),
	}, nil
}

// Next takes the next input and returns the next MovingAverage value
func (ma *MovingAverage) Next(input float64) float64 {
	// add input to data
	ma.index = (ma.index + 1) % ma.n
	oldValue := ma.data[ma.index]
	ma.data[ma.index] = input

	if ma.count < ma.n {
		ma.count++
	}

	ma.sum += input - oldValue
	return ma.sum / float64(ma.count)
}

// Reset resets the indicators to a clean state
func (ma *MovingAverage) Reset() {
	ma.index = 0
	ma.count = 0

	ma.sum = 0

	ma.data = make([]float64, ma.n)
}

func (ma *MovingAverage) String() string {
	return fmt.Sprintf("MA(%d)", ma.n)
}
