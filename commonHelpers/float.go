package commonHelpers

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func Float64ToDateTimeUTC(input float64) time.Time {
	ts := fmt.Sprintf("%f", input)
	v := strings.Split(ts, ".")
	if len(v[1]) > 0 {
		for len(v[1]) < 9 {
			v[1] += "0"
		}
	}

	a, _ := strconv.ParseInt(v[0], 10, 64)
	b, _ := strconv.ParseInt(v[1], 10, 64)
	t := time.Unix(a, b).UnixNano()
	return time.Unix(0, t).UTC()
}
