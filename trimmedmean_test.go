package trimmedmean

import (
	"math"
	"testing"
)

func floatEquals(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}

func TestComputeTrimmedMean_Symmetric(t *testing.T) {
	data := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := 5.5 // mean of 2-9 (with 10% trimming from both ends)

	result, err := ComputeTrimmedMean(data, 0.1, 0.1)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
	if !floatEquals(result, expected, 1e-6) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestComputeTrimmedMean_Asymmetric(t *testing.T) {
	data := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := 7.0 // mean of 4-10 (30% low trim)

	result, err := ComputeTrimmedMean(data, 0.3, 0.0)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
	if !floatEquals(result, expected, 1e-6) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestComputeSymmetricTrimmedMean(t *testing.T) {
	data := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := 5.5 // mean of 2-9 (with 10% symmetric trimming)

	result, err := ComputeSymmetricTrimmedMean(data, 0.1)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
	if !floatEquals(result, expected, 1e-6) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestComputeTrimmedMean_InvalidTrim(t *testing.T) {
	data := []float64{1, 2, 3}
	_, err := ComputeTrimmedMean(data, -0.1, 0.1)
	if err == nil {
		t.Error("Expected error for negative trimming")
	}

	_, err = ComputeTrimmedMean(data, 0.6, 0.5)
	if err == nil {
		t.Error("Expected error for excessive trimming")
	}
}

func TestComputeTrimmedMean_Empty(t *testing.T) {
	data := []float64{}
	_, err := ComputeTrimmedMean(data, 0.1, 0.1)
	if err == nil {
		t.Error("Expected error for empty data slice")
	}
}
