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

func TestNewMean(t *testing.T) {
	tests := map[string]struct {
		input   int
		want    *Mean
		wantErr error
	}{
		"negative n": {input: -3, want: nil, wantErr: ErrInvalidParameters},
		"zero n":     {input: 0, want: nil, wantErr: ErrInvalidParameters},
		"positive n": {input: 9, want: &Mean{n: 9, data: make([]float64, 9)}, wantErr: nil},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			gotSD, gotErr := NewMean(tc.input)
			if tc.wantErr != nil { // only check error returned if expecting one
				assert.EqualError(t, gotErr, tc.wantErr.Error(), "must return the correct error")
			}
			assert.Equal(t, tc.want, gotSD, "must return the correct value")
		})
	}
}

func TestNextMean(t *testing.T) {
	sd, _ := NewMean(4)
	tests := []struct {
		input float64
		want  float64
	}{
		{input: 10., want: 10.},
		{input: 20., want: 15.},
		{input: 30., want: 20.},
		{input: 20., want: 20.},
		{input: 10., want: 20.},
		{input: 100., want: 40.},
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

func TestMeanNextSameValue(t *testing.T) {
	sd, _ := NewMean(3)
	tests := []struct {
		input float64
		want  float64
	}{
		{input: 10., want: 10.},
		{input: 10., want: 10.},
		{input: 10., want: 10.},
		{input: 10., want: 10.},
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

func TestMeanReset(t *testing.T) {
	sd, _ := NewMean(4)
	tests := []struct {
		input float64
		want  float64
	}{
		{input: 10., want: 10.},
		{input: 20., want: 15.},
		{input: 30., want: 20.},
		{input: 20., want: 20.},
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
	diff := cmp.Diff(10., sd.Next(10.), floatComparer)
	if diff != "" {
		t.Fatalf(diff)
	}
}

func TestMeanString(t *testing.T) {
	sd, _ := NewMean(4)
	want := "Mean(4)"
	got := sd.String()
	diff := cmp.Diff(want, got, floatComparer)
	if diff != "" {
		t.Fatalf(diff)
	}
}
