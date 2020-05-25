/*
Copyright 2020 Binh Nguyen
Licensed under terms of MIT license (see LICENSE)
*/

package tago

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestNewExponentialMovingAverage(t *testing.T) {
	tests := map[string]struct {
		input   int
		want    *ExponentialMovingAverage
		wantErr error
	}{
		"negative n": {input: -3, want: nil, wantErr: ErrInvalidParameters},
		"zero n":     {input: 0, want: nil, wantErr: ErrInvalidParameters},
		"positive n": {input: 9, want: &ExponentialMovingAverage{n: 9, k: 2. / 10., isNew: true}, wantErr: nil},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			gotSD, gotErr := NewExponentialMovingAverage(tc.input)
			if tc.wantErr != nil { // only check error returned if expecting one
				assert.EqualError(t, gotErr, tc.wantErr.Error(), "must return the correct error")
			}
			assert.Equal(t, tc.want, gotSD, "must return the correct value")
		})
	}
}

func TestExponentialMovingAverageNext(t *testing.T) {
	sd, _ := NewExponentialMovingAverage(3)
	tests := []struct {
		input float64
		want  float64
	}{
		{input: 2., want: 2.},
		{input: 5., want: 3.5},
		{input: 1., want: 2.25},
		{input: 6.25, want: 4.25},
	}
	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			got := sd.Next(tc.input)
			diff := cmp.Diff(tc.want, got, floatComparer)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}

func TestExponentialMovingAverageReset(t *testing.T) {
	sd, _ := NewExponentialMovingAverage(3)
	tests := []struct {
		input float64
		want  float64
	}{
		{input: 2., want: 2.},
		{input: 5., want: 3.5},
		{input: 1., want: 2.25},
	}
	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			got := sd.Next(tc.input)
			diff := cmp.Diff(tc.want, got, floatComparer)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
	diff := cmp.Diff(4., sd.Next(4.), floatComparer)
	if diff == "" {
		t.Fatal(diff)
	}

	sd.Reset()
	diff = cmp.Diff(4., sd.Next(4.), floatComparer)
	if diff != "" {
		t.Fatalf(diff)
	}
}

func TestExponentialMovingAverageString(t *testing.T) {
	sd, _ := NewExponentialMovingAverage(4)
	want := "EMA(4)"
	got := sd.String()
	diff := cmp.Diff(want, got, floatComparer)
	if diff != "" {
		t.Fatalf(diff)
	}
}
