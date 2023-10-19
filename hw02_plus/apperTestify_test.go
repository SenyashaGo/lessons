package main

import (
	"github.com/kulti/titlecase"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestEmpty2(t *testing.T) {
	const str, minor, want = "", "", ""
	got := titlecase.TitleCase(str, minor)
	require.Equalf(t, want, got, "TitleCase(%v, %v) = %v; want %v", str, minor, got, want)
}

func TestWithoutMinor2(t *testing.T) {
	const str, minor, want = "the quick fox in the bag", "", "The Quick Fox In The Bag"
	got := titlecase.TitleCase(str, minor)
	require.Equalf(t, want, got, "TitleCase(%v, %v) = %v; want %v", str, minor, got, want)
}

func TestWithMinorInFirst2(t *testing.T) {
	const str, minor, want = "the quick fox in the bag", "in the", "The Quick Fox in the Bag"
	got := titlecase.TitleCase(str, minor)
	require.Equalf(t, want, got, "TitleCase(%v, %v) = %v; want %v", str, minor, got, want)
}

func TestExample2(t *testing.T) {
	const str, minor, want = "The quick fox in the bag", "The the", "The Quick Fox In the Bag"
	got := titlecase.TitleCase(str, minor)
	require.Equalf(t, want, got, "TitleCase(%v, %v) = %v; want %v", str, minor, got, want)
}
