// trimmedmean/trimmedmean.go
package trimmedmean

import (
	"errors"
	"sort"
)

// ComputeSymmetricTrimmedMean computes the symmetric trimmed mean of a float64 slice.
func ComputeSymmetricTrimmedMean(data []float64, trim float64) (float64, error) {
	return ComputeTrimmedMean(data, trim, trim)
}

// ComputeTrimmedMean computes the trimmed mean of a float64 slice.
// lowTrim and highTrim are proportions (0-1) of elements to trim from each end.
func ComputeTrimmedMean(data []float64, lowTrim, highTrim float64) (float64, error) {
	if len(data) == 0 {
		return 0, errors.New("empty data slice")
	}
	if lowTrim < 0 || highTrim < 0 || lowTrim+highTrim >= 1 {
		return 0, errors.New("invalid trimming proportions")
	}
	// Sort data
	sorted := make([]float64, len(data))
	copy(sorted, data)
	sort.Float64s(sorted)

	// Calculate indices to trim
	n := len(sorted)
	lowCount := int(float64(n) * lowTrim)
	highCount := int(float64(n) * highTrim)
	trimmed := sorted[lowCount : n-highCount]

	// Compute mean
	sum := 0.0
	for _, v := range trimmed {
		sum += v
	}
	return sum / float64(len(trimmed)), nil
}
