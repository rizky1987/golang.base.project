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
	Product struct {
		ID                     mssql.UniqueIdentifier `gorm:"column:Id"`
		ProductCode            string                 `gorm:"column:ProductCode"`
		DosageDescription      *string                `gorm:"column:DosageDescription"`
		UsabilityDescription   *string                `gorm:"column:UsabilityDescription"`
		CompositionDescription *string                `gorm:"column:CompositionDescription"`
		HowToUseDescription    *string                `gorm:"column:HowToUseDescription"`
		CreatedBy              string                 `gorm:"column:CreatedBy"`
		CreatedByName          string                 `gorm:"column:CreatedByName"`
		CreatedDate            time.Time              `gorm:"column:CreatedDate"`
		ModifiedBy             *string                `gorm:"column:ModifiedBy"`
		ModifiedByName         *string                `gorm:"column:ModifiedByName"`
		ModifiedDate           pq.NullTime            `gorm:"column:ModifiedDate"`
		DeletedBy              *string                `gorm:"column:DeletedBy"`
		DeletedByName          *string                `gorm:"column:DeletedByName"`
		DeletedDate            *gorm.DeletedAt        `gorm:"column:DeletedDate"`
	}
)

func (x Product) TableName() string {
	return "Product"
}

func (u *Product) BeforeCreate(tx *gorm.DB) error {
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
