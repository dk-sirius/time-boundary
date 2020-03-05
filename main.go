package main

import (
	"fmt"
	"time"

	"github.com/e-Sirius/tools/constants"
	toolTime "github.com/e-Sirius/tools/time"
)

func main() {
	s := toolTime.CalTimeRange(time.Now(), constants.TIME_TYPE__DAY)
	s1 := toolTime.CalTimeRange(time.Now(), constants.TIME_TYPE__WEEK)
	s2 := toolTime.CalTimeRange(time.Now(), constants.TIME_TYPE__YEAR)
	s3 := toolTime.CalTimeRange(time.Now(), constants.TIME_TYPE__SEASON)
	s4 := toolTime.CalTimeRange(time.Now(), constants.TIME_TYPE__MON)
	fmt.Println(s, "\n", s1, "\n", s2, "\n", s3, "\n", s4)
}
