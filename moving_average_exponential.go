/*
Copyright 2020 Binh Nguyen
Licensed under terms of MIT license (see LICENSE)
*/

package tago

import "fmt"

/*
ExponentialMovingAverage returns the exponential average in _n_ number of periods

# Formula

![EMA formula](https://wikimedia.org/api/rest_v1/media/math/render/svg/05d06bdbee2c14031fd91ead6f5f772aec1ec964)

Where:

* _EMA<sub>t</sub>_ - is the value of the EMA at any time period _t_.
* _EMA<sub>t-1</sub>_ - is the value of the EMA at the previous period _t-1_.
* _p<sub>t</sub>_ - is the input value at a time period t.
* _α_ - is the coefficient that represents the degree of weighting decrease, a constant smoothing factor between 0 and 1.

_α_ is calculated with the following formula:

![alpha formula](https://wikimedia.org/api/rest_v1/media/math/render/svg/d9f6258e152db0644af548972bd6c50a8becf7ee)

# Parameters

* _n_ - number of periods (integer greater than 0)

# Example
```
```
*/
type ExponentialMovingAverage struct {
	// number of periods (must be an integer greater than 0)
	n int

	// internal parameters for calculation
	k       float64
	current float64
	isNew   bool
}

// NewExponentialMovingAverage creates a new ExponentialMovingAverage with the given number of periods
// Example: NewExponentialMovingAverage(9)
func NewExponentialMovingAverage(n int) (*ExponentialMovingAverage, error) {
	if n <= 0 {
		return nil, ErrInvalidParameters
	}

	return &ExponentialMovingAverage{
		n: n,

		k:       2. / (float64(n) + 1.),
		current: 0,
		isNew:   true,
	}, nil
}

// Next takes the next input and returns the next ExponentialMovingAverage value
func (ma *ExponentialMovingAverage) Next(input float64) float64 {
	if ma.isNew {
		ma.isNew = false
		ma.current = input
	} else {
		ma.current = ma.k*input + (1.-ma.k)*ma.current
	}
	return ma.current
}

// Reset resets the indicators to a clean state
func (ma *ExponentialMovingAverage) Reset() {
	ma.isNew = true
	ma.current = 0
}

func (ma *ExponentialMovingAverage) String() string {
	return fmt.Sprintf("EMA(%d)", ma.n)
}
