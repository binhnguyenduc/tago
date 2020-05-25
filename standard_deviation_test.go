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

func TestNewStandardDeviation(t *testing.T) {
	tests := map[string]struct {
		input   int
		want    *StandardDeviation
		wantErr error
	}{
		"negative n": {input: -3, want: nil, wantErr: ErrInvalidParameters},
		"zero n":     {input: 0, want: nil, wantErr: ErrInvalidParameters},
		"positive n": {input: 9, want: &StandardDeviation{n: 9, data: make([]float64, 9)}, wantErr: nil},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			gotSD, gotErr := NewStandardDeviation(tc.input)
			if tc.wantErr != nil { // only check error returned if expecting one
				assert.EqualError(t, gotErr, tc.wantErr.Error(), "must return the correct error")
			}
			assert.Equal(t, tc.want, gotSD, "must return the correct value")
		})
	}
}

func TestNextStandardDeviation(t *testing.T) {
	sd, _ := NewStandardDeviation(4)
	tests := []struct {
		input float64
		want  float64
	}{
		{input: 10., want: 0.},
		{input: 20., want: 5.},
		{input: 30., want: 8.165},
		{input: 20., want: 7.071},
		{input: 10., want: 7.071},
		{input: 100., want: 35.355},
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

func TestStandardDeviationNextSameValue(t *testing.T) {
	sd, _ := NewStandardDeviation(3)
	tests := []struct {
		input float64
		want  float64
	}{
		{input: 10., want: 0.},
		{input: 10., want: 0.},
		{input: 10., want: 0.},
		{input: 10., want: 0.},
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

func TestStandardDeviationReset(t *testing.T) {
	sd, _ := NewStandardDeviation(4)
	tests := []struct {
		input float64
		want  float64
	}{
		{input: 10., want: 0.},
		{input: 20., want: 5.},
		{input: 30., want: 8.165},
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
	diff := cmp.Diff(0., sd.Next(20.), floatComparer)
	if diff != "" {
		t.Fatalf(diff)
	}
}

func TestStandardDeviationString(t *testing.T) {
	sd, _ := NewStandardDeviation(4)
	want := "SD(4)"
	got := sd.String()
	diff := cmp.Diff(want, got, floatComparer)
	if diff != "" {
		t.Fatalf(diff)
	}
}
