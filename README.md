# Go Trimmed Mean Package

This is a Go package for calculating the **trimmed mean**, supporting both symmetric and asymmetric trimming.

---

# Features
- Supports symmetric trimming (low and high ends trimmed by the same proportion) or a simplified single-parameter interface.  
- Supports asymmetric trimming (low and high ends trimmed by different proportions).  
- Sorts the data before trimming.  
- Validates input (non-empty slice, valid proportions, etc.).  

## Usage

1, Initialize a Go module in your project (if not already done):  

go mod init yourprojects  

2, Get the trimmed mean package:  
go get [gitlink](https://github.com/MikeLiang2/gotrimmedmean.git)  

## symmetric trimming usage
import "[gitlink](https://github.com/MikeLiang2/gotrimmedmean.git)"  

data := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}  

// Symmetric trimming: trim 10% from both the low end and high end.  
// This removes the lowest 10% and highest 10% of sorted data elements.  
meanSym, err := trimmedmean.ComputeSymmetricTrimmedMean(data, 0.1)  
if err != nil {  
    // Handle error  
}  
fmt.Printf("Symmetric trimmed mean (10%%): %.2f\n", meanSym)  


## asymmetric trimming
import "[gitlink](https://github.com/MikeLiang2/gotrimmedmean.git)"  

data := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}  

// Asymmetric trimming: trim 20% from the low end and 10% from the high end.  
// This removes the lowest 20% and highest 10% of sorted data elements.  
meanAsym, err := trimmedmean.ComputeTrimmedMean(data, 0.2, 0.1)  
if err != nil {  
    // Handle error  
}  
fmt.Printf("Asymmetric trimmed mean (low 20%%, high 10%%): %.2f\n", meanAsym)  
