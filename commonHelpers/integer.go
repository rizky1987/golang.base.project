package commonHelpers

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func ConvertIntegerToString(inputInt int) string {
	return strconv.Itoa(inputInt)
}

func ToPadNumberWithZero(value, length int) string {
	return fmt.Sprintf("%0"+strconv.Itoa(length)+"d", value)
}

func RandomInteger() int {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 999999999
	return rand.Intn(max-min+1) + min
}

func ConvertNullAbleIntegerToInteger(input *int) int {

	if input == nil {
		return 0
	}

	return *input
}

func ConvertIntegerToBoolen(input int) bool {

	result := false

	if input == 1 {
		return true
	}
	return result
}
