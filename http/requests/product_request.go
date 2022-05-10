package requests

type (
	CreateProductRequest struct {
		ProductCode            string `form:"productCode" validate:"required"`
		DosageDescription      string `form:"dosageDescription"`
		UsabilityDescription   string `form:"usabilityDescription"`
		CompositionDescription string `form:"composition"`
		HowToUseDescription    string `form:"howToUseDescription"`
	}
)
