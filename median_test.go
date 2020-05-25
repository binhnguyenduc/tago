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

func TestNewMedian(t *testing.T) {
	tests := map[string]struct {
		input   int
		want    *Median
		wantErr error
	}{
		"negative n": {input: -3, want: nil, wantErr: ErrInvalidParameters},
		"zero n":     {input: 0, want: nil, wantErr: ErrInvalidParameters},
		"positive n": {input: 9, want: &Median{n: 9, data: make([]float64, 0, 9)}, wantErr: nil},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			gotSD, gotErr := NewMedian(tc.input)
			if tc.wantErr != nil { // only check error returned if expecting one
				assert.EqualError(t, gotErr, tc.wantErr.Error(), "must return the correct error")
			}
			assert.Equal(t, tc.want, gotSD, "must return the correct value")
		})
	}
}

func TestMedianNextOddLength(t *testing.T) {
	sd, _ := NewMedian(3)
	tests := []struct {
		input float64
		want  float64
	}{
		{input: 10., want: 10.},
		{input: 20., want: 15.},
		{input: 30., want: 20.},
		{input: 15., want: 20.},
		{input: 40., want: 30.},
		{input: 25., want: 25.},
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
func TestMedianNextEvenLength(t *testing.T) {
	sd, _ := NewMedian(4)
	tests := []struct {
		input float64
		want  float64
	}{
		{input: 10., want: 10.},
		{input: 20., want: 15.},
		{input: 30., want: 20.},
		{input: 15., want: 17.5},
		{input: 40., want: 25.},
		{input: 25., want: 27.5},
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

func TestMedianReset(t *testing.T) {
	sd, _ := NewMedian(4)
	tests := []struct {
		input float64
		want  float64
	}{
		{input: 10., want: 10.},
		{input: 20., want: 15.},
		{input: 30., want: 20.},
		{input: 15., want: 17.5},
		{input: 40., want: 25.},
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
	diff := cmp.Diff(25., sd.Next(25.), floatComparer)
	if diff != "" {
		t.Fatalf(diff)
	}
}

func TestMedianString(t *testing.T) {
	sd, _ := NewMedian(4)
	want := "Median(4)"
	got := sd.String()
	diff := cmp.Diff(want, got, floatComparer)
	if diff != "" {
		t.Fatalf(diff)
	}
}
