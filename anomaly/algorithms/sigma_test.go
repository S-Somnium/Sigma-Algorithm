package algorithms

import (
	"testing"
)

var sample = &[]float64{4, 14, 16, 22, 24, 64}

func TestSigmaDetection(t *testing.T) {
	s := GetSigma(sample)
	warns, alarms, err := s.Detect(2, 3)
	if err != nil {
		t.Errorf("SigmaDetection FAILED. Expected %d %d, got %d %d. Parameters: 4, 14, 16, 22, 24, 64\n", 1, 0, warns, alarms)
	} else if (len(warns) != 1) && (len(alarms) != 0) {
		t.Errorf("SigmaDetection FAILED. Expected %d %d, got %d %d. Parameters: 4, 14, 16, 22, 24, 64\n", 1, 0, warns, alarms)
	} else {
		t.Logf("SigmaDetection PASSED. Expected %d %d, got %d %d. Parameters: 4, 14, 16, 22, 24, 64\n", 1, 0, warns, alarms)
	}

}

func TestMeanCalc(t *testing.T) {
	s := GetSigma(sample)
	s.meanCalc()
	if s.mean != 24.00 {
		t.Errorf("meanCalc(4, 14, 16, 22, 24, 64) FAILED. Expected %f, got %f\n", 24.0, s.mean)
	} else {
		t.Logf("meanCalc(4, 14, 16, 22, 24, 64) PASSED. Expected %f, got %f\n", 24.0, s.mean)
	}
}

func TestDeviationCalc(t *testing.T) {
	s := GetSigma(sample)
	s.meanCalc()
	s.deviationCalc()
	if s.derivation != 19.008769905844336 {
		t.Errorf("deviationCalc FAILED. Expected %f, got %f. Parameters: 4, 14, 16, 22, 24, 64\n", 19.008769905844336, s.derivation)
	} else {
		t.Logf("deviationCalc PASSED. Expected %f, got %f. Parameters: 4, 14, 16, 22, 24, 64\n", 19.008769905844336, s.derivation)
	}
}

func TestSigmaCalc(t *testing.T) {
	s := GetSigma(sample)
	s.meanCalc()
	s.deviationCalc()
	result := s.sigmaCalc(4)
	if result != 1.0521459357478415 {
		t.Errorf("sigmaCalc(4) FAILED. Expected %f, got %f\n", 1.0521459357478415, result)
	} else {
		t.Logf("sigmaCalc(4) PASSED. Expected %f, got %f\n", 1.0521459357478415, result)
	}

}
