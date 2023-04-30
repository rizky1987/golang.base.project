package commonHelpers

import (
	"errors"
	"strings"

	mssql "github.com/denisenkom/go-mssqldb"
)

func StringSplitToArrayString(inputString, delimiter string) []string {
	return strings.Split(TrimWhiteSpace(inputString), TrimWhiteSpace(delimiter))
}

func TrimWhiteSpace(inputString string) string {
	return strings.TrimSpace(inputString)
}

func StringReplace(original, from, to string) string {
	return strings.ReplaceAll(original, from, to)
}

func StringToNewUUID(UUIDString string) (mssql.UniqueIdentifier, error) {
	var newUUID mssql.UniqueIdentifier

	err := newUUID.Scan(UUIDString)
	if err != nil {

		return newUUID, errors.New("Please input a valid UUID")
	}

	return newUUID, nil

}
