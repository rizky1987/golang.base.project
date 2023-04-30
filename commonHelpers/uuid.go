package commonHelpers

import (
	mssql "github.com/denisenkom/go-mssqldb"
	"github.com/gofrs/uuid"
)

func GenerateNewUUID() mssql.UniqueIdentifier {
	var newUUID mssql.UniqueIdentifier

	u2, _ := uuid.NewV4()

	newUUID.Scan(u2.String())

	return newUUID

}
