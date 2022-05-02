package anomaly

import (
	"fmt"
)

type Algorithm interface {
	Detect(warnTrigger float64, alarmTrigger float64) (warns []int, alarms []int, err error)
	SetList(list *[]float64)
}

func AutoDetect(outliers Algorithm) {
	fmt.Println(outliers.Detect(2, 3))
}
