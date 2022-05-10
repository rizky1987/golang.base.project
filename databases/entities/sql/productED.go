package entities

import (
	"errors"
	"time"

	mssql "github.com/denisenkom/go-mssqldb"
	"github.com/gofrs/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type (
	ProductED struct {
		ID             mssql.UniqueIdentifier `gorm:"column:Id"`
		ProductCode    string                 `gorm:"column:ProductCode"`
		PrincipalCode  string                 `gorm:"column:PrincipalCode"`
		StatusEd       int                    `gorm:"column:StatusEd"`
		StartEd        time.Time              `gorm:"column:StartEd"`
		EndEd          time.Time              `gorm:"column:EndEd"`
		CreatedBy      *string                `gorm:"column:CreatedBy"`
		CreatedByName  *string                `gorm:"column:CreatedByName"`
		CreatedDate    time.Time              `gorm:"column:CreatedDt"`
		ModifiedBy     *string                `gorm:"column:UpdatedBy"`
		ModifiedByName *string                `gorm:"column:UpdatedByName"`
		ModifiedDate   pq.NullTime            `gorm:"column:UpdatedDt"`
		DeletedBy      *string                `gorm:"column:DeletedBy"`
		DeletedByName  *string                `gorm:"column:DeletedByName"`
		DeletedDate    gorm.DeletedAt         `gorm:"column:DeletedDt"`
	}
)

func (x ProductED) TableName() string {
	return "ProductEd"
}

func (u *ProductED) BeforeCreate(tx *gorm.DB) error {
	u2, err := uuid.NewV4()
	if err != nil {
		errors.New("can't save invalid data")
		return err
	}
	err = u.ID.Scan(u2.String())
	if err != nil {
		errors.New("can't save invalid data")
		return err
	}
	return nil
}
