package main

import (
	"sigma-algorithm/anomaly"
	"sigma-algorithm/anomaly/algorithms"
)

func main() {
	//receive the dataset from somewhere with its relevance attributes.
	var sample = &[]float64{4, 14, 16, 22, 24, 64}
	alg := algorithms.GetSigma(sample)
	anomaly.AutoDetect(alg)

}
