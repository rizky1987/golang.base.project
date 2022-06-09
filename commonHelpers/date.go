package commonHelpers

import (
	"strings"
	"time"
)

func GetCurrentTimeUTC() time.Time {
	return time.Now().In(time.UTC)
}

func GetHorizonTimeOnStringType() string {
	return time.Now().In(time.UTC).Add(7 * time.Hour).String()
}

func GetCurrentTimeAsiaJakarta() time.Time {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	return time.Now().In(loc).Add(7 * time.Hour)
}

func GetCurrentTimeUTCOnStringFormatYYYYMMDD() string {
	return time.Now().In(time.UTC).Format("2006/01/02")
}

func ConvertDateToStringFormatYYYYMMDDHIS(date *time.Time) string {
	if date == nil {
		return ""
	}
	return date.Format("2006/01/02 15:04:05")
}

func ConvertDateToStringFormatYYYYMMDD(date *time.Time) string {
	if date == nil {
		return ""
	}
	return date.Format("2006/01/02")
}

func ConvertDateToStringFormatMMYYYY(date *time.Time) string {
	if date == nil {
		return ""
	}
	return date.Format("01-2006")
}

func ConvertTimeToStringFormatHIS(date *time.Time) string {
	if date == nil {
		return ""
	}
	return date.Format("15:04:05")
}

func ConvertStringToDateFormatYYYYMMDD(dateString string) (*time.Time, error) {
	date, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		return nil, err
	}

	return &date, nil
}

func ConvertStringToDateFormatRFC3339(dateString string) (*time.Time, error) {
	date, err := time.Parse("2006-01-02T15:04:05.000Z", dateString)
	if err != nil {
		return nil, err
	}

	return &date, nil
}

func IsFirstDateBeforeSecondDate(firstDate, SecondDate *time.Time) bool {

	return firstDate.Before(*SecondDate)
}

func IsFirstDateEqualSecondDate(firstDate, SecondDate *time.Time) bool {

	return firstDate.Equal(*SecondDate)
}

func IsFirstDateAfterSecondDate(firstDate, SecondDate *time.Time) bool {

	return firstDate.After(*SecondDate)
}

func IsFirstDateBeforeOrEqualSecondDate(firstDate, SecondDate *time.Time) bool {

	return IsFirstDateBeforeSecondDate(firstDate, SecondDate) || IsFirstDateEqualSecondDate(firstDate, SecondDate)
}

func IsFirstDateAfterOrEqualSecondDate(firstDate, SecondDate *time.Time) bool {

	return IsFirstDateAfterSecondDate(firstDate, SecondDate) || IsFirstDateEqualSecondDate(firstDate, SecondDate)
}

func GetCurrentTimeUTConStringFormatDDMMYYYY() string {

	return time.Now().In(time.UTC).Format("02/01/2006")
}

func GetTimeHorizonFromRFC3339String(dateTime string) string {
	explode := strings.Split(dateTime, "T")
	replace := strings.Replace(explode[1], "Z", "", 1)
	return replace
}
