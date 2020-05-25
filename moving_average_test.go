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

func TestNewMovingAverage(t *testing.T) {
	tests := map[string]struct {
		input   int
		want    *MovingAverage
		wantErr error
	}{
		"negative n": {input: -3, want: nil, wantErr: ErrInvalidParameters},
		"zero n":     {input: 0, want: nil, wantErr: ErrInvalidParameters},
		"positive n": {input: 9, want: &MovingAverage{n: 9, data: make([]float64, 9)}, wantErr: nil},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			gotSD, gotErr := NewMovingAverage(tc.input)
			if tc.wantErr != nil { // only check error returned if expecting one
				assert.EqualError(t, gotErr, tc.wantErr.Error(), "must return the correct error")
			}
			assert.Equal(t, tc.want, gotSD, "must return the correct value")
		})
	}
}

func TestMovingAverageNext(t *testing.T) {
	sd, _ := NewMovingAverage(4)
	tests := []struct {
		input float64
		want  float64
	}{
		{input: 4., want: 4.},
		{input: 5., want: 4.5},
		{input: 6., want: 5.},
		{input: 6., want: 5.25},
		{input: 6., want: 5.75},
		{input: 6., want: 6.},
		{input: 2., want: 5.},
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

func TestMovingAverageReset(t *testing.T) {
	sd, _ := NewMovingAverage(4)
	tests := []struct {
		input float64
		want  float64
	}{
		{input: 4., want: 4.},
		{input: 5., want: 4.5},
		{input: 6., want: 5.},
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

	sd.Reset()
	diff := cmp.Diff(20., sd.Next(20.), floatComparer)
	if diff != "" {
		t.Fatalf(diff)
	}
}

func TestMovingAverageString(t *testing.T) {
	sd, _ := NewMovingAverage(4)
	want := "MA(4)"
	got := sd.String()
	diff := cmp.Diff(want, got, floatComparer)
	if diff != "" {
		t.Fatalf(diff)
	}
}
