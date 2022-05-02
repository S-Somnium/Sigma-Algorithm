package algorithms

import (
	"errors"
	"math"
)

type sigma struct {
	list       *[]float64
	mean       float64
	derivation float64
}

func GetSigma(list *[]float64) *sigma {
	// does list size can be 0?
	s := sigma{
		list: list,
	}
	return &s
}

func (s *sigma) SetList(list *[]float64) {
	s.list = list
}

// calculate the sigma value of each element on the list and if its above the warnTrigger or alarmTrigger will return the position of each warn/alarm.
func (s *sigma) Detect(warnTrigger float64, alarmTrigger float64) (warns []int, alarms []int, err error) {
	if s.list == nil {
		return nil, nil, errors.New("list is nil")
	}
	s.meanCalc()
	s.deviationCalc()
	for i, x := range *s.list {
		anomaly := s.sigmaCalc(x)
		if anomaly >= alarmTrigger {
			alarms = append(alarms, i)
			continue
		}
		if anomaly >= warnTrigger {
			warns = append(warns, i)
			continue
		}
	}
	return warns, alarms, nil
}

func (s *sigma) meanCalc() {
	mean := 0.0
	for _, x := range *s.list {
		mean += x
	}
	s.mean = mean / float64(len(*s.list))
}

// Calculates the deviation value of an array, considering it as a population.
func (s *sigma) deviationCalc() {
	derivation := 0.0
	for _, x := range *s.list {
		derivation += math.Pow((x - s.mean), 2)
	}
	// sample or *population ?
	derivation = derivation / float64(len(*s.list))
	s.derivation = math.Sqrt(derivation)
}

// return the sigma position of a single item.
func (s *sigma) sigmaCalc(value float64) float64 {
	return math.Abs(((value - s.mean) / s.derivation))
}
