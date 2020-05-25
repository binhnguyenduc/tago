/*
Copyright 2020 Binh Nguyen
Licensed under terms of MIT license (see LICENSE)
*/

package tago

/*
BollingerBands
The Bollinger Bands are represented by Average EMA and standard deviaton that is moved 'k' times away in both directions from calculated average value.

# Formula

See SMA, SD documentation.

* _BB<sub>Middle Band</sub>_ - Simple Moving Average (SMA).
* _BB<sub>Upper Band</sub>_ = SMA + SD of observation * multipler (usually 2.0)
* _BB<sub>Lower Band</sub>_ = SMA - SD of observation * multipler (usually 2.0)

# Parameters

* _n_ - number of periods (integer greater than 0)

# Example
```
```
*/
type BollingerBands struct {
	// number of periods (must be an integer greater than 0)
	n          int
	multiplier int

	// internal parameters for calculations
	sd *StandardDeviation
}

// // NewBollingerBands creates a new BollingerBands with the given number of periods
// // Example: NewBollingerBands(9)
// func NewBollingerBands(n, multiplier int) (*BollingerBands, error) {
// 	if n <= 0 || multiplier <= 0 {
// 		return nil, ErrInvalidParameters
// 	}

// 	sd, err := NewStandardDeviation(n)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &BollingerBands{
// 		n:          n,
// 		multiplier: multiplier,

// 		sd: sd,
// 	}, nil
// }

// // Next takes the next input and returns the next BollingerBands value
// func (bb *BollingerBands) Next(input float64) (float64, float64, float64) {
// 	sd := bb.sd.Next(input)
// 	// mean := bb.sd.Mean()

// 	// return mean + float64(bb.multiplier)*sd, mean, mean - float64(bb.multiplier)*sd
// }
