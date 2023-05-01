package requests

type (
	CreateBookingRequest struct {
		StartDate        string                        `json:"startDate" validate:"required"`
		EndDate          string                        `json:"endDate" validate:"required"`
		DownPayment      int                           `json:"downPayment" validate:"required"`
		BookedBy         string                        `json:"BookedBy" validate:"required"`
		IsTimeRulesAgree bool                          `json:"isTimeRulesAgree"`
		BookingDetails   []*CreateBookingDetailRequest `json:"bookingDetails" validate:"required"`
	}

	CreateBookingDetailRequest struct {
		RoomId string `json:"roomId" validate:"required"`
	}
)
